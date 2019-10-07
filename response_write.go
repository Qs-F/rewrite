package rewrite

import (
	"net/http"

	"github.com/Qs-F/bort"
)

type ResponseWrite struct {
	w  http.ResponseWriter
	rw *Rewrite
}

func (w *ResponseWrite) Write(p []byte) (int, error) {
	if !bort.IsBin(p) {
		return w.w.Write([]byte(w.rw.ToReplacer().Replace(string(p))))
	}
	return w.w.Write(p)
}

func (w *ResponseWrite) Header() http.Header {
	return w.Header()
}

func (w *ResponseWrite) WriteHeader(statusCode int) {
	w.WriteHeader(statusCode)
}
