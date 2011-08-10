package go2d

type Panel struct {
	Container
	BackgroundColor
}

func NewPanel(x, y, width, height int) *Panel {
	panel := &Panel{}
	panel.Init(x, y, width, height)
	return panel
}

func (p *Panel) Draw(drawArea *Rect) {
	if p.backgroundColor != nil {
		realRect := NewRect(p.Rect().X+drawArea.X, p.Rect().Y+drawArea.Y, p.Rect().Width, p.Rect().Height)
		inRect := drawArea.Intersection(realRect)

		DrawFillRect(inRect, p.backgroundColor.R, p.backgroundColor.G, p.backgroundColor.B, 255)
	}
	
	p.Container.Draw(drawArea)
}