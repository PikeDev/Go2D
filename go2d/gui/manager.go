package gui

type Manager struct {
	root *Window
}

func NewManager(x, y, width, height int) *Manager {
	manager := &Manager{}
	manager.root = NewWindow(x, y, width, height)
	return manager
}

func (m *Manager) Draw() {
	m.root.Draw(m.root.Rect())
}

func (m *Manager) MouseDown(x, y int) {
	if m.root != nil {
		m.MouseDown(x, y)
	}
}

func (m *Manager) MouseUp(x, y int) {
	if m.root != nil {
		m.MouseUp(x, y)
	}
}

func (m *Manager) MouseMove(x, y int) {
	if m.root != nil {
		m.MouseMove(x, y)
	}
}
func (m *Manager) MouseScroll(direction int) {
	if m.root != nil {
		m.MouseScroll(direction)
	}
}

func (m *Manager) KeyDown(button int) {
	if m.root != nil {
		m.KeyDown(button)
	}
}

func (m *Manager) KeyUp(button int) {
	if m.root != nil {
		m.KeyUp(button)
	}
}

func (m *Manager) TextInput(character uint8) {
	if m.root != nil {
		m.TextInput(character)
	}
}