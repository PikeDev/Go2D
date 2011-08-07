package gui

import . "go2d"

type IElement interface {
	Draw(drawArea *Rect)
	DrawElement(drawArea *Rect)
	MouseDown(x, y int)
	MouseUp(x, y int)
	MouseMove(x, y int)
	MouseScroll(direction int)
	KeyDown(button int)
	KeyUp(button int)
	TextInput(character uint8)
	Rect() *Rect
	Visible() bool
}

type Element struct {
	rect *Rect
	visible bool
}

func (e *Element) Init(x, y, width, height int) {
	e.rect = NewRect(x, y, width, height)
}

func (e *Element) Draw(drawArea *Rect) {
	e.DrawElement(drawArea)
}

func (e *Element) DrawElement(drawArea *Rect) {
	//If there's nothing to draw
}

func (e *Element) Visible() bool {
	return e.visible
}

func (e *Element) Rect() *Rect {
	return e.rect
}
