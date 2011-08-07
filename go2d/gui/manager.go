package go2d

type GUIManager struct {
	root *Window
}

func NewGUIManager(x, y, width, height int) *GUIManager {
	manager := &GUIManager{}
	manager.root = NewWindow(x, y, width, height)
	return manager
}

func (m *GUIManager) Draw() {
	m.root.Draw(m.root.Rect())
}

func (m *GUIManager) MouseDown(x, y int) {
	if m.root != nil {
		m.MouseDown(x, y)
	}
}

func (m *GUIManager) MouseUp(x, y int) {
	if m.root != nil {
		m.MouseUp(x, y)
	}
}

func (m *GUIManager) MouseMove(x, y int) {
	if m.root != nil {
		m.MouseMove(x, y)
	}
}
func (m *GUIManager) MouseScroll(direction int) {
	if m.root != nil {
		m.MouseScroll(direction)
	}
}

func (m *GUIManager) KeyDown(button int) {
	if m.root != nil {
		m.KeyDown(button)
	}
}

func (m *GUIManager) KeyUp(button int) {
	if m.root != nil {
		m.KeyUp(button)
	}
}

func (m *GUIManager) TextInput(character uint8) {
	if m.root != nil {
		m.TextInput(character)
	}
}