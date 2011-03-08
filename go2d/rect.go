package go2d

import "sdl"

type Rect struct {
	X, Y, Width, Height int
}

func NewRect(_x int, _y int, _width int, _height int) *Rect {
	return &Rect{X: _x, Y: _y, Width: _width, Height: _height}
}

func NewRectFrom(_rect *Rect) *Rect {
	return NewRect(_rect.X, _rect.Y, _rect.Width, _rect.Height)
}

func (rect *Rect) Equals(_rect *Rect) bool {
	return ((rect.X == _rect.X) && (rect.Y == _rect.Y) && (rect.Width == _rect.Width) && (rect.Height == _rect.Height))
}

func (rect *Rect) Contains(_x int, _y int) bool {
	return (_x >= rect.X && _x <= rect.X+rect.Width && _y >= rect.Y && _y <= rect.Y+rect.Height)
}

func (rect *Rect) ContainsRect(_rect *Rect) bool {
	return (rect.Contains(_rect.X, _rect.Y) &&
		rect.Contains(_rect.X+_rect.Width, _rect.Y) &&
		rect.Contains(_rect.X, _rect.Y+_rect.Height) &&
		rect.Contains(_rect.X+_rect.Width, _rect.Y+_rect.Height))
}

func (rect *Rect) Intersects(_rect *Rect) bool {
	return !(rect.X > _rect.X+_rect.Width || _rect.X > rect.X+rect.Width ||
		rect.Y > _rect.Y+_rect.Height || _rect.Y > rect.Y+rect.Height)
}

func (rect *Rect) Intersection(_rect *Rect) *Rect {
	if rect.Intersects(_rect) {
		tempX := _rect.X
		if rect.X > _rect.X {
			tempX = rect.X
		}

		tempY := _rect.Y
		if rect.Y > _rect.Y {
			tempY = rect.Y
		}

		tempW := _rect.X + _rect.Width
		if rect.X+rect.Width < _rect.X+_rect.Width {
			tempW = rect.X + rect.Width
		}

		tempH := _rect.Y + _rect.Height
		if rect.Y+rect.Height < _rect.Y+_rect.Height {
			tempH = rect.Y + rect.Height
		}

		tempW -= tempX
		tempH -= tempY

		return NewRect(tempX, tempY, tempW, tempH)
	}
	return NewRect(0, 0, 0, 0)
}

//Internal: convert Rect to sdl.Rect
func (rect *Rect) toSDL() *sdl.Rect {
	return &sdl.Rect{int32(rect.X), int32(rect.Y), int32(rect.Width), int32(rect.Height)}
}
