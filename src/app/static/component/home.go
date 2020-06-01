package component

// HomeComponent .
type HomeComponent struct{}

// Home .
var Home HomeComponent

// Render .
func (h HomeComponent) Render() string {
	return "<h2>Home Component</h2>"
}
