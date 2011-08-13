package go2d

import (
	"sdl"
	"exec"
	"path"
	"os"
)

func Sleep(_ticks uint32) {
	sdl.Delay(_ticks)
}

func GetTicks() uint32 {
	return sdl.GetTicks()
}

func KeyDown(_key int) bool {
	return sdl.KeyDown(_key)
}

func GetPath() string {
	file, _ := exec.LookPath(os.Args[0])
	dir, _ := path.Split(file)
	os.Chdir(dir)
	path, _ := os.Getwd()
	return path + "/"
}

func DrawFillRect(_rect *Rect, _red, _green, _blue, _alpha uint8) {
	sdl.SetRenderDrawColor(g_game.renderer, _red, _green, _blue, _alpha)
	sdl.SetRenderDrawBlendMode(g_game.renderer, sdl.BLENDMODE_BLEND)
	sdl.RenderFillRect(g_game.renderer, *_rect.toSDL())
	sdl.SetRenderDrawBlendMode(g_game.renderer, sdl.BLENDMODE_NONE)
	sdl.SetRenderDrawColor(g_game.renderer, 0, 0, 0, 255)
}

func DrawRect(_rect *Rect, _red, _green, _blue, _alpha uint8) {
	sdl.SetRenderDrawColor(g_game.renderer, _red, _green, _blue, _alpha)
	sdl.SetRenderDrawBlendMode(g_game.renderer, sdl.BLENDMODE_BLEND)
	sdl.RenderDrawRect(g_game.renderer, *_rect.toSDL())
	sdl.SetRenderDrawBlendMode(g_game.renderer, sdl.BLENDMODE_NONE)
	sdl.SetRenderDrawColor(g_game.renderer, 0, 0, 0, 255)
}

func DrawLine(_red, _green, _blue, _alpha uint8, _x1, _y1, _x2, _y2 int) {
	sdl.SetRenderDrawColor(g_game.renderer, _red, _green, _blue, _alpha)
	sdl.SetRenderDrawBlendMode(g_game.renderer, sdl.BLENDMODE_BLEND)
	sdl.RenderDrawLine(g_game.renderer, _x1, _y1, _x2, _y2)
	sdl.SetRenderDrawBlendMode(g_game.renderer, sdl.BLENDMODE_NONE)
	sdl.SetRenderDrawColor(g_game.renderer, 0, 0, 0, 255)
}


