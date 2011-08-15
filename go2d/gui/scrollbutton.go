package go2d

type ScrollButton struct {
	Button
	
	OnScrollChangeListener
	
	boundaries *Rect
	startX, startY int
	mouseDownX, mouseDownY int
	scrolling bool
}

func NewScrollButton(x, y, width, height int, boundaries *Rect) *ScrollButton {
	scrollButton := &ScrollButton{}
	scrollButton.Init(x, y, width, height)
	scrollButton.boundaries = boundaries
	return scrollButton
}

func (s *ScrollButton) Boundaries() *Rect {
	return s.boundaries
}

func (s *ScrollButton) SetBoundaries(boundaries *Rect) {
	s.boundaries = boundaries
}

func (s *ScrollButton) MouseDown(x, y int) {
	s.startX = s.Rect().X
	s.startY = s.Rect().Y
	
	s.mouseDownX = x
	s.mouseDownY = y
	
	s.Button.MouseDown(x, y)
}

func (s *ScrollButton) UpdateScrollChangeListener() {
	if s.onScrollChange != nil {
		s.onScrollChange(s.ScrolledX(), s.ScrolledY())
	}
}

func (s *ScrollButton) MouseUp(x, y int) {
	if !s.scrolling && !s.mouseDown && s.boundaries != nil && s.boundaries.Contains(x, y) {
		newX := (s.boundaries.X + (x - s.boundaries.X)) - (s.Rect().Width/2)
		if newX >= s.boundaries.X && newX+s.Rect().Width <= s.boundaries.X+s.boundaries.Width {
			s.Rect().X = newX
		} 
		
		newY := (s.boundaries.Y + (y - s.boundaries.Y)) - (s.Rect().Height/2)
		if newY >= s.boundaries.Y && newY+s.Rect().Height <= s.boundaries.Y+s.boundaries.Height {
			s.Rect().Y = newY
		}
		
		s.UpdateScrollChangeListener()
	}
	s.scrolling = false
	s.Button.MouseUp(x, y)
}

func (s *ScrollButton) MouseMove(x, y int) {
	if s.boundaries != nil  && s.mouseDown {
		s.scrolling = true
		
		newX := s.startX + (x - s.mouseDownX)
		if newX >= s.boundaries.X && newX+s.Rect().Width <= s.boundaries.X+s.boundaries.Width {
			s.Rect().X = newX
		} else if newX < s.boundaries.X {
			s.Rect().X = s.boundaries.X
		} else if newX > s.boundaries.X+s.boundaries.Width {
			s.Rect().X = s.boundaries.Width - s.Rect().Width
		}
				
		newY := s.startY + (y - s.mouseDownY)
		if newY >= s.boundaries.Y && newY+s.Rect().Height <= s.boundaries.Y+s.boundaries.Height {
			s.Rect().Y = newY
		} else if newY < s.boundaries.Y {
			s.Rect().Y = s.boundaries.Y
		} else if newY > s.boundaries.Height {
			s.Rect().Y = s.boundaries.Height
		}
		
		s.UpdateScrollChangeListener()
	}
}

func (s *ScrollButton) ScrolledX() int {
	return s.Rect().X - s.boundaries.X
}

func (s *ScrollButton) ScrolledY() int {
	return s.Rect().Y - s.boundaries.Y
}