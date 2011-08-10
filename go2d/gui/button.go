package go2d

import "sdl"

type Button struct {
	Element
	
	//Properties
	BackgroundColor
	ElementImage
	TextElement
	
	//Listeners
	OnClickListener
	
	//Members
	caption string
	mouseDown bool
	hover bool
}

func NewButton(x, y, width, height int, caption string) *Button {
	button := &Button{}
	button.Init(x, y, width, height)
	button.caption = caption
	button.SetFont(g_game.guiManager.defaultFont)
	button.fontColor = &sdl.Color{uint8(255), uint8(255), uint8(255), uint8(255)}
	return button
}

func (b *Button) Caption() string {
	return b.caption
}

func (b *Button) SetCaption(caption string) {
	b.caption = caption
}

func (b *Button) Draw(drawArea *Rect) {
	realRect := NewRect(b.Rect().X+drawArea.X, b.Rect().Y+drawArea.Y, b.Rect().Width, b.Rect().Height)
	inRect := drawArea.Intersection(realRect)
	
	if b.backgroundColor != nil {
		DrawFillRect(inRect, b.backgroundColor.R, b.backgroundColor.G, b.backgroundColor.B, 255)
	}
	
	var drawable *Image
	switch {
		case b.mouseDown && b.mouseDownImage != nil:
			drawable = b.mouseDownImage
			
		case b.hover && b.hoverImage != nil:
			drawable = b.hoverImage
			
		default:
			drawable = b.image
	}
	if drawable != nil {
		drawable.DrawRectInRect(b.Rect(), drawArea)
	}
	
	if b.caption != "" &&  b.font != nil {
		captionX := b.Rect().X+((b.Rect().Width/2)-(b.font.GetStringWidth(b.caption)/2));
		captionY := b.Rect().Y+((b.Rect().Height/2)-(b.font.GetStringHeight()/2));
		b.font.SetStyle(b.bold, b.italic, b.underlined)
		b.font.SetColor(b.fontColor.R, b.fontColor.G, b.fontColor.B)
		b.font.DrawTextInRect(b.caption, drawArea.X + captionX, drawArea.Y + captionY, drawArea)
	}
}

func (b *Button) MouseMove(x, y int) {
	if b.Rect().Contains(x, y) {
		b.hover = true
	} else {
		b.hover = false
	}
}

func (b *Button) MouseDown(x, y int) {
	if b.Rect().Contains(x, y) {
		b.mouseDown = true
	}
}

func (b *Button) MouseUp(x, y int) {
	if b.Rect().Contains(x, y) {
		if b.mouseDown && b.onClick != nil {
			b.onClick(x, y)
		}
	}
	b.mouseDown = false
}
