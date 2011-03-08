package sdl

//#include "SDL.h"
import "C" 
import "unsafe"

type cast unsafe.Pointer

func Delay(_ticks uint32) {
	C.SDL_Delay(C.Uint32(_ticks))
}

func GetTicks() uint32 { 
    return uint32(C.SDL_GetTicks())
}   

func Quit() {
	C.SDL_Quit()
}

func GetError() (ret string) {
	ret = C.GoString(C.SDL_GetError())
	C.SDL_ClearError()	
	return
}

func KeyDown(_key int) bool {
	zero := C.int(0)
	var state = uintptr(unsafe.Pointer(C.SDL_GetKeyboardState(&zero)))+uintptr(_key)
	down := (*uint8)(cast(state))
	if *down == 1 {
		return true
	}
	return false
}

func Init() (error string) {
	flags := int64(C.SDL_INIT_VIDEO)
    if C.SDL_Init(C.Uint32(flags)) != 0 {
        error = C.GoString(C.SDL_GetError())
        return
    }
    return ""
}

type Window struct { 
	window *C.SDL_Window
}

func CreateWindow(_title string, _width int, _height int) (*Window, string) {
	ctitle := C.CString(_title)
    var window *C.SDL_Window = C.SDL_CreateWindow(ctitle, 
					 							  C.SDL_WINDOWPOS_CENTERED, 
												  C.SDL_WINDOWPOS_CENTERED,
												  C.int(_width), 
												  C.int(_height), 
												  C.SDL_WINDOW_SHOWN | C.SDL_WINDOW_OPENGL)
	C.free(unsafe.Pointer(ctitle))
    if window == nil {
        return nil, GetError()
    }
    return &Window{window}, ""
}

func DestroyWindow(_window *Window) {
	C.SDL_DestroyWindow(_window.window)
}

func PollEvent() (*SDLEvent, bool) {
	var ev *SDLEvent = &SDLEvent{}
    if C.SDL_PollEvent((*C.SDL_Event)(cast(ev))) != 0 {
		return ev, true
    }
	return nil, false
}


