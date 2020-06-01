package main

import (
	"app/static/component"
	"app/static/framework"
	"app/static/pkg"
)

func main() {
	// Starts the Oak framework
	framework.Start()

	// Starts our Router
	pkg.NewRouter()
	pkg.RegisterRoute("home", component.Home)
	pkg.RegisterRoute("about", component.About)

	// keeps our app running
	done := make(chan struct{}, 0)
	<-done
}
