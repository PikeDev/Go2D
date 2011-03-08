package sdl

//#include "SDL.h"
import "C"

type SDLEvent struct {
    Evtype uint32
    rest [48]byte
}

func (e *SDLEvent) Window() *WindowEvent {
    return (*WindowEvent)(cast(e))
}

func (e *SDLEvent) Keyboard() *KeyboardEvent {
    return (*KeyboardEvent)(cast(e))
}

func (e *SDLEvent) TextEdit() *TextEditingEvent {
    return (*TextEditingEvent)(cast(e))
}

func (e *SDLEvent) TextInput() *TextInputEvent {
    return (*TextInputEvent)(cast(e))
}

func (e *SDLEvent) MouseMotion() *MouseMotionEvent {
    return (*MouseMotionEvent)(cast(e))
}

func (e *SDLEvent) MouseButton() *MouseButtonEvent {
    return (*MouseButtonEvent)(cast(e))
}

func (e *SDLEvent) MouseWheel() *MouseWheelEvent {
    return (*MouseWheelEvent)(cast(e))
}

func (e *SDLEvent) Quit() *QuitEvent {
    return (*QuitEvent)(cast(e))
}

func (e *SDLEvent) User() *UserEvent {
    return (*UserEvent)(cast(e))
}

func (e *SDLEvent) SysWM() *SysWMEvent {
    return (*SysWMEvent)(cast(e))
}

type WindowEvent struct {
    Evtype uint32
    WindowID uint32
    Event uint8
    Padding1 uint8
    Padding2 uint8
    Padding3 uint8
    Data1 int32
    Data2 int32
}

type KeyboardEvent struct {
    Evtype uint32
    WindowID uint32
    State uint8
    Repeat uint8
    Padding2 uint8
    Padding3 uint8
    keysym C.SDL_keysym
}

func (e *KeyboardEvent) Keysym() *KeySym {
	return (*KeySym)(cast(&e.keysym))
}

type KeySym struct {
	Scancode int32//C.enum_SDL_scancode
	Sym int32
	Mod uint16
	Unicode uint32
}

type TextEditingEvent struct {
    Evtype uint32
    WindowID uint32
    Text [32]byte
    Start int32
    Length int32
}

type TextInputEvent struct {
    Evtype uint32
    WindowID uint32
    Text [32]byte
}

type MouseMotionEvent struct {
    Evtype uint32
    WindowID uint32
    State uint8
    Padding1 uint8
    Padding2 uint8
    Padding3 uint8
    X int32
    Y int32
    Xrel int32
    Yrel int32
}

type MouseButtonEvent struct {
    Evtype uint32
    WindowID uint32
    Button uint8
    State uint8
    Padding1 uint8
    Padding2 uint8
    X int32
    Y int32
}

type MouseWheelEvent struct {
    Evtype uint32
    WindowID uint32
    X int32
    Y int32
}

type QuitEvent struct {
    Evtype uint32
}

type UserEvent struct {
    Evtype uint32
    WindowID uint32
    Code int32
    Data1 *[0]byte
    Cata2 *[0]byte
}

type SysWMEvent struct {
    Evtype uint32
}
