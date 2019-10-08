package rewrite

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type T struct {
	Handler http.Handler
	Rule    *Rule
	Expect  string
}

func TestMap(t *testing.T) {
	tests := []T{
		{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello Japan")
			}),
			Rule: &Rule{
				{
					Old: "Japan",
					New: "World",
				},
			},
			Expect: "Hello World",
		},
		{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello Japan")
			}),
			Rule: &Rule{
				{
					Old: "World",
					New: "Gopher",
				},
			},
			Expect: "Hello Japan",
		},
		{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello Japan")
			}),
			Rule: &Rule{
				{
					Old: "Japan",
					New: "日本",
				},
			},
			Expect: "Hello 日本",
		},
		{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("gopher", "hello")
				fmt.Fprintf(w, "Hello World")
			}),
			Rule: &Rule{
				{
					Old: "World",
					New: "Gopher",
				},
			},
			Expect: "Hello 日本",
		},
		{
			Handler: http.FileServer(http.Dir("./_testdata/helloworld")),
			Rule: &Rule{
				{
					Old: "world",
					New: "gopher",
				},
			},
			Expect: "hello gopher",
		},
	}

	for i, test := range tests {
		t.Log("testcase: ", i)
		srv := httptest.NewServer(test.Rule.Map(test.Handler))
		resp, err := http.Get(srv.URL)
		if err != nil {
			t.Error(err)
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			t.Error(err)
			continue
		}

		if !bytes.Equal(body, []byte(test.Expect)) {
			t.Error(fmt.Sprintf("test %d: expect: %s but get: %s\n", i, test.Expect, string(body)))
			continue
		}
		srv.Close()
	}
}
