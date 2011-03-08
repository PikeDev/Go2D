package go2d

import (
	"sdl"
)

func EventHandler(_event *sdl.SDLEvent) {
	switch _event.Evtype {
	case sdl.SDL_WINDOWEVENT:
		HandleWindowEvent(_event.Window())

	case sdl.SDL_KEYDOWN, sdl.SDL_TEXTINPUT:
		HandleKeyboardEvent(_event.Keyboard())

	case sdl.SDL_MOUSEBUTTONUP, sdl.SDL_MOUSEBUTTONDOWN:
		HandleMouseButtonEvent(_event.MouseButton())

	case sdl.SDL_MOUSEMOTION:
		HandleMouseMotionEvent(_event.MouseMotion())

	case sdl.SDL_MOUSEWHEEL:
		HandleMouseWheelEvent(_event.MouseWheel())
	}
}

func HandleWindowEvent(_event *sdl.WindowEvent) {
	switch _event.Event {
	case sdl.SDL_WINDOWEVENT_CLOSE:
		g_running = false
	}
}

func HandleKeyboardEvent(_event *sdl.KeyboardEvent) {
	switch _event.Evtype {
	case sdl.SDL_KEYDOWN:
		if g_game.keydownFun != nil {
			g_game.keydownFun(int(_event.Keysym().Scancode))
		}

	case sdl.SDL_TEXTINPUT:
		keysym := uint8(_event.State)
		if keysym != 0 && keysym > 31 {
			if g_game.textinputFun != nil {
				g_game.textinputFun(keysym)
			}
		}
	}
}

func HandleMouseButtonEvent(_event *sdl.MouseButtonEvent) {
	switch _event.Evtype {
	case sdl.SDL_MOUSEBUTTONUP:
		if g_game.mouseupFun != nil {
			g_game.mouseupFun(int16(_event.X), int16(_event.Y))
		}
	case sdl.SDL_MOUSEBUTTONDOWN:
		if g_game.mousedownFun != nil {
			g_game.mousedownFun(int16(_event.X), int16(_event.Y))
		}
	}
}

func HandleMouseMotionEvent(_event *sdl.MouseMotionEvent) {
	if g_game.mousemoveFun != nil {
		g_game.mousemoveFun(int16(_event.X), int16(_event.Y))
	}
}

func HandleMouseWheelEvent(_event *sdl.MouseWheelEvent) {
	if 0-_event.Y < 0 {
		//scroll up
	} else if (0 - _event.Y) > 0 {
		//scroll down
	}
}
