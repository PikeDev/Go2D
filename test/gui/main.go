 package main

import (
	"go2d"
	"fmt"
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
	
	//A panel
	panel := go2d.NewPanel(20, 20, 200, 200)
	panel.SetBackgroundColor(80, 80, 255)
	
	//Another panel in the panel that gets clipped
	innerPanel := go2d.NewPanel(180, 180, 50,50)
	innerPanel.SetBackgroundColor(255, 80, 80)
	panel.AddChild(innerPanel)
	
	//A label
	label := go2d.NewLabel(10, 10, "This is a test label")
	panel.AddChild(label)
	
	//A button
	button := go2d.NewButton(10, 40, 100, 30, "Click here")
	button.SetFontColor(80, 80, 80)
	button.SetBackgroundColor(80, 255, 80)
	button.SetOnClickListener(func(x, y int) {
		println("Button clicked!")
	})
	panel.AddChild(button)
	
	//A textfield
	textfield := go2d.NewTextField(10, 80, 150, 30)
	textfield.SetFontColor(0, 0, 0)
	textfield.SetBackgroundColor(255, 255, 255)
	textfield.SetBorderColor(0, 0, 0)
	panel.AddChild(textfield)
	
	//A password textfield
	password := go2d.NewTextField(10, 120, 150, 30)
	password.SetPassword(true)
	password.SetFontColor(0, 0, 0)
	password.SetBackgroundColor(255, 255, 255)
	password.SetBorderColor(0, 0, 0)
	panel.AddChild(password)
	
	g_window.AddChild(panel)
	
	//A customized button
	customButton := go2d.NewButton(20, 240, 186, 52, "")
	customButton.SetImage(go2d.NewImage("button_normal.png"))
	customButton.SetHoverImage(go2d.NewImage("button_hover.png"))
	customButton.SetMouseDownImage(go2d.NewImage("button_down.png"))
	customButton.SetOnClickListener(CustomButtonOnClick)
	g_window.AddChild(customButton)
	
	//A scrollbar
	scrollBar := go2d.NewScrollbar(300, 20, 20, 140, go2d.SCROLLBAR_VERTICAL)
	scrollBar.SetOnValueChangeListener(func(value int) {
		fmt.Printf("Scrollbar value: %d\n", value)
	})
	g_window.AddChild(scrollBar)
}

func CustomButtonOnClick(x, y int)  {
	println("Custom button clicked!")
}

func update(dt uint32) {
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