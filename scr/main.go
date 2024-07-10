package main

import (
	"github.com/HaraldWik/go-game-2/scr/app"
	mod "github.com/HaraldWik/go-game-2/scr/modules"

	"github.com/HaraldWik/go-game-2/scr/abus"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

func main() {
	app := app.NewApp()
	win := app.NewWindow("Window!!!", vec2.New(1920, 1075))
	win.Flags = win.FLAG_RESIZABLE
	win.Open()

	scene := abus.SceneManager.NewScene()

	scene.Create(&mod.Cam2D{Win: win, Zoom: 10.0})

	for i := 0; i < 50; i++ {
		stdRect := abus.NewObject(&mod.RenderRect2D{Transform: mod.Transform2D{Pos: vec2.New(float32(i), float32(i))}, Color: vec3.All(1.0)})
		scene.Instance(stdRect)
	}

	for !win.CloseEvent() {
		win.BeginDraw(vec3.All(0.175))

		scene.Update()

		win.EndDraw(60)
	}
	/*
	   cam := scene.Instance(abus.NewObject(&mod.Transform3D{Pos: vec3.New(0.0, 0.0, 50.0)}, &mod.Cam3D{Win: win, Fov: 45.0}))
	   transform := mod.GetTransform3D(cam)

	   scene.Instance(abus.NewObject(

	   	&mod.Transform3D{Pos: vec3.New(90.9, 3.0, -30.0), Size: vec3.All(2.0), Rot: vec3.Zero()},
	   	&mod.RenderCube3D{Color: vec3.All(1.0), Angle: 1.0}))

	   obj := mod.GetTransform3D(scene.Objs[1])
	   fmt.Println(obj.Pos)

	   	for !win.CloseEvent() {
	   		win.BeginDraw(vec3.All(0.175))

	   		scene.Update()

	   		if input.IsPressed("SPACE") {
	   			transform.Pos.Y++
	   		}

	   		if input.IsPressed("SHIFT") {
	   			transform.Pos.Y--
	   		}

	   		if input.IsPressed("W") {
	   			transform.Pos.Z++
	   		}

	   		if input.IsPressed("S") {
	   			transform.Pos.Z--
	   		}

	   		if input.IsJustPressed("I") {
	   			fmt.Println(transform)
	   		}

	   		win.EndDraw(60)
	   	}
	*/
}
