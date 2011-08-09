package go2d

type IContainer interface {
	Children() []IElement
}

type Container struct {
	Element
	children []IElement
}

func (c *Container) Init(x, y, width, height int) {
	c.Element.Init(x, y, width, height)
	c.children = make([]IElement, 0)
}

func (c *Container) Children() []IElement {
	return c.children
}

func (c *Container) AddChild(element IElement) {
	element.SetParent(element)
	c.children = append(c.children, element)
}

func (c *Container) RemoveChild(element IElement) {
	for index, child := range c.children {
		if child == element {
			h := c.children[0:index]
			l := c.children[index+1:]
			c.children = append(h, l...)
		}
	}
}

func (c *Container) Draw(drawArea *Rect) {
	childDrawArea := NewRectFrom(drawArea)
	childDrawArea.X += c.Rect().X
	childDrawArea.Y += c.Rect().Y
	
	for _, child := range c.Children() {
		if child.Visible() {
			child.Draw(childDrawArea) 
		}
	}
}

func (c *Container) MouseDown(x, y int) {
	for _, child := range c.Children() {
		if child.Visible() {
			child.MouseDown(x - c.Rect().X, y - c.Rect().Y)
		}
	}
}

func (c *Container) MouseUp(x, y int) {
	for _, child := range c.Children() {
		if child.Visible() {
			child.MouseUp(x - c.Rect().X, y - c.Rect().Y)
		}
	}
}

func (c *Container) MouseMove(x, y int) {
	for _, child := range c.Children() {
		if child.Visible() {
			child.MouseMove(x - c.Rect().X, y - c.Rect().Y)
		}
	}
}

func (c *Container) MouseScroll(direction int) {
	for _, child := range c.Children() {
		if child.Visible() {
			child.MouseScroll(direction)
		}
	}
}

func (c *Container) KeyDown(button int) {
	for _, child := range c.Children() {
		if child.Visible() {
			child.KeyDown(button)
		}
	}
}

func (c *Container) KeyUp(button int) {
	for _, child := range c.Children() {
		if child.Visible() {
			child.KeyUp(button)
		}
	}
}

func (c *Container) TextInput(character uint8) {
	for _, child := range c.Children() {
		if child.Visible() {
			child.TextInput(character)
		}
	}
}