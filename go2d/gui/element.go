package go2d

type IElement interface {
	Draw(drawArea *Rect)
	MouseDown(x, y int)
	MouseUp(x, y int)
	MouseMove(x, y int)
	MouseScroll(direction int)
	KeyDown(button int)
	KeyUp(button int)
	TextInput(character uint8)
	Rect() *Rect
	Visible() bool
	Parent() IElement
	SetParent(element IElement)
}

type Element struct {
	rect *Rect
	visible bool
	parent IElement
}

func (e *Element) Init(x, y, width, height int) {
	e.rect = NewRect(x, y, width, height)
	e.visible = true
}

func (e *Element) Draw(drawArea *Rect) {
}

func (e *Element) MouseDown(x, y int) {
}

func (e *Element) MouseUp(x, y int) {
}

func (e *Element) MouseMove(x, y int) {
}

func (e *Element) MouseScroll(direction int) {
}

func (e *Element) KeyDown(button int) {
}

func (e *Element) KeyUp(button int) {
}

func (e *Element) TextInput(character uint8) {
}

func (e *Element) Visible() bool {
	return e.visible
}

func (e *Element) Rect() *Rect {
	return e.rect
}

func (e *Element) Parent() IElement {
	return e.parent
}

func (e *Element) SetParent(element IElement) {
	e.parent = element
}