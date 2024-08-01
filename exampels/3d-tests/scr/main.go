package main

import (
	"github.com/HaraldWik/go-game-2/scr/app"
	dt "github.com/HaraldWik/go-game-2/scr/data-types"
	"github.com/HaraldWik/go-game-2/scr/input"
	load "github.com/HaraldWik/go-game-2/scr/loaders"
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

	sceneMain := ups.NewScene()

	sceneMain.NewObject(
		"Camera3D",
		ups.Data{
			"Window":    win,
			"Transform": dt.NewTransform3D(vec3.New(0.0, 0.0, 0.0), vec3.All(1.0), vec3.New(0.0, 45.0, 0.0)),
			"Fov":       float32(50),
			"Speed":     float32(10.0),
		},
		[]ups.System{
			sys.Camera3D{},
			CameraController3D{},
		},
		"Camera",
	)

	sceneMain.NewObject(
		"Cube3D",
		ups.Data{
			"Color":     vec3.All(1.0),
			"Transform": dt.NewTransform3D(vec3.New(0.0, 0.0, -10.0), vec3.All(0.1), vec3.New(45.0, 45.0, 45.0)),
		},
		[]ups.System{
			sys.RenderCube3D{},
			Jump{},
		},
	)

	sceneMain.NewObject(
		"Model",
		ups.Data{
			"Material":  dt.NewMaterial(load.PNG("../assets/Chair jungle_1.png"), vec3.All(1.0)),
			"Transform": dt.NewTransform3D(vec3.New(0.0, 0.0, -10.0), vec3.All(0.1), vec3.New(45.0, 45.0, 45.0)),
			"Model":     load.OBJ("../assets/Chair jungle.obj"),
		},
		[]ups.System{
			sys.RenderObj3D{},
		},
	)

	for !win.CloseEvent() {
		win.BeginDraw(vec3.New(0.0, 0.144, 0.856))

		sceneMain.Update(win.GetDeltaTime())

		win.EndDraw(60)
	}
}

type Jump struct{}

func (j Jump) Start(obj *ups.Object) {}

func (j Jump) Update(obj *ups.Object, deltaTime float32) {
	var (
		transform = obj.Data.Get("Transform").(dt.Transform3D)
	)

	transform.Pos.X = 2
}

type CameraController3D struct{}

func (c CameraController3D) Start(obj *ups.Object) {
	var (
		transform = obj.Data.Get("Transform").(dt.Transform3D)
	)

	obj.Data.Set("Target", transform)
}

func (c CameraController3D) Update(obj *ups.Object, deltaTime float32) {
	var (
		transform = obj.Data.Get("Transform").(dt.Transform3D)
		speed     = obj.Data.Get("Speed").(float32)
		target    = obj.Data.Get("Target").(dt.Transform3D)
	)

	if input.IsPressed(input.K_W) {
		target.Pos.Z += speed * deltaTime
	}

	if input.IsPressed(input.K_S) {
		target.Pos.Z -= speed * deltaTime
	}

	if input.IsPressed(input.K_A) {
		target.Pos.X += speed * deltaTime
	}
	if input.IsPressed(input.K_D) {
		target.Pos.X -= speed * deltaTime
	}

	if input.IsPressed(input.K_SPACE) {
		target.Pos.Y += speed * deltaTime
	}
	if input.IsPressed(input.K_LEFT_CONTROL) {
		target.Pos.Y -= speed * deltaTime
	}

	transform.Pos.X = Lerp(transform.Pos.X, target.Pos.X, 0.2*deltaTime)
	transform.Pos.Y = Lerp(transform.Pos.Y, target.Pos.Y, 0.2*deltaTime)
	transform.Pos.Z = Lerp(transform.Pos.Z, target.Pos.Z, 0.2*deltaTime)

	obj.Data.Set("Transform", transform)
}

func Lerp(a, b, t float32) float32 {
	return a + t*(b-a)
}
