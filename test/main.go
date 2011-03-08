package main

import (
	"go2d"
	"fmt"
)

var image *go2d.Image
var arial16 *go2d.Font

func start() {
	image = go2d.NewImage("test.png")
	
	arial16 = go2d.NewFont("arial.ttf", 16)
	arial16.SetStyle(true, false, false)
	arial16.SetColor(255, 255, 255)
}

func update() {
	//game mechanics here
}

func draw() {
	image.DrawRect(go2d.NewRect(10, 10, 100, 100))
	
	image.Draw(10, 200)
	
	arial16.DrawText("Testing...", 200, 10)
	
	go2d.DrawFillRect(go2d.NewRect(350, 200, 100, 100), 255, 100, 0, 255)
} 

func mouseUp(x, y int16) {
	fmt.Printf("Mouse up x: %d y: %d\n", x, y)
}

func textInput(char uint8) {
	fmt.Printf("Text input: %c\n", char)
}

func keyDown(key int) {
	switch key {
	case go2d.KEY_UP:
		fmt.Println("Going up")
	case go2d.KEY_DOWN:
		fmt.Println("Going down")
	}
}

func main() {
	game := go2d.NewGame("Test Game")
	game.SetDimensions(800, 600)
	game.SetD3D(true)
	
	game.SetInitFun(start)
	game.SetUpdateFun(update)
	game.SetDrawFun(draw)
	
	game.SetMouseUpFun(mouseUp)
	game.SetTextInputFun(textInput)
	game.SetKeyDownFun(keyDown)
	
	game.Run()
}