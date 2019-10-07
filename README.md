# pkg `rewrite`

pkg `rewrite` provides the http.ResponseWriter implementation and Replacer of the response.

```go
func main() {
	rw := &Rewrite{
		Replaces: []*Replace{
			{
				Old: "world",
				New: "google",
			},
		},
	}
	http.Handle("/", rw.RewriteFunc(handler))
	http.ListenAndServe(":8080", nil)
}
```
