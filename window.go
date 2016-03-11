package gonsole

// Window is the top-level struct in gonsole library.
type Window struct {
	BaseWindow
}

// NewWindow creates a new window for later display
func NewWindow(app *App, id string) *Window {
	win := &Window{}
	win.Init(app, id)
	return win
}
