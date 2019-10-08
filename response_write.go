package rewrite

import (
	"net/http"
)

type ResponseWrite struct {
	w  http.ResponseWriter
	rw Rewriter
}

type Rewriter interface {
	Rewrite(p []byte) []byte
}

func NewResponseWrite(w http.ResponseWriter, rw Rewriter) *ResponseWrite {
	return &ResponseWrite{
		w:  w,
		rw: rw,
	}
}

func (w *ResponseWrite) Write(p []byte) (int, error) {
	return w.w.Write(w.rw.Rewrite(p))
}

func (w *ResponseWrite) Header() http.Header {
	return w.Header()
}

func (w *ResponseWrite) WriteHeader(statusCode int) {
	w.WriteHeader(statusCode)
}
