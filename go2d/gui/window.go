package go2d

type IWindow interface {
	NextFocus()
	FocusElement(element IElement)
}

type Window struct {
	Container

	title string

	popups []*Window
}

func NewWindow(x, y, width, height int) *Window {
	window := &Window{}
	window.Init(x, y, width, height)
	return window
}

func (w *Window) AddChild(element IElement) {
	element.SetParent(w)
	w.children = append(w.children, element)
}

func (w *Window) KeyDown(button int) {
	if button == KEY_TAB {
		w.NextFocus()
	} else {
		w.Container.KeyDown(button)
	}
}

func (w *Window) NextFocus() {
	var elements []IElement = make([]IElement, 0)
	w.FillFocusList(&elements)
	
	if len(elements) > 0 {
		if len(elements) > 1 {
			currentFocusIndex := 0
			for index, element := range elements {
				if element.Focus() {
					currentFocusIndex = index
				}
				element.SetFocus(false)
			}

			currentFocusIndex++
			if currentFocusIndex > (len(elements)-1) {
				currentFocusIndex = 0
			}
			
			elements[currentFocusIndex].SetFocus(true)
		} else {
			elements[0].SetFocus(true)
		}
	}
}

func (w *Window) FocusElement(element IElement) {
	var elements []IElement = make([]IElement, 0)
	w.FillFocusList(&elements)
	
	for _, elem := range elements {
		elem.SetFocus(false)
	}
	
	element.SetFocus(true)
}