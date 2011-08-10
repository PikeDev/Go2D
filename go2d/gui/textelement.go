package go2d
 
import "sdl"
 
type TextElement struct {
	font *Font
	fontColor *sdl.Color
	bold, italic, underlined bool
}

func (t *TextElement) SetFont(font *Font) {
	t.font = font
}

func (t *TextElement) Font() *Font {
	return t.font
}

func (t *TextElement) SetFontColor(red, green, blue int) {
	t.fontColor.R = uint8(red)
	t.fontColor.G = uint8(green)
	t.fontColor.B = uint8(blue)
}

func (t *TextElement) FontColor() *sdl.Color {
	return t.fontColor
}

func (t *TextElement) SetFontStyle(bold, italic, underlined bool) {
	t.bold = bold
	t.italic = italic
	t.underlined = underlined
}

func (t *TextElement) Bold() bool {
	return t.bold
}

func (t *TextElement) Italic() bool {
	return t.italic
}

func (t *TextElement) Underlined() bool {
	return t.underlined
}