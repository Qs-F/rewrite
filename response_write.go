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
	return w.Header()
}

func (w *ResponseWrite) WriteHeader(statusCode int) {
	w.WriteHeader(statusCode)
}
