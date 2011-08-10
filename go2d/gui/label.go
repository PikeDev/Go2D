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
	
	label.SetFont(g_game.guiManager.defaultFont)
	label.foregroundColor = &sdl.Color{uint8(255), uint8(255), uint8(255), uint8(255)}
	
	label.caption = caption
	
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
		height = l.font.GetStringHeight()
	}
	l.Rect().Width = width
	l.Rect().Height = height
}

func (l *Label) Draw(drawArea *Rect) {
	if l.font != nil {
		l.font.SetStyle(l.bold, l.italic, l.underlined)
		l.font.SetColor(l.foregroundColor.R, l.foregroundColor.G, l.foregroundColor.B)
		l.font.DrawTextInRect(l.caption, drawArea.X + l.Rect().X, drawArea.Y + l.Rect().Y, drawArea)
	}
}