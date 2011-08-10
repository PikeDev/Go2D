package go2d

import "sdl"

type BackgroundColor struct {
 	backgroundColor *sdl.Color
}

func (b *BackgroundColor) BackgroundColor() *sdl.Color {
	return b.backgroundColor
}

func (b *BackgroundColor) SetBackgroundColor(red, green, blue int) {
	if b.backgroundColor == nil {
		b.backgroundColor = &sdl.Color{uint8(red), uint8(green), uint8(blue), uint8(255)}
	} else {
		b.backgroundColor.R = uint8(red)
		b.backgroundColor.G = uint8(green)
		b.backgroundColor.B = uint8(blue)
	}
}

type ElementImage struct {
	image *Image
	hoverImage *Image
	mouseDownImage *Image
}

func (e *ElementImage) Image() *Image {
	return e.image
}

func (e *ElementImage) SetImage(image *Image) {
	e.image = image
}

func (e *ElementImage) HoverImage() *Image {
	return e.hoverImage
}

func (e *ElementImage) SetHoverImage(image *Image) {
	e.hoverImage = image
}

func (e *ElementImage) MouseDownImage() *Image {
	return e.mouseDownImage
}

func (e *ElementImage) SetMouseDownImage(image *Image) {
	e.mouseDownImage = image
}