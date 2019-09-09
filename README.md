# snippetgo: Generate packed go file with snippet code marked in source code

It is used for embed code blocks into your documentations. because code blocks are in your compilable source code, So you will never generate invalid code blocks in your docs.

Since I will write docs in go source code, So right now it generate to go source code as global variables.

Install

```
$ go get -u -v github.com/sunfmin/snippetgo
```

Mark your source code with these
```go
// @snippet_begin(VariableName)
func Hello() {
    fmt.Println("Hello")
}
// @snippet_end
```

Then run

```bash
$ snippetgo -pkg=mypkg > ./mypkg/snippets_generated.go
```
