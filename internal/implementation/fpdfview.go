package implementation

// #cgo pkg-config: pdfium
// #include "fpdfview.h"
import "C"
import (
	"errors"
	"github.com/google/uuid"

	pdfium_errors "github.com/klippa-app/go-pdfium/errors"
	"github.com/klippa-app/go-pdfium/references"
	"github.com/klippa-app/go-pdfium/requests"
	"github.com/klippa-app/go-pdfium/responses"
)

// FPDF_LoadDocument opens and load a PDF document from a file path.
// Loaded document can be closed by FPDF_CloseDocument().
// If this function fails, you can use FPDF_GetLastError() to retrieve
// the reason why it failed.
func (p *PdfiumImplementation) FPDF_LoadDocument(request *requests.FPDF_LoadDocument) (*responses.FPDF_LoadDocument, error) {
	// Don't lock, OpenDocument will do that.
	doc, err := p.OpenDocument(&requests.OpenDocument{
		FilePath: request.Path,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	return &responses.FPDF_LoadDocument{
		Document: doc.Document,
	}, nil
}

// FPDF_LoadMemDocument opens and load a PDF document from memory.
// Loaded document can be closed by FPDF_CloseDocument().
// If this function fails, you can use FPDF_GetLastError() to retrieve
// the reason why it failed.
func (p *PdfiumImplementation) FPDF_LoadMemDocument(request *requests.FPDF_LoadMemDocument) (*responses.FPDF_LoadMemDocument, error) {
	// Don't lock, OpenDocument will do that.
	doc, err := p.OpenDocument(&requests.OpenDocument{
		File:     request.Data,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	return &responses.FPDF_LoadMemDocument{
		Document: doc.Document,
	}, nil
}

// FPDF_LoadMemDocument64 opens and load a PDF document from memory.
// Loaded document can be closed by FPDF_CloseDocument().
// If this function fails, you can use FPDF_GetLastError() to retrieve
// the reason why it failed.
func (p *PdfiumImplementation) FPDF_LoadMemDocument64(request *requests.FPDF_LoadMemDocument64) (*responses.FPDF_LoadMemDocument64, error) {
	// Don't lock, OpenDocument will do that.
	doc, err := p.OpenDocument(&requests.OpenDocument{
		File:     request.Data,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	return &responses.FPDF_LoadMemDocument64{
		Document: doc.Document,
	}, nil
}

// FPDF_LoadCustomDocument loads a PDF document from a custom access descriptor.
// This is implemented as an io.ReadSeeker in go-pdfium.
// This is only really efficient for single threaded usage, the multi-threaded
// usage will just load the file in memory because it can't transfer readers
// over gRPC. The single-threaded usage will actually efficiently walk over
// the PDF as it's being used by PDFium.
// Loaded document can be closed by FPDF_CloseDocument().
// If this function fails, you can use FPDF_GetLastError() to retrieve
// the reason why it failed.
func (p *PdfiumImplementation) FPDF_LoadCustomDocument(request *requests.FPDF_LoadCustomDocument) (*responses.FPDF_LoadCustomDocument, error) {
	// Don't lock, OpenDocument will do that.
	doc, err := p.OpenDocument(&requests.OpenDocument{
		FileReader:     request.Reader,
		FileReaderSize: request.Size,
		Password:       request.Password,
	})
	if err != nil {
		return nil, err
	}

	return &responses.FPDF_LoadCustomDocument{
		Document: doc.Document,
	}, nil
}

// FPDF_CloseDocument closes the references, releases the resources.
func (p *PdfiumImplementation) FPDF_CloseDocument(request *requests.FPDF_CloseDocument) (*responses.FPDF_CloseDocument, error) {
	p.Lock()
	defer p.Unlock()

	nativeDocument, err := p.getDocumentHandle(request.Document)
	if err != nil {
		return nil, err
	}

	err = nativeDocument.Close()
	if err != nil {
		return nil, err
	}

	delete(p.documentRefs, nativeDocument.nativeRef)

	return &responses.FPDF_CloseDocument{}, nil
}

// FPDF_GetLastError returns the last error generated by PDFium.
func (p *PdfiumImplementation) FPDF_GetLastError(request *requests.FPDF_GetLastError) (*responses.FPDF_GetLastError, error) {
	p.Lock()
	defer p.Unlock()

	return &responses.FPDF_GetLastError{
		Error: responses.FPDF_GetLastErrorError(C.FPDF_GetLastError()),
	}, nil
}

// FPDF_SetSandBoxPolicy set the policy for the sandbox environment.
func (p *PdfiumImplementation) FPDF_SetSandBoxPolicy(request *requests.FPDF_SetSandBoxPolicy) (*responses.FPDF_SetSandBoxPolicy, error) {
	p.Lock()
	defer p.Unlock()

	enable := C.int(0)
	if request.Enable {
		enable = C.int(1)
	}

	C.FPDF_SetSandBoxPolicy(C.FPDF_DWORD(request.Policy), enable)

	return &responses.FPDF_SetSandBoxPolicy{}, nil
}

// FPDF_LoadPage loads a page and returns a reference.
func (p *PdfiumImplementation) FPDF_LoadPage(request *requests.FPDF_LoadPage) (*responses.FPDF_LoadPage, error) {
	p.Lock()
	defer p.Unlock()

	documentHandle, err := p.getDocumentHandle(request.Document)
	if err != nil {
		return nil, err
	}

	pageObject := C.FPDF_LoadPage(documentHandle.handle, C.int(request.Index))
	if pageObject == nil {
		return nil, pdfium_errors.ErrPage
	}

	pageRef := uuid.New()
	pageHandle := &PageHandle{
		handle:      pageObject,
		index:       request.Index,
		nativeRef:   references.FPDF_PAGE(pageRef.String()),
		documentRef: documentHandle.nativeRef,
	}

	documentHandle.pageRefs[pageHandle.nativeRef] = pageHandle
	p.pageRefs[pageHandle.nativeRef] = pageHandle

	return &responses.FPDF_LoadPage{
		Page: pageHandle.nativeRef,
	}, nil
}

// FPDF_ClosePage unloads a page by reference.
func (p *PdfiumImplementation) FPDF_ClosePage(request *requests.FPDF_ClosePage) (*responses.FPDF_ClosePage, error) {
	p.Lock()
	defer p.Unlock()

	pageRef, err := p.getPageHandle(request.Page)
	if err != nil {
		return nil, err
	}

	pageRef.Close()
	delete(p.pageRefs, request.Page)

	// Remove page reference from document.
	documentHandle, err := p.getDocumentHandle(pageRef.documentRef)
	if err != nil {
		return nil, err
	}
	delete(documentHandle.pageRefs, request.Page)

	return &responses.FPDF_ClosePage{}, nil
}

// FPDF_GetFileVersion returns the version of the PDF file.
func (p *PdfiumImplementation) FPDF_GetFileVersion(request *requests.FPDF_GetFileVersion) (*responses.FPDF_GetFileVersion, error) {
	p.Lock()
	defer p.Unlock()

	documentHandle, err := p.getDocumentHandle(request.Document)
	if err != nil {
		return nil, err
	}

	fileVersion := C.int(0)

	success := C.FPDF_GetFileVersion(documentHandle.handle, &fileVersion)
	if int(success) == 0 {
		return nil, errors.New("could not get file version")
	}

	return &responses.FPDF_GetFileVersion{
		FileVersion: int(fileVersion),
	}, nil
}

// FPDF_GetDocPermissions returns the permissions of the PDF.
func (p *PdfiumImplementation) FPDF_GetDocPermissions(request *requests.FPDF_GetDocPermissions) (*responses.FPDF_GetDocPermissions, error) {
	p.Lock()
	defer p.Unlock()

	documentHandle, err := p.getDocumentHandle(request.Document)
	if err != nil {
		return nil, err
	}

	permissions := C.FPDF_GetDocPermissions(documentHandle.handle)

	docPermissions := &responses.FPDF_GetDocPermissions{
		DocPermissions: uint32(permissions),
	}

	PrintDocument := uint32(1 << 2)
	ModifyContents := uint32(1 << 3)
	CopyOrExtractText := uint32(1 << 4)
	AddOrModifyTextAnnotations := uint32(1 << 5)
	FillInExistingInteractiveFormFields := uint32(1 << 8)
	ExtractTextAndGraphics := uint32(1 << 9)
	AssembleDocument := uint32(1 << 10)
	PrintDocumentAsFaithfulDigitalCopy := uint32(1 << 11)

	hasPermission := func(permission uint32) bool {
		if docPermissions.DocPermissions&permission > 0 {
			return true
		}

		return false
	}

	docPermissions.PrintDocument = hasPermission(PrintDocument)
	docPermissions.ModifyContents = hasPermission(ModifyContents)
	docPermissions.CopyOrExtractText = hasPermission(CopyOrExtractText)
	docPermissions.AddOrModifyTextAnnotations = hasPermission(AddOrModifyTextAnnotations)
	docPermissions.FillInInteractiveFormFields = hasPermission(AddOrModifyTextAnnotations)
	docPermissions.FillInExistingInteractiveFormFields = hasPermission(FillInExistingInteractiveFormFields)
	docPermissions.ExtractTextAndGraphics = hasPermission(ExtractTextAndGraphics)
	docPermissions.AssembleDocument = hasPermission(AssembleDocument)
	docPermissions.PrintDocumentAsFaithfulDigitalCopy = hasPermission(PrintDocumentAsFaithfulDigitalCopy)

	// Calculated permissions
	docPermissions.CreateOrModifyInteractiveFormFields = docPermissions.ModifyContents && docPermissions.AddOrModifyTextAnnotations

	return docPermissions, nil
}

// FPDF_GetSecurityHandlerRevision returns the revision number of security handlers of the file.
func (p *PdfiumImplementation) FPDF_GetSecurityHandlerRevision(request *requests.FPDF_GetSecurityHandlerRevision) (*responses.FPDF_GetSecurityHandlerRevision, error) {
	p.Lock()
	defer p.Unlock()

	documentHandle, err := p.getDocumentHandle(request.Document)
	if err != nil {
		return nil, err
	}

	securityHandlerRevision := C.FPDF_GetSecurityHandlerRevision(documentHandle.handle)

	return &responses.FPDF_GetSecurityHandlerRevision{
		SecurityHandlerRevision: int(securityHandlerRevision),
	}, nil
}

// FPDF_GetPageCount counts the amount of pages.
func (p *PdfiumImplementation) FPDF_GetPageCount(request *requests.FPDF_GetPageCount) (*responses.FPDF_GetPageCount, error) {
	p.Lock()
	defer p.Unlock()

	documentHandle, err := p.getDocumentHandle(request.Document)
	if err != nil {
		return nil, err
	}

	return &responses.FPDF_GetPageCount{
		PageCount: int(C.FPDF_GetPageCount(documentHandle.handle)),
	}, nil
}

// FPDF_GetPageWidth returns the width of a page.
func (p *PdfiumImplementation) FPDF_GetPageWidth(request *requests.FPDF_GetPageWidth) (*responses.FPDF_GetPageWidth, error) {
	p.Lock()
	defer p.Unlock()

	pageHandle, err := p.loadPage(request.Page)
	if err != nil {
		return nil, err
	}

	width := C.FPDF_GetPageWidth(pageHandle.handle)

	return &responses.FPDF_GetPageWidth{
		Page:  pageHandle.index,
		Width: float64(width),
	}, nil
}

// FPDF_GetPageHeight returns the height of a page.
func (p *PdfiumImplementation) FPDF_GetPageHeight(request *requests.FPDF_GetPageHeight) (*responses.FPDF_GetPageHeight, error) {
	p.Lock()
	defer p.Unlock()

	pageHandle, err := p.loadPage(request.Page)
	if err != nil {
		return nil, err
	}

	height := C.FPDF_GetPageHeight(pageHandle.handle)

	return &responses.FPDF_GetPageHeight{
		Page:   pageHandle.index,
		Height: float64(height),
	}, nil
}

// FPDF_GetPageSizeByIndex returns the size of a page by the page index.
func (p *PdfiumImplementation) FPDF_GetPageSizeByIndex(request *requests.FPDF_GetPageSizeByIndex) (*responses.FPDF_GetPageSizeByIndex, error) {
	p.Lock()
	defer p.Unlock()

	documentHandle, err := p.getDocumentHandle(request.Document)
	if err != nil {
		return nil, err
	}

	width := C.double(0)
	height := C.double(0)

	result := C.FPDF_GetPageSizeByIndex(documentHandle.handle, C.int(request.Index), &width, &height)
	if int(result) == 0 {
		return nil, errors.New("could not load page size by index")
	}

	return &responses.FPDF_GetPageSizeByIndex{
		Page:   request.Index,
		Width:  float64(width),
		Height: float64(height),
	}, nil
}