package main

import (
	"fmt"

	"github.com/HaraldWik/go-game-2/scr/app"
	"github.com/HaraldWik/go-game-2/scr/scene"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

func main() {
	app := app.NewApp()
	win := app.NewWindow("Window-1", vec2.New(500, 400))
	win.Flags = win.FLAG_RESIZABLE
	win.Open()

	lvl1 := scene.New2D()
	obj1 := lvl1.Create(vec2.New(20, 20), vec2.New(1.0, 1.0), 0.0, vec3.Zero(), nil)
	obj1.Comp = append(obj1.Comp, WHY)
	defer fmt.Println(obj1.Pos)

	for !win.CloseEvent() {
		lvl1.Update()

		win.BeginDraw(0.3, 0.2, 0.5)

		lvl1.Update()

		win.EndDrawOpenGL(60)
	}
}

func WHY(obj *scene.Obj2D) {
	print("HELLO ")
}
