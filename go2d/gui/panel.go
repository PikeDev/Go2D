package go2d

import "sdl"

type Panel struct {
	Container
	
	backgroundColor *sdl.Color
}

func NewPanel(x, y, width, height int) *Panel {
	panel := &Panel{}
	panel.Init(x, y, width, height)
	return panel
}

func (p *Panel) SetBackgroundColor(red, green, blue int) {
	if p.backgroundColor == nil {
		p.backgroundColor = &sdl.Color{uint8(red), uint8(green), uint8(blue), uint8(255)}
	} else {
		p.backgroundColor.R = uint8(red)
		p.backgroundColor.G = uint8(green)
		p.backgroundColor.B = uint8(blue)
	}
}

func (p *Panel) Draw(drawArea *Rect) {
	if p.backgroundColor != nil {
		DrawFillRect(p.Rect(), p.backgroundColor.R, p.backgroundColor.G, p.backgroundColor.B, 255)
	}
	
	p.Container.Draw(drawArea)
}