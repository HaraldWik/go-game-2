package main

import (
	"fmt"
	"time"

	"github.com/HaraldWik/go-game-2/scr/app"
	dt "github.com/HaraldWik/go-game-2/scr/data-types"
	"github.com/HaraldWik/go-game-2/scr/input"
	sys "github.com/HaraldWik/go-game-2/scr/systems"
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

func main() {
	app := app.New()
	win := app.NewWindow("Window!!!", vec2.New(1920, 1075))
	win.Flags = win.FLAG_RESIZABLE
	win.Open()

	zoom := float32(10)

	cam := ups.NewObject()
	cam.AddData("Window", win)
	cam.AddData("Transform", dt.NewTransform2D(vec2.Zero(), vec2.All(1.0), 0.0))
	cam.AddData("Zoom", zoom)
	cam.AddSystems(sys.Camera2D{})
	ups.Manager.AddObject("Camera", cam)

	cir := ups.NewObject()
	cir.AddData("Color", vec3.New(1.0, 0.5, 0.0))
	cir.AddData("Transform", dt.NewTransform2D(vec2.New(-0.5, -0.9), vec2.All(1.0), 0.0))
	cir.AddData("Segments", uint32(30))
	cir.AddSystems(sys.RenderCircle2D{})
	ups.Manager.AddObject("Circle", cir)

	tri := ups.NewObject()
	tri.AddData("Color", vec3.New(0.0, 1.0, 0.0))
	tri.AddData("Transform", dt.NewTransform2D(vec2.New(2.0, -3.0), vec2.New(1.0, 1.0), 0.0))
	tri.AddData("Flip", true)
	tri.AddSystems(sys.RenderTriangle2D{})
	ups.Manager.AddObject("Triangle", tri)

	rect := ups.NewObject()
	rect.AddData("Color", vec3.New(1.0, 0.0, 0.0))
	rect.AddData("Transform", dt.NewTransform2D(vec2.New(-1.0, 2.0), vec2.All(4.0), 60.0))
	rect.AddSystems(sys.RenderRect2D{})
	ups.Manager.AddObject("Rectangle", rect)

	for !win.CloseEvent() {
		win.BeginDraw(vec3.New(0.0, 0.144, 0.856))

		if input.IsPressed(input.K_W) {
			zoom += 0.1
		}
		if input.IsPressed(input.K_S) {
			zoom -= 0.1
			fmt.Println(zoom)
		}

		cam.AddData("Zoom", zoom)

		ups.Manager.Update()

		win.EndDraw(60)
	}
}

type Sys1 struct{}

func (s Sys1) Update() {
	obj := ups.Manager.GetParent()
	obj.Data.Set("Data", obj.Data.Get("Data").(int32)+int32(time.Now().Second()))
	fmt.Println("Updated")

	fmt.Println(obj.Data.Get("Data"))
}

type Sys2 struct{}

func (s Sys2) Update() {
	obj := ups.Manager.GetParent()
	obj.Data.Set("Data2", obj.Data.Get("Data2").(int32)-int32(time.Now().Second()))
	fmt.Println("Updated")

	fmt.Println(obj.Data.Get("Data2"))
}

/*
// TEST

var obj = ups.NewObject() // <-- this is temporary and I already know how to fix this!

func main() {
	for {
		myvar := int32(78)

		obj.AddData("Hello", myvar)
		obj.AddSystem(MMM{})
		obj.Update()

		fmt.Print(obj.Data.Get("Hello"))

		time.Sleep(time.Second)
	}
}

type MMM struct{}

func (mmm MMM) Update() {
	obj.Data.Set("Hello", obj.Data.Get("Hello").(int32)+1)
	fmt.Println("Updated")
}
*/
