package go2d

type Window struct {
	Container

	title string

	popups []*Window
}

func NewWindow(x, y, width, height int) *Window {
	window := &Window{}
	window.Init(x, y, width, height)
	return window
}