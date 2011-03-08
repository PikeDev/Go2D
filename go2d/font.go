package go2d

import (
	"fmt"
	"sdl"
)

type Font struct {
	font    *sdl.Font
	fontmap map[uint32]map[uint16]*Image
	color   *sdl.Color
	size    int
	style   uint32
	alpha   uint8
}

//Load font from file and specify font size (each font-size needs a seperate font instance)
func NewFont(_file string, _size int) *Font {
	sdlfont := sdl.LoadFont(_file, _size)
	if sdlfont == nil {
		fmt.Printf("Go2D Error: Loading Font: %v", sdl.GetError())
	}
	f := &Font{font: sdlfont,
		fontmap: make(map[uint32]map[uint16]*Image),
		color:   &sdl.Color{255, 255, 255, 255},
		alpha:   255,
		size:    _size}
	addResource(f)
	f.build()
	return f
}

//Internal: release the font resource
func (f *Font) release() {
	f.font.Release()
}

//Set the color (RGB values 0-255)
func (f *Font) SetColor(_red uint8, _green uint8, _blue uint8) {
	f.color.R = _red
	f.color.G = _green
	f.color.B = _blue
}

//Set the opacity (0-255)
func (f *Font) SetAlpha(_alpha uint8) {
	f.alpha = _alpha
}

//Set the style (bold, italic, underlined)
func (f *Font) SetStyle(_bold bool, _italic bool, _underline bool) {
	var b, i, u uint32 = 0, 0, 0
	if _bold {
		b = 1
	}
	if _italic {
		i = 1
	}
	if _underline {
		u = 1
	}

	f.style = (b << 16) | (i << 8) | (u)

	_, present := f.fontmap[f.style]
	if !present {
		f.font.SetStyle(_bold, _italic, _underline)
		f.build()
	}
}

//Internal: build a map of all character textures
func (f *Font) build() {
	f.fontmap[f.style] = make(map[uint16]*Image)
	for c := 32; c <= 127; c++ {
		surface := f.font.RenderText_Blended(fmt.Sprintf("%c", c), sdl.Color{255, 255, 255, 255})
		img := newImageFromSurface(surface)

		f.fontmap[f.style][uint16(c)] = img
	}
}

//Draw a string on the given coordinates
func (f *Font) DrawText(_text string, _x, _y int) {
	prev_char := -1
	for c := 0; c < len(_text); c++ {
		_, _, _, _, advance := f.font.GetMetrics(uint16(_text[c]))

		if prev_char != -1 {
			kerning := f.font.GetKerning(prev_char, int(_text[c]))
			_x += kerning

			prev_char = int(_text[c])
		}

		img := f.fontmap[f.style][uint16(_text[c])]
		if img != nil {
			img.SetColorMod(f.color.R, f.color.G, f.color.B)
			img.SetAlphaMod(f.alpha)
			img.Draw(_x, _y)
			_x += advance
		}
	}
}

//Draw a string on the given coordinates, within a rect (text outside of the rect will be omitted)
func (f *Font) DrawTextInRect(_text string, _x int, _y int, _rect *Rect) {
	prev_char := -1
	for c := 0; c < len(_text); c++ {
		_, _, _, _, advance := f.font.GetMetrics(uint16(_text[c]))

		if prev_char != -1 {
			kerning := f.font.GetKerning(prev_char, int(_text[c]))
			_x += kerning

			prev_char = int(_text[c])
		}

		img := f.fontmap[f.style][uint16(_text[c])]
		if img != nil {
			img.SetColorMod(f.color.R, f.color.G, f.color.B)
			img.SetAlphaMod(f.alpha)
			img.DrawInRect(_x, _y, _rect)
			_x += advance
		}
	}
}

//Get the pixel width of the given string
func (f *Font) GetStringWidth(_text string) int {
	w := 0
	prev_char := -1
	for c := 0; c < len(_text); c++ {
		_, _, _, _, advance := f.font.GetMetrics(uint16(_text[c]))

		if prev_char != -1 {
			kerning := f.font.GetKerning(prev_char, int(_text[c]))
			w += kerning

			prev_char = int(_text[c])
		}

		w += advance
	}
	return w
}

//Get the pixel height of the font
func (f *Font) GetStringHeight() int {
	return f.font.GetHeight() - 4
}
