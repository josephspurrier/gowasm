package framework

import "syscall/js"

// RegisterFunction .
func RegisterFunction(funcName string, myfunc func(value js.Value, args []js.Value) interface{}) {
	js.Global().Set(funcName, js.FuncOf(myfunc))
}

// Start .
func Start() {
	println("GoWASM Framework Initialized")
}
