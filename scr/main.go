package main

import (
	"fmt"

	"github.com/HaraldWik/go-game-2/scr/app"
	component "github.com/HaraldWik/go-game-2/scr/components"

	"github.com/HaraldWik/go-game-2/scr/ecs"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

func main() {
	app := app.NewApp()
	win := app.NewWindow("Window-1", vec2.New(500, 400))
	win.Flags = win.FLAG_RESIZABLE
	win.Open()

	scene := ecs.NewScene()

	cam := scene.Create(&component.Camera2D{Window: win, Transform: component.Transform2D{Pos: vec2.Zero(), Size: vec2.All(1.0), Rot: 0.0}, Zoom: 5.0})
	fmt.Println(cam)

	scene.Create(&component.RenderRect{component.Transform2D{Pos: vec2.Zero(), Size: vec2.All(1.0), Rot: 90.0}, vec3.New(1.0, 1.0, 0.0)})
	scene.Create(&component.RenderRect{component.Transform2D{Pos: vec2.New(1.0, -1.5), Size: vec2.All(1.0), Rot: -90.0}, vec3.New(1.0, 0.0, 0.0)})
	scene.Create(&component.RenderRect{component.Transform2D{Pos: vec2.New(-1.4, 0.34), Size: vec2.All(1.0), Rot: -90.0}, vec3.New(0.0, 1.0, 0.0)})

	for !win.CloseEvent() {
		win.Update()
		win.BeginDraw(0.1, 0.2, 0.7)

		scene.Update()

		win.EndDraw(60)
	}
}
