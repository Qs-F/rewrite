package rewrite

import (
	"net/http"
	"strings"
)

type Replace struct {
	Old string
	New string
}

type Rewrite struct {
	Replaces []*Replace
}

func (r *Rewrite) ToReplacer() *strings.Replacer {
	replaces := []string{}
	for _, rw := range r.Replaces {
		replaces = append(replaces, rw.Old, rw.New)
	}
	return strings.NewReplacer(replaces...)
}

func (rw *Rewrite) NewResponseWrite(w http.ResponseWriter) *ResponseWrite {
	return &ResponseWrite{w: w, rw: rw}
}

func (rw *Rewrite) Rewrite(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nw := rw.NewResponseWrite(w)
		handler.ServeHTTP(nw, r)
	})
}

func (rw *Rewrite) RewriteFunc(hf http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nw := rw.NewResponseWrite(w)
		hf.ServeHTTP(nw, r)
	})
}
