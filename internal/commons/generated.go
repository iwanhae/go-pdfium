// Code generated by tool. DO NOT EDIT.
// See the code_generation package.

package commons

import (
	"github.com/klippa-app/go-pdfium/requests"
	"github.com/klippa-app/go-pdfium/responses"
)

type Pdfium interface {
	Ping() (string, error)
	OpenDocument(*requests.OpenDocument) error
    GetDocPermissions(*requests.GetDocPermissions) (*responses.GetDocPermissions, error)
    GetFileVersion(*requests.GetFileVersion) (*responses.GetFileVersion, error)
    GetMetadata(*requests.GetMetadata) (*responses.GetMetadata, error)
    GetPageCount(*requests.GetPageCount) (*responses.GetPageCount, error)
    GetPageSize(*requests.GetPageSize) (*responses.GetPageSize, error)
    GetPageSizeInPixels(*requests.GetPageSizeInPixels) (*responses.GetPageSizeInPixels, error)
    GetPageText(*requests.GetPageText) (*responses.GetPageText, error)
    GetPageTextStructured(*requests.GetPageTextStructured) (*responses.GetPageTextStructured, error)
    GetSecurityHandlerRevision(*requests.GetSecurityHandlerRevision) (*responses.GetSecurityHandlerRevision, error)
    RenderPageInDPI(*requests.RenderPageInDPI) (*responses.RenderPage, error)
    RenderPageInPixels(*requests.RenderPageInPixels) (*responses.RenderPage, error)
    RenderPagesInDPI(*requests.RenderPagesInDPI) (*responses.RenderPages, error)
    RenderPagesInPixels(*requests.RenderPagesInPixels) (*responses.RenderPages, error)
    RenderToFile(*requests.RenderToFile) (*responses.RenderToFile, error)
	Close() error
}


func (g *PdfiumRPC) GetDocPermissions(request *requests.GetDocPermissions) (*responses.GetDocPermissions, error) {
	resp := &responses.GetDocPermissions{}
	err := g.client.Call("Plugin.GetDocPermissions", request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *PdfiumRPC) GetFileVersion(request *requests.GetFileVersion) (*responses.GetFileVersion, error) {
	resp := &responses.GetFileVersion{}
	err := g.client.Call("Plugin.GetFileVersion", request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *PdfiumRPC) GetMetadata(request *requests.GetMetadata) (*responses.GetMetadata, error) {
	resp := &responses.GetMetadata{}
	err := g.client.Call("Plugin.GetMetadata", request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *PdfiumRPC) GetPageCount(request *requests.GetPageCount) (*responses.GetPageCount, error) {
	resp := &responses.GetPageCount{}
	err := g.client.Call("Plugin.GetPageCount", request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *PdfiumRPC) GetPageSize(request *requests.GetPageSize) (*responses.GetPageSize, error) {
	resp := &responses.GetPageSize{}
	err := g.client.Call("Plugin.GetPageSize", request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *PdfiumRPC) GetPageSizeInPixels(request *requests.GetPageSizeInPixels) (*responses.GetPageSizeInPixels, error) {
	resp := &responses.GetPageSizeInPixels{}
	err := g.client.Call("Plugin.GetPageSizeInPixels", request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *PdfiumRPC) GetPageText(request *requests.GetPageText) (*responses.GetPageText, error) {
	resp := &responses.GetPageText{}
	err := g.client.Call("Plugin.GetPageText", request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *PdfiumRPC) GetPageTextStructured(request *requests.GetPageTextStructured) (*responses.GetPageTextStructured, error) {
	resp := &responses.GetPageTextStructured{}
	err := g.client.Call("Plugin.GetPageTextStructured", request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *PdfiumRPC) GetSecurityHandlerRevision(request *requests.GetSecurityHandlerRevision) (*responses.GetSecurityHandlerRevision, error) {
	resp := &responses.GetSecurityHandlerRevision{}
	err := g.client.Call("Plugin.GetSecurityHandlerRevision", request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *PdfiumRPC) RenderPageInDPI(request *requests.RenderPageInDPI) (*responses.RenderPage, error) {
	resp := &responses.RenderPage{}
	err := g.client.Call("Plugin.RenderPageInDPI", request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *PdfiumRPC) RenderPageInPixels(request *requests.RenderPageInPixels) (*responses.RenderPage, error) {
	resp := &responses.RenderPage{}
	err := g.client.Call("Plugin.RenderPageInPixels", request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *PdfiumRPC) RenderPagesInDPI(request *requests.RenderPagesInDPI) (*responses.RenderPages, error) {
	resp := &responses.RenderPages{}
	err := g.client.Call("Plugin.RenderPagesInDPI", request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *PdfiumRPC) RenderPagesInPixels(request *requests.RenderPagesInPixels) (*responses.RenderPages, error) {
	resp := &responses.RenderPages{}
	err := g.client.Call("Plugin.RenderPagesInPixels", request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *PdfiumRPC) RenderToFile(request *requests.RenderToFile) (*responses.RenderToFile, error) {
	resp := &responses.RenderToFile{}
	err := g.client.Call("Plugin.RenderToFile", request, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}



func (s *PdfiumRPCServer) GetDocPermissions(request *requests.GetDocPermissions, resp *responses.GetDocPermissions) error {
	var err error
	implResp, err := s.Impl.GetDocPermissions(request)
	if err != nil {
		return err
	}

	// Overwrite the target address of resp to the target address of implResp.
	*resp = *implResp

	return nil
}

func (s *PdfiumRPCServer) GetFileVersion(request *requests.GetFileVersion, resp *responses.GetFileVersion) error {
	var err error
	implResp, err := s.Impl.GetFileVersion(request)
	if err != nil {
		return err
	}

	// Overwrite the target address of resp to the target address of implResp.
	*resp = *implResp

	return nil
}

func (s *PdfiumRPCServer) GetMetadata(request *requests.GetMetadata, resp *responses.GetMetadata) error {
	var err error
	implResp, err := s.Impl.GetMetadata(request)
	if err != nil {
		return err
	}

	// Overwrite the target address of resp to the target address of implResp.
	*resp = *implResp

	return nil
}

func (s *PdfiumRPCServer) GetPageCount(request *requests.GetPageCount, resp *responses.GetPageCount) error {
	var err error
	implResp, err := s.Impl.GetPageCount(request)
	if err != nil {
		return err
	}

	// Overwrite the target address of resp to the target address of implResp.
	*resp = *implResp

	return nil
}

func (s *PdfiumRPCServer) GetPageSize(request *requests.GetPageSize, resp *responses.GetPageSize) error {
	var err error
	implResp, err := s.Impl.GetPageSize(request)
	if err != nil {
		return err
	}

	// Overwrite the target address of resp to the target address of implResp.
	*resp = *implResp

	return nil
}

func (s *PdfiumRPCServer) GetPageSizeInPixels(request *requests.GetPageSizeInPixels, resp *responses.GetPageSizeInPixels) error {
	var err error
	implResp, err := s.Impl.GetPageSizeInPixels(request)
	if err != nil {
		return err
	}

	// Overwrite the target address of resp to the target address of implResp.
	*resp = *implResp

	return nil
}

func (s *PdfiumRPCServer) GetPageText(request *requests.GetPageText, resp *responses.GetPageText) error {
	var err error
	implResp, err := s.Impl.GetPageText(request)
	if err != nil {
		return err
	}

	// Overwrite the target address of resp to the target address of implResp.
	*resp = *implResp

	return nil
}

func (s *PdfiumRPCServer) GetPageTextStructured(request *requests.GetPageTextStructured, resp *responses.GetPageTextStructured) error {
	var err error
	implResp, err := s.Impl.GetPageTextStructured(request)
	if err != nil {
		return err
	}

	// Overwrite the target address of resp to the target address of implResp.
	*resp = *implResp

	return nil
}

func (s *PdfiumRPCServer) GetSecurityHandlerRevision(request *requests.GetSecurityHandlerRevision, resp *responses.GetSecurityHandlerRevision) error {
	var err error
	implResp, err := s.Impl.GetSecurityHandlerRevision(request)
	if err != nil {
		return err
	}

	// Overwrite the target address of resp to the target address of implResp.
	*resp = *implResp

	return nil
}

func (s *PdfiumRPCServer) RenderPageInDPI(request *requests.RenderPageInDPI, resp *responses.RenderPage) error {
	var err error
	implResp, err := s.Impl.RenderPageInDPI(request)
	if err != nil {
		return err
	}

	// Overwrite the target address of resp to the target address of implResp.
	*resp = *implResp

	return nil
}

func (s *PdfiumRPCServer) RenderPageInPixels(request *requests.RenderPageInPixels, resp *responses.RenderPage) error {
	var err error
	implResp, err := s.Impl.RenderPageInPixels(request)
	if err != nil {
		return err
	}

	// Overwrite the target address of resp to the target address of implResp.
	*resp = *implResp

	return nil
}

func (s *PdfiumRPCServer) RenderPagesInDPI(request *requests.RenderPagesInDPI, resp *responses.RenderPages) error {
	var err error
	implResp, err := s.Impl.RenderPagesInDPI(request)
	if err != nil {
		return err
	}

	// Overwrite the target address of resp to the target address of implResp.
	*resp = *implResp

	return nil
}

func (s *PdfiumRPCServer) RenderPagesInPixels(request *requests.RenderPagesInPixels, resp *responses.RenderPages) error {
	var err error
	implResp, err := s.Impl.RenderPagesInPixels(request)
	if err != nil {
		return err
	}

	// Overwrite the target address of resp to the target address of implResp.
	*resp = *implResp

	return nil
}

func (s *PdfiumRPCServer) RenderToFile(request *requests.RenderToFile, resp *responses.RenderToFile) error {
	var err error
	implResp, err := s.Impl.RenderToFile(request)
	if err != nil {
		return err
	}

	// Overwrite the target address of resp to the target address of implResp.
	*resp = *implResp

	return nil
}
