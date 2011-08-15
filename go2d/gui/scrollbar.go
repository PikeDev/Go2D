package go2d

const (
	SCROLLBAR_VERTICAL = iota
	SCROLLBAR_HORIZONTAL 
)

type Scrollbar struct {
	Container
	
	//Properties
	Value
	
	//Listeners
	OnValueChangeListener
	
	//Members
	orientation int
	pnlBackground *Panel
	btnLeftTop, btnRightDown *Button
	scbScroller *ScrollButton
	minValue, maxValue int
}

func NewScrollbar(x, y, width, height, orientation int) *Scrollbar {
	scrollbar := &Scrollbar{}
	scrollbar.Init(x, y, width, height)
	scrollbar.orientation = orientation
	
	scrollbar.value = 0
	scrollbar.minValue = 0
	scrollbar.maxValue = 100
	
	scrollbar.pnlBackground = NewPanel(0, 0, width, height)
	scrollbar.pnlBackground.SetBackgroundColor(80, 80, 80)
	scrollbar.AddChild(scrollbar.pnlBackground)
	
	if orientation == SCROLLBAR_VERTICAL {
		scrollbar.btnLeftTop = NewButton(0, 0, width, width, "")
		scrollbar.btnRightDown = NewButton(0, height-width, width, width, "")
		scrollbar.scbScroller = NewScrollButton(0, width, width, width, NewRect(0, width, width, height-(2*width)))
	} else {
		scrollbar.btnLeftTop = NewButton(0, 0, height, height, "")
		scrollbar.btnRightDown = NewButton(width-height, 0, height, height, "")	
		scrollbar.scbScroller = NewScrollButton(height, 0, height, height, NewRect(height, 0, height, width-(2*height)))
	}
	scrollbar.btnLeftTop.SetBackgroundColor(120,120,120)
	scrollbar.btnRightDown.SetBackgroundColor(120,120,120)
	scrollbar.scbScroller.SetBackgroundColor(100,100,100)
	
	scrollbar.AddChild(scrollbar.btnLeftTop)
	scrollbar.AddChild(scrollbar.btnRightDown)
	scrollbar.AddChild(scrollbar.scbScroller)
	
	scrollbar.scbScroller.onScrollChange = func(scrolledX, scrolledY int) {
		scrollbar.ScrollButtonChanged(scrolledX, scrolledY)
	}
	
	scrollbar.btnLeftTop.SetOnClickListener(func(x, y int) {
		if scrollbar.value-1 >= scrollbar.minValue {
			scrollbar.value--
			scrollbar.UpdateScrollerPos()
			
			if scrollbar.onValueChange != nil {
				scrollbar.onValueChange(scrollbar.value)
			}
		}
	})
	
	scrollbar.btnRightDown.SetOnClickListener(func(x, y int) {
		if scrollbar.value+1 <= scrollbar.maxValue {
			scrollbar.value++
			scrollbar.UpdateScrollerPos()
			
			if scrollbar.onValueChange != nil {
				scrollbar.onValueChange(scrollbar.value)
			}
		}
	})
	
	return scrollbar
}

func (s *Scrollbar) ScrollButtonChanged(scrolledX, scrolledY int) {
	s.value = s.minValue + int((float32(scrolledY) / float32(s.ScrollAreaSize())) * float32(s.maxValue))

	if s.onValueChange != nil {
		s.onValueChange(s.value)
	}
}

func (s *Scrollbar) ButtonLeftTop() *Button {
	return s.btnLeftTop
}

func (s *Scrollbar) ButtonRightDown() *Button {
	return s.btnRightDown
}

func (s *Scrollbar) Scroller() *ScrollButton {
	return s.scbScroller
}

func (s *Scrollbar) MinValue() int {
	return s.minValue
}

func (s *Scrollbar) SetMinValue(minValue int) {
	s.minValue = minValue
}

func (s *Scrollbar) MaxValue() int {
	return s.maxValue
}

func (s *Scrollbar) SetMaxValue(maxValue int) {
	s.maxValue = maxValue
}

func (s *Scrollbar) ScrollAreaSize() int {
	if s.orientation == SCROLLBAR_VERTICAL {
		return s.Rect().Height-s.btnLeftTop.Rect().Height-s.btnRightDown.Rect().Height-s.scbScroller.Rect().Height;
	} 
	return s.Rect().Width-s.btnLeftTop.Rect().Width-s.btnRightDown.Rect().Width-s.scbScroller.Rect().Width;
}

func (s *Scrollbar) UpdateScrollerPos() {
	delta := s.maxValue - s.minValue
	size := 0
	if delta != 0 {
		size = int(float32(s.value - s.minValue) * float32(s.ScrollAreaSize()) / float32(delta))
	}
		
	if s.orientation == SCROLLBAR_VERTICAL {
		size += s.btnLeftTop.Rect().Y + s.btnLeftTop.Rect().Height
		s.scbScroller.Rect().Y = size
	} else {
		size += s.btnRightDown.Rect().X + s.btnRightDown.Rect().Width
		s.scbScroller.Rect().X = size
	}
}
