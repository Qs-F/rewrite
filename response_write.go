package rewrite

import (
	"net/http"
)

// ResponseWrite is http.ResponseWriter interface implementation.
type ResponseWrite struct {
	w  http.ResponseWriter
	rw Rewriter
}

// Rewriter interface is used to rewrite content based on original content.
type Rewriter interface {
	Rewrite(p []byte) []byte
}

func NewResponseWrite(w http.ResponseWriter, rw Rewriter) *ResponseWrite {
	return &ResponseWrite{
		w:  w,
		rw: rw,
	}
}

// Write writes content which is returned from Rewriter to p.
func (w *ResponseWrite) Write(p []byte) (int, error) {
	return w.w.Write(w.rw.Rewrite(p))
}

func (w *ResponseWrite) Header() http.Header {
	return w.w.Header()
}

func (w *ResponseWrite) WriteHeader(statusCode int) {
	// http.FileServer seems automatically set Content-Length at called time.
	// But it is needed to rewrite the content after that, depending Go's functionality that if there is no Cotent-Length it will be automatically solved is good idea.
	w.w.Header().Del("Content-Length")
	w.w.WriteHeader(statusCode)
}
