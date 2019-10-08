package rewrite

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/Qs-F/bort"
)

type Replace struct {
	Old string
	New string
}

type Rule []*Replace

func (rl *Rule) ToReplacer() *strings.Replacer {
	replaces := []string{}
	for _, rep := range *rl {
		replaces = append(replaces, rep.Old, rep.New)
	}
	return strings.NewReplacer(replaces...)
}

func (rl *Rule) Rewrite(p []byte) []byte {
	if isbin, err := bort.IsBin(bytes.NewReader(p)); err != nil || isbin {
		return p
	}
	return []byte(rl.ToReplacer().Replace(string(p)))
}

func (rl *Rule) Map(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := NewResponseWrite(w, rl)
		h.ServeHTTP(rw, r)
	})
}

func (rl *Rule) MapFunc(hf http.HandlerFunc) http.Handler {
	return rl.Map(hf)
}
