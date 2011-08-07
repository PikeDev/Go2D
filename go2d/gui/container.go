package gui

import . "go2d"

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

func (c *Container) Draw(drawArea *Rect) {
	for _, child := range c.Children() {
		if child.Visible() {
			child.Draw(c.rect) 
		}
	}
	c.DrawElement(drawArea)
}

func (c *Container) MouseDown(x, y int) {
	for _, child := range c.Children() {
		if child.Visible() {
			child.MouseDown(x, y)
		}
	}
}

func (c *Container) MouseUp(x, y int) {
	for _, child := range c.Children() {
		if child.Visible() {
			child.MouseUp(x, y)
		}
	}
}

func (c *Container) MouseMove(x, y int) {
	for _, child := range c.Children() {
		if child.Visible() {
			child.MouseMove(x, y)
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