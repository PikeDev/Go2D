package gui

import . "go2d"

type Manager struct {
	root *Window
}

func NewManager(x, y, width, height int) {
	
}

func (m *Manager) Draw(drawArea *Rect) {
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