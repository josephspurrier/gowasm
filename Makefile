# This Makefile is an easy way to run common operations.
# Execute commands this:
# * make command-here
#
# Tip: Each command is run on its own line so you can't CD unless you
# connect commands together using operators. See examples:
# A; B    # Run A and then B, regardless of success of A
# A && B  # Run B if and only if A succeeded
# A || B  # Run B if and only if A failed
# A &     # Run A in background.
# Source: https://askubuntu.com/a/539293
#
# Tip: Use $(shell app param) syntax when expanding a shell return value.

.PHONY: files
files:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" src/app/static/
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.html" src/app/static/index.html

.PHONY: build
build:
	cd src/app/static && GOOS=js GOARCH=wasm go build -o test.wasm main.go

.PHONY: serve
serve: build
	cd src/app/webserver && go build -o webserver main.go && ./webserver