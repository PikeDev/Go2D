package go2d

import "sdl"

type Label struct {
	Element
	TextElement

	caption string
}
 
func NewLabel(x, y int, caption string) *Label {
	label := &Label{}
	
	label.Init(x, y, 0, 0)
	
	label.caption = caption
	
	label.SetFont(g_game.guiManager.defaultFont)
	label.fontColor = &sdl.Color{uint8(255), uint8(255), uint8(255), uint8(255)}

	return label
}

func (l *Label) SetCaption(caption string) {
	l.caption = caption
}

func (l *Label) Caption() string {
	return l.caption
}

func (l *Label) SetFont(font *Font) {
	l.font = font
	width, height := 0, 0
	if l.font != nil {
		width = l.font.GetStringWidth(l.caption)
		height = l.font.GetStringHeight()+2
	}
	l.Rect().Width = width
	l.Rect().Height = height
}

func (l *Label) Draw(drawArea *Rect) {
	realRect := NewRect(l.Rect().X+drawArea.X, l.Rect().Y+drawArea.Y, l.Rect().Width, l.Rect().Height)
	inRect := drawArea.Intersection(realRect)
	
	if l.font != nil {
		l.font.SetStyle(l.bold, l.italic, l.underlined)
		l.font.SetColor(l.fontColor.R, l.fontColor.G, l.fontColor.B)
		l.font.DrawTextInRect(l.caption, drawArea.X + l.Rect().X, drawArea.Y + l.Rect().Y, inRect)
	}
}