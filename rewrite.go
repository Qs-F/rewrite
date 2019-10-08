package rewrite

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/Qs-F/bort"
)

// Replace is struct consisting of Old and New string
type Replace struct {
	Old string
	New string
}

// Rule is slice of Replace
type Rule []*Replace

// Replacer returns strings.Replacer by converting from Rule to strings.Replacer
func (rl *Rule) ToReplacer() *strings.Replacer {
	replaces := []string{}
	for _, rep := range *rl {
		replaces = append(replaces, rep.Old, rep.New)
	}
	return strings.NewReplacer(replaces...)
}

// Rewrite is implementation of Rewriter.
// If the content is binary, Rewrite returns original content.
func (rl *Rule) Rewrite(p []byte) []byte {
	if isbin, err := bort.IsBin(bytes.NewReader(p)); err != nil || isbin {
		return p
	}
	return []byte(rl.ToReplacer().Replace(string(p)))
}

// Map provides wrapper of http.Handler.
// Map returns http.Handler.
// By passing http.Handler, Map autonatically rewrite the content which handler write out.
func (rl *Rule) Map(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := NewResponseWrite(w, rl)
		h.ServeHTTP(rw, r)
	})
}

// MapFunc provides wrapper of http.HandlerFunc.
// MapFunc returns http.Handler.
// By passing http.Handler, Map autonatically rewrite the content which handler write out.
func (rl *Rule) MapFunc(hf http.HandlerFunc) http.Handler {
	return rl.Map(hf)
}
