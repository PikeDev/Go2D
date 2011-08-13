package go2d

import "sdl"

const (
	CARET_TIME = 500
)

type TextField struct {
	Element
	
	//Properties
	BackgroundColor
	BorderColor
	ElementImage
	TextElement
	
	//Listeners
	OnKeyDownListener
	
	//Members
	text string
	readonly bool
	password bool
	caret bool
	caretLast uint32
}

func NewTextField(x, y, width, height int) *TextField {
	textfield := &TextField{}
	textfield.Init(x, y, width, height)
	textfield.focusable = true
	textfield.SetFont(g_game.guiManager.defaultFont)
	textfield.fontColor = &sdl.Color{uint8(255), uint8(255), uint8(255), uint8(255)}
	textfield.caretLast = sdl.GetTicks()
	return textfield
}

func (t *TextField) Text() string {
	return t.text
}

func (t *TextField) SetText(text string) {
	t.text = text
}

func (t *TextField) ReadOnly() bool {
	return t.readonly
}

func (t *TextField) SetReadOnly(readonly bool) {
	t.readonly = readonly
}

func (t *TextField) Password() bool {
	return t.password
}

func (t *TextField) SetPassword(password bool) {
	t.password = password
}

func (t *TextField) KeyDown(button int) {
	if !t.readonly && t.focus {
		if button == KEY_BACKSPACE {
			if len(t.text) > 0 {
				textlength := len(t.text)
				t.text = t.text[0 : textlength-1]
			}
		}
	}
	
	if t.onKeyDown != nil {
		t.onKeyDown(button)
	}
}

func (t *TextField) TextInput(char uint8) {
	if !t.readonly && t.focus {
		if char != 0 && char > 31 { 
			t.text += string(char)
		}
	}
}

func (t *TextField) Draw(drawArea *Rect) {
	realRect := NewRect(t.Rect().X+drawArea.X, t.Rect().Y+drawArea.Y, t.Rect().Width, t.Rect().Height)
	inRect := drawArea.Intersection(realRect)
	
	if t.backgroundColor != nil {
		DrawFillRect(inRect, t.backgroundColor.R, t.backgroundColor.G, t.backgroundColor.B, 255)
	}
	
	if t.borderColor != nil {
		DrawRect(inRect, t.borderColor.R, t.borderColor.G, t.borderColor.B, 255)
	}
	
	if t.image != nil {
		t.image.DrawRectInRect(t.Rect(), drawArea)
	}
	
	caretX := 0
	textX := t.Rect().X+(t.font.GetStringWidth(" "));
	textY := t.Rect().Y+((t.Rect().Height/2)-(t.font.GetStringHeight()/2));
	
	if t.text != "" &&  t.font != nil {
		var drawText string
		if t.password {
			for i := 0; i < len(t.text); i++ {
				drawText += "*"
			}
		} else {
			drawText = t.text
		}

		t.font.SetStyle(t.bold, t.italic, t.underlined)
		t.font.SetColor(t.fontColor.R, t.fontColor.G, t.fontColor.B)
		t.font.DrawTextInRect(drawText, drawArea.X + textX, drawArea.Y + textY, inRect)
		
		caretX = t.font.GetStringWidth(drawText)
	}
	
	//draw caret
	if t.caret && !t.readonly && t.focus {
		caretX += textX
		if inRect.Contains(drawArea.X+caretX, drawArea.Y+textY) {
			DrawLine(t.fontColor.R, t.fontColor.G, t.fontColor.B, 255, drawArea.X+caretX, drawArea.Y+textY, drawArea.X+caretX, drawArea.Y+textY+t.font.GetStringHeight())
		}
	}
	if sdl.GetTicks()-t.caretLast >= CARET_TIME {
		t.caret = !t.caret
		t.caretLast = sdl.GetTicks()
	}
}

func (t *TextField) MouseUp(x, y int) {
	if t.Rect().Contains(x, y) {
		if window := t.Window(); window != nil {
			window.FocusElement(t)
		}
	}
}

