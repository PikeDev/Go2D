package go2d

import (
	"sdl"
)

type Image struct {
	Width, Height uint16
	blendmode     int

	surface *sdl.Surface
	texture *sdl.Texture
}

//Load image from file
func NewImage(_file string) *Image {
	image := &Image{blendmode: sdl.BLENDMODE_BLEND}
	addResource(image)

	image.surface = sdl.LoadImage(GetPath() + _file)

	image.reload()
	return image
}

//Internal: load image from sdl surface
func newImageFromSurface(_surface *sdl.Surface) *Image {
	image := &Image{surface: _surface,
		blendmode: sdl.BLENDMODE_BLEND}
	addResource(image)
	image.reload()
	return image
}

//Internal: create texture to render
func (image *Image) reload() {
	image.texture = image.surface.CreateTexture(g_game.renderer)
	image.Width = uint16(image.texture.W)
	image.Height = uint16(image.texture.H)
}

//Internal: release the image's resources
func (image *Image) release() {
	if image.surface != nil {
		image.surface.Release()
	}
	if image.texture != nil {
		image.texture.Release()
	}
}

//Set the blendmode (BLENDMODE_NONE, BLENDMODE_BLEND, BLENDMODE_ADD, BLENDMODE_MOD)
func (image *Image) SetBlendMode(_blendmode int) {
	image.blendmode = _blendmode
	image.texture.SetBlendMode(_blendmode)
}

//Set the color modifier (used for color blending)
func (image *Image) SetColorMod(_red uint8, _green uint8, _blue uint8) {
	image.texture.SetColorMod(_red, _green, _blue)
}

//Set the alpha value (transparency (0 for transparent - 255 for opaque)
func (image *Image) SetAlphaMod(_alpha uint8) {
	image.texture.SetAlpha(_alpha)
}

//Internal: render image to buffer
func (image *Image) render(_src *sdl.Rect, _dst *sdl.Rect) {
	image.texture.SetBlendMode(image.blendmode)
	image.texture.RenderCopy(g_game.renderer, _src, _dst)
}

//Draw image at given coordinates
func (image *Image) Draw(_x, _y int) {
	src := &sdl.Rect{0, 0, int32(image.Width), int32(image.Height)}
	dst := &sdl.Rect{int32(_x), int32(_y), int32(image.Width), int32(image.Height)}

	image.render(src, dst)
}

//Draw the image as a rect (e.g. can be stretched) at the given coordinates
func (image *Image) DrawRect(_rect *Rect) {
	src := &sdl.Rect{0, 0, int32(image.Width), int32(image.Height)}

	image.render(src, _rect.toSDL())
}

//Draw a part of the image on the given coordinates
func (image *Image) DrawClip(_x, _y int, _clip *Rect) {
	dst := &sdl.Rect{0, 0, int32(image.Width), int32(image.Height)}

	image.render(_clip.toSDL(), dst)
}

//Draw a part of the image as a rect (e.g. can be stretched) at the given coordinates
func (image *Image) DrawRectClip(_rect *Rect, _clip *Rect) {
	image.render(_clip.toSDL(), _rect.toSDL())
}

//Draw the image within a rect (cut off pieces that fall outside of the rect)
func (image *Image) DrawInRect(_x, _y int, _inrect *Rect) {
	imgRect := NewRect(_x, _y, int(image.Width), int(image.Height))
	inRect := _inrect.Intersection(imgRect)
	dstRect := NewRect(inRect.X, inRect.Y, inRect.Width, inRect.Height)
	inRect.X -= _x
	inRect.Y -= _y

	image.render(inRect.toSDL(), dstRect.toSDL())
}

//Draw the image as a rect (e.g. can be stretched) within a rect (cut off pieces that fall outside of the rect)
func (image *Image) DrawRectInRect(_rect *Rect, _inrect *Rect) {
	inRect := _inrect.Intersection(_rect)
	dstRect := NewRect(inRect.X, inRect.Y, inRect.Width, inRect.Height)
	inRect.X -= _rect.X
	inRect.Y -= _rect.Y

	image.render(inRect.toSDL(), dstRect.toSDL())
}
