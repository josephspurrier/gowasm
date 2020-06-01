package component

import (
	"fmt"
	"syscall/js"

	"app/static/framework"
)

// AboutComponent .
type AboutComponent struct{}

// About .
var About AboutComponent

func init() {
	framework.RegisterFunction("coolFunc", CoolFunc)
}

// Render .
func (a AboutComponent) Render() string {
	return `
<div>
	<h2>About Component Actually Works</h2>
	<button onClick="coolFunc();">Cool Func</button>
</div>`
}

// CoolFunc .
func CoolFunc(value js.Value, args []js.Value) interface{} {
	fmt.Println("does stuff")
	return nil
}
