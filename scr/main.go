package main

import (
	"fmt"

	"github.com/HaraldWik/go-game-2/scr/app"
	commponent "github.com/HaraldWik/go-game-2/scr/commponents"
	"github.com/HaraldWik/go-game-2/scr/ecs"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
)

func main() {
	app := app.NewApp()
	win := app.NewWindow("Window-1", vec2.New(500, 400))
	win.Flags = win.FLAG_RESIZABLE
	win.Open()

	scene := ecs.NewScene()
	o := scene.Create(&commponent.Transform2D{})
	defer fmt.Println(o.Components)

	for !win.CloseEvent() {
		win.BeginDraw(0.3, 0.2, 0.5)

		scene.Update()

		win.EndDrawOpenGL(60)
	}
}
