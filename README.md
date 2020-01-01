# pkg `rewrite`

pkg `rewrite` provides the http.ResponseWriter implementation and Replacer of the response.

## Installation

`go get -u github.com/Qs-F/rewrite`

## Example

```go
func main() {
  rw := &Rule{
    {
      Old: "world",
      New: "gopher",
    },
  }
  http.Handle("/", rw.MapFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello world")
  }))
  http.ListenAndServe(":8080", nil)
}
// Open localhost:8080, then you will get "hello gopher"
```

## LICENSE

MIT
