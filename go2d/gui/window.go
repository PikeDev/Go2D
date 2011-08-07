package gui

type Window struct {
	Container

	title string
	root IElement
	
	popups []*Window
}

func NewWindow(x, y, width, height int) *Window {
	window := &Window{}
	window.Init(x, y, width, height)
	return window
}