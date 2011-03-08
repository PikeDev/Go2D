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
