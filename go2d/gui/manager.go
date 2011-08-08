package go2d

type GUIManager struct {
	root *Window
	defaultFont *Font
}

func NewGUIManager(x, y, width, height int, defaultFont *Font) *GUIManager {
	manager := &GUIManager{}
	manager.root = NewWindow(x, y, width, height)
	manager.defaultFont = defaultFont
	return manager
}

func (m *GUIManager) Root() *Window {
	return m.root
}

func (m *GUIManager) DefaultFont() *Font {
	return m.defaultFont
}

func (m *GUIManager) Draw() {
	if m.root != nil {
		m.root.Draw(m.root.Rect())
	}
}

func (m *GUIManager) MouseDown(x, y int) {
	if m.root != nil {
		m.root.MouseDown(x, y)
	}
}

func (m *GUIManager) MouseUp(x, y int) {
	if m.root != nil {
		m.root.MouseUp(x, y)
	}
}

func (m *GUIManager) MouseMove(x, y int) {
	if m.root != nil {
		m.root.MouseMove(x, y)
	}
}
func (m *GUIManager) MouseScroll(direction int) {
	if m.root != nil {
		m.root.MouseScroll(direction)
	}
}

func (m *GUIManager) KeyDown(button int) {
	if m.root != nil {
		m.root.KeyDown(button)
	}
}

func (m *GUIManager) KeyUp(button int) {
	if m.root != nil {
		m.root.KeyUp(button)
	}
}

func (m *GUIManager) TextInput(character uint8) {
	if m.root != nil {
		m.root.TextInput(character)
	}
}