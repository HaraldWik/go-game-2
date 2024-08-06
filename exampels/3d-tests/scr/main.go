package main

import (
	s2d "github.com/HaraldWik/go-game-2/scr/2d/systems"
	d3d "github.com/HaraldWik/go-game-2/scr/3d/data"
	s3d "github.com/HaraldWik/go-game-2/scr/3d/systems"
	"github.com/HaraldWik/go-game-2/scr/app"
	"github.com/HaraldWik/go-game-2/scr/input"
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

var SceneMain = ups.SceneManager.New()

func main() {
	application := app.New()
	window := application.NewWindow("Window!!!", vec2.New(1920, 1075))
	window.SetFlags(window.FLAG_RESIZABLE)
	window.SetMaxFPS(60)
	window.Open()

	SceneMain.New(
		"Camera3D",
		ups.Data{
			"Window": window,
			"Transform": d3d.NewTransform3D(
				vec3.New(0.0, 0.0, 0.0),
				vec3.All(1.0),
				vec3.New(0.0, 45.0, 0.0),
			),
			"Fov":   float32(50),
			"Speed": float32(10.0),
		},
		[]ups.System{
			s3d.Camera3D{},
			CameraController3D{},
		},
	)

	SceneMain.New(
		"Skybox",
		ups.Data{
			"Color": vec3.New(0.0, 0.1, 0.9),
		},
		[]ups.System{
			s2d.Skybox2D{},
		},
	)

	SceneMain.New(
		"Cube3D",
		ups.Data{
			"Color": vec3.All(1.0),
			"Transform": d3d.NewTransform3D(
				vec3.New(0.0, 0.0, -10.0),
				vec3.All(100.0),
				vec3.New(45.0, 45.0, 45.0),
			),
		},
		[]ups.System{
			s3d.RenderCube3D{},
		},
	)

	for !window.CloseEvent() {
		window.BeginDraw()

		ups.SceneManager.Update(window.GetDeltaTime())

		window.EndDraw()
	}
}

type CameraController3D struct{}

func (c CameraController3D) Start(obj *ups.Object) {
	var (
		transform = obj.Data.Get("Transform").(d3d.Transform3D)
	)

	obj.Data.Set("Target", transform)
}

func (c CameraController3D) Update(obj *ups.Object, deltaTime float32) {
	var (
		transform = obj.Data.Get("Transform").(d3d.Transform3D)
		speed     = obj.Data.Get("Speed").(float32)
		target    = obj.Data.Get("Target").(d3d.Transform3D)
	)

	if input.IsPressed(input.K_W) {
		target.Position.Z += speed * deltaTime
	}

	if input.IsPressed(input.K_S) {
		target.Position.Z -= speed * deltaTime
	}

	if input.IsPressed(input.K_A) {
		target.Position.X += speed * deltaTime
	}
	if input.IsPressed(input.K_D) {
		target.Position.X -= speed * deltaTime
	}

	if input.IsPressed(input.K_SPACE) {
		target.Position.Y += speed * deltaTime
	}
	if input.IsPressed(input.K_LEFT_CONTROL) {
		target.Position.Y -= speed * deltaTime
	}

	transform.Position.X = Lerp(transform.Position.X, target.Position.X, 0.2*deltaTime)
	transform.Position.Y = Lerp(transform.Position.Y, target.Position.Y, 0.2*deltaTime)
	transform.Position.Z = Lerp(transform.Position.Z, target.Position.Z, 0.2*deltaTime)

	obj.Data.Set("Transform", transform)
}

func Lerp(a, b, t float32) float32 {
	return a + t*(b-a)
}
