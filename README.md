# WASM and Go

This project is a sample Go application that compiles to a WebAssembly (WASM) application for use in a web browser. This project is based on this article:
https://tutorialedge.net/golang/writing-frontend-web-framework-webassembly-go/.

This project utilizes a `Makefile`.

```bash
# Compiles a WASM application called test.wasm.
make build

# Calls 'make serve' and then runs a web server on localhost:80.
make serve
```

You can how open your browser to: http://localhost.

References:

- https://dev.to/enbis/generate-and-run-webassembly-code-using-go-4fbp
- https://blog.gopheracademy.com/advent-2018/go-in-the-browser/
- https://tutorialedge.net/golang/writing-frontend-web-framework-webassembly-go/
- https://github.com/golang/go/wiki/WebAssembly
- https://golang.org/pkg/syscall/js/