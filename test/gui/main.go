 package main

import (
	"go2d"
)

//Constants
const (
	WINDOW_WIDTH = 800
	WINDOW_HEIGHT = 600
)

//Globals
var (
	g_game *go2d.Game
	g_window *go2d.Window
)

func start() {
	//Load up our default font for GUI text elements
	arial := go2d.NewFont("arial.ttf", 14)
	arial.SetStyle(true, false, false)
	arial.SetColor(255, 255, 255)
	
	//Initialize the GUI system (use whole window area)
	g_window = g_game.InitGUI(0, 0, WINDOW_WIDTH, WINDOW_HEIGHT, arial)

	//Set up some elements
	panel := go2d.NewPanel(20, 20, 200, 200)
	panel.SetBackgroundColor(80, 80, 255)
	
	label := go2d.NewLabel(10, 10, "This is a test label")
	panel.AddChild(label)
	
	g_window.AddChild(panel)
}

func update() {
	//game mechanics here
}

func draw() {

} 

func main() {
	g_game = go2d.NewGame("Test Game")
	g_game.SetDimensions(WINDOW_WIDTH, WINDOW_HEIGHT)
	g_game.SetD3D(true)
	
	g_game.SetInitFun(start)
	g_game.SetUpdateFun(update)
	g_game.SetDrawFun(draw)
	
	g_game.Run()
}