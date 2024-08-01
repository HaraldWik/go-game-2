package main

import (
	"github.com/HaraldWik/go-game-2/scr/app"
	dt "github.com/HaraldWik/go-game-2/scr/data-types"
	load "github.com/HaraldWik/go-game-2/scr/loaders"
	sys "github.com/HaraldWik/go-game-2/scr/systems"
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

func main() {
	app := app.New()
	window := app.NewWindow("Gnome Jumper 24", vec2.New(800, 600))
	window.Flags = window.FLAG_RESIZABLE
	window.MaxFPS = 60
	window.Open()

	sceneMain := ups.SceneManager.NewScene()

	sceneMain.NewObject(
		"Skybox",
		ups.Data{
			"Color": vec3.New(0.0, 0.25, 0.75),
		},
		[]ups.System{
			sys.Skybox2D{},
		},
	)

	sceneMain.NewObject(
		"Camera2D",
		ups.Data{
			"Window": window,
			"Transform": dt.NewTransform2D(
				vec2.Zero(), vec2.Zero(), 0.0,
			),
			"Zoom": float32(10),
		},
		[]ups.System{
			sys.Camera2D{},
		},
	)

	sceneMain.NewObject(
		"Obj",
		ups.Data{
			"Material": dt.NewMaterial(
				load.EmptyTexture(),
				vec3.New(1.0, 0.0, 0.0),
			),

			"Transform": dt.NewTransform2D(
				vec2.New(0.0, 5.0), vec2.New(1.0, 3.0), 0.0,
			),
		},
		[]ups.System{
			sys.RenderRectangle2D{},
		},
	).Clone(
		ups.Data{
			"Transform": dt.NewTransform2D(
				vec2.New(-5.0, 5.0), vec2.New(1.0, 3.0), 0.0,
			),
		},
		ups.Data{
			"Transform": dt.NewTransform2D(
				vec2.New(5.0, 5.0), vec2.New(1.0, 3.0), 0.0,
			),
		},
	)

	ups.SceneManager.SetCurrentScenes(sceneMain.ID)

	for !window.CloseEvent() {
		window.BeginDraw()

		ups.SceneManager.Update(window.GetDeltaTime())

		window.EndDraw()
	}
}
