package main

import "go2d"

func start() {
	//load resources here
}

func update(dt uint32) {
	//game mechanics here
	//this is called every frame before the draw function
}

func draw() {
	//rendering here
	//this is called every frame
} 

func mouseMove(x, y int16) {
	//mouse move events
}

func mouseUp(x, y int16) {
	//mouse up events
}

func mouseDown(x, y int16) {
	//mouse down events
}

func textInput(char uint8) {
	//text input events
}

func keyDown(key int) {
	//key down events
}

func main() {
	game := go2d.NewGame("Skeleton")
	game.SetDimensions(200, 200)
	
	//Set to false when OpenGL should also be defaulted on Windows
	game.SetD3D(true)
	
	game.SetInitFun(start)
	game.SetUpdateFun(update)
	game.SetDrawFun(draw)
	
	game.SetMouseMoveFun(mouseMove)
	game.SetMouseDownFun(mouseDown)
	game.SetMouseUpFun(mouseUp)
	game.SetTextInputFun(textInput)
	game.SetKeyDownFun(keyDown)
	
	game.Run()
}