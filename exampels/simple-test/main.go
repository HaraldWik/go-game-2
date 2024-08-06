package main

import (
	"fmt"

	d2d "github.com/HaraldWik/go-game-2/scr/2d/data"
	s2d "github.com/HaraldWik/go-game-2/scr/2d/systems"
	"github.com/HaraldWik/go-game-2/scr/app"
	"github.com/HaraldWik/go-game-2/scr/input"
	load "github.com/HaraldWik/go-game-2/scr/loaders"
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
	vec4 "github.com/HaraldWik/go-game-2/scr/vector/4"
)

func main() {
	app := app.New()
	window := app.NewWindow(
		"MyWindowName",
		vec2.New(500, 400),
	)

	window.SetMaxFPS(65)
	window.SetFlags(window.FLAG_RESIZABLE)
	window.Open()

	MainScene := ups.SceneManager.New()

	MainScene.New(
		"Skybox",
		ups.Data{
			"Color": vec3.New(0.0, 0.25, 0.65),
		},
		[]ups.System{
			s2d.Skybox2D{},
		},
	)

	MainScene.New(
		"Camera2D",
		ups.Data{
			"Window": window,
			"Transform": d2d.NewTransform2D(
				vec2.New(0.0, 0.0),
				vec2.All(1.0), 0.0,
			),
			"Zoom": float32(10),
		},
		[]ups.System{
			s2d.Camera2D{},
		},
	)

	MainScene.New(
		"MyObject",
		ups.Data{
			"MyData": int32(30.0),
		},
		[]ups.System{},
		"MyTag",
	)

	MainScene.New(
		"RenderingObject",
		ups.Data{
			"Material": d2d.NewMaterial2D(
				load.NewEmptyTexture(), // <-- Texture, it can load real textures!
				vec4.Zero(),
				vec3.All(1.0), // <-- Alpha
				1.0,           // <-- Transparency
				1.0,           // <-- Z level
			),
			"Transform": d2d.NewTransform2D(
				vec2.Zero(),        // <-- Position
				vec2.New(4.0, 2.0), // <-- Scale
				45.0,               // <-- Rotation in degress
			),

			"Speed": float32(2.5),

			"Window": window,
		},
		[]ups.System{
			s2d.RenderRectangle2D{},
			Rotate{},
			MoveToMouse{},
		},
		"MyTag",
	)

	for !window.CloseEvent() {
		window.BeginDraw()
		ups.SceneManager.Update(window.GetDeltaTime())
		window.EndDraw()
	}
}

func WindowToNDC(window app.Window) vec2.Type {
	mousePos := input.MousePosition()
	windowSize := window.GetSize()

	// Debug output
	fmt.Printf("Mouse Position (Window Coords): X = %.2f, Y = %.2f\n", mousePos.X, mousePos.Y)
	fmt.Printf("Window Size: Width = %.2f, Height = %.2f\n", windowSize.X, windowSize.Y)

	// Convert mouse position to NDC
	ndcX := (2.0*mousePos.X)/windowSize.X - 1.0
	ndcY := 1.0 - (2.0*mousePos.Y)/windowSize.Y

	// Debug output
	fmt.Printf("NDC Coordinates: X = %.2f, Y = %.2f\n", ndcX, ndcY)

	return vec2.New(ndcX, ndcY)
}

type MoveToMouse struct{}

func (m MoveToMouse) FixedUpdate(obj *ups.Object) {
	var (
		window    = obj.Data.Get("Window").(app.Window)
		transform = obj.Data.Get("Transform").(d2d.Transform2D)
	)

	transform.Position = WindowToNDC(window)

	obj.Data.Set("Transform", transform)
}

type Rotate struct{}

func (r Rotate) FixedUpdate(obj *ups.Object) {
	var (
		transform = obj.Data.Get("Transform").(d2d.Transform2D)
		speed     = obj.Data.Get("Speed").(float32)
	)

	transform.Rotation += speed

	obj.Data.Set("Transform", transform)
}

type MySystem struct{}

func (s MySystem) Start(obj *ups.Object) {
	fmt.Printf("Object '%s' in scene '%d' has Started\n", obj.Name, obj.Scene.ID)

	if obj.Tags.Has("Mytag") {
		fmt.Printf("Object '%s' has tag MyTag!\n", obj.Name)
	}
}

func (s MySystem) Update(obj *ups.Object, deltaTime float32) {
	var (
		myData = obj.Data.Get("MyData").(int32)
	)

	if input.IsJustPressed(input.M_LMB) {
		myData += 1
	}

	if input.IsJustPressed(input.M_RMB) {
		myData += 1
	}

	obj.Data.Set("MyData", myData)
}
