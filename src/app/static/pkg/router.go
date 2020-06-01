package pkg

import (
	"fmt"
	"syscall/js"

	"app/static/component"
)

// Router .
type Router struct {
	Routes map[string]component.Component
}

var router Router

func init() {
	router.Routes = make(map[string]component.Component)
}

// NewRouter .
func NewRouter() {
	js.Global().Set("Link", js.FuncOf(Link))
	js.Global().Get("document").Call("getElementById", "view").Set("innerHTML", "")
}

// RegisterRoute .
func RegisterRoute(path string, component component.Component) {
	router.Routes[path] = component
}

// Link .
func Link(value js.Value, args []js.Value) interface{} {
	fmt.Println("Link clicked:", args[0].String())

	comp := router.Routes[args[0].String()]
	html := comp.Render()

	js.Global().Get("document").Call("getElementById", "view").Set("innerHTML", html)
	return nil
}
