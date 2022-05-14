package rewrite

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"testing"
)

type T struct {
	Handler  http.Handler
	FilePath string
	Rule     *Rule
	Expect   string
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
			Expect: "Hello Gopher",
		},
		{
			Handler:  http.FileServer(http.Dir("./testdata")),
			FilePath: "/helloworld",
			Rule: &Rule{
				{
					Old: "world",
					New: "gopher",
				},
			},
			Expect: "hello gopher\n",
		},
		{
			Handler:  http.FileServer(http.Dir("./testdata")),
			FilePath: "/notawesomefile",
			Rule: &Rule{
				{
					Old: "404",
					New: "four-zero-four",
				},
			},
			Expect: "four-zero-four page not found\n",
		},
	}

	for i, test := range tests {
		srv := httptest.NewServer(test.Rule.Map(test.Handler))
		url, _ := url.Parse(srv.URL)
		url.Path = path.Join(url.Path, test.FilePath)
		t.Log(url)
		resp, err := http.Get(url.String())
		if err != nil {
			t.Error(err)
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
			continue
		}

		if !bytes.Equal(body, []byte(test.Expect)) {
			t.Error(fmt.Sprintf("test %d: expect: %s but get: %s\n", i, test.Expect, string(body)))
			continue
		}

		t.Log(fmt.Sprintf("test %d: get: %s\n", i, string(body)))
		resp.Body.Close()
		srv.Close()
	}
}
