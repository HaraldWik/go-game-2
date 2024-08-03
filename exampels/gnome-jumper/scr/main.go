package main

import (
	"github.com/HaraldWik/go-game-2/exampels/gnome-jumper/scr/data"
	d2d "github.com/HaraldWik/go-game-2/scr/2d/data"
	s2d "github.com/HaraldWik/go-game-2/scr/2d/systems"
	"github.com/HaraldWik/go-game-2/scr/app"
	load "github.com/HaraldWik/go-game-2/scr/loaders"
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

var score uint32

var (
	application = app.New()
	window      = application.NewWindow("Gnome Jumper 24", vec2.New(1920, 1075))
)

var (
	WorldScene    = ups.SceneManager.New()
	MainMenuScene = ups.SceneManager.New()
)

func main() {
	window.SetFlags(window.FLAG_RESIZABLE)
	window.SetMaxFPS(60)
	window.Open()

	// Main menu scene
	MainMenuScene.New(
		"Skybox",
		ups.Data{
			"Color": vec3.Zero(),
		},
		[]ups.System{
			s2d.Skybox2D{},
		},
	)

	MainMenuScene.New(
		"Camera2D",
		ups.Data{
			"Window":    window,
			"Transform": d2d.NewTransform2D(vec2.New(0.0, 0.0), vec2.All(1.0), 0.0),
			"Zoom":      float32(10),
		},
		[]ups.System{
			s2d.Camera2D{},
		},
	)

	MainMenuScene.New(
		"Secret",
		ups.Data{
			"Material": d2d.NewMaterial2D(
				load.NewTexture(data.Get(data.AssetPath)+"texture.png"),
				vec3.All(0.5), 1.0,
			),
			"Transform": d2d.NewTransform2D(
				vec2.New(0.0, 0.0),
				vec2.New(36.75, 20.0),
				0.0,
			),
		},
		[]ups.System{
			s2d.RenderRectangle2D{},
		},
	)

	MainMenuScene.New(
		"Backround",
		ups.Data{
			"Material": d2d.NewMaterial2D(
				load.NewTexture(data.Get(data.AssetPath)+"texture.png"),
				vec3.All(0.5), 1.0,
			),
			"Transform": d2d.NewTransform2D(
				vec2.New(0.0, 0.0),
				vec2.New(36.75, 20.0),
				0.0,
			),
		},
		[]ups.System{
			s2d.RenderRectangle2D{},
		},
	)

	MainMenuScene.New(
		"Button",
		ups.Data{
			"Material": d2d.NewMaterial2D(
				load.NewTexture(data.Get(data.AssetPath)+"texture.png"),
				vec3.All(0.5), 1.0,
			),
			"Transform": d2d.NewTransform2D(
				vec2.New(0.0, 0.0),
				vec2.New(36.75, 20.0),
				0.0,
			),
		},
		[]ups.System{
			s2d.RenderRectangle2D{},
		},
	)

	// World scene

	WorldScene.New(
		"BackroundGnome",
		ups.Data{
			"Material": d2d.NewMaterial2D(
				load.NewTexture(data.Get(data.AssetPath)+"gnome.png"),
				vec3.All(0.5),
				-9.0,
			),
			"Transform": d2d.NewTransform2D(
				vec2.New(0.0, 0.0),
				vec2.New(16.0, 13.0),
				25.0,
			),

			"Min":   float32(-40.0),
			"Max":   float32(40.0),
			"Speed": float32(150.0),
		},
		[]ups.System{
			s2d.RenderRectangle2D{},
			FlipRotate{},
		},
	)

	WorldScene.New(
		"Skybox",
		ups.Data{
			"Color": vec3.New(0.0, 0.25, 0.65),
		},
		[]ups.System{
			s2d.Skybox2D{},
		},
	)

	WorldScene.New(
		"Camera2D",
		ups.Data{
			"Window": window,
			"Transform": d2d.NewTransform2D(
				vec2.New(0.0, 7.5),
				vec2.All(1.0), 0.0,
			),
			"Zoom": float32(10),
		},
		[]ups.System{
			s2d.Camera2D{},
		},
	)

	WorldScene.New(
		"Player",
		ups.Data{
			"Material": d2d.NewMaterial2D(
				load.EmptyTexture(),
				vec3.New(1.0, 0.0, 0.0),
				2.0,
			),
			"Transform": d2d.NewTransform2D(
				vec2.New(-6.0, 0.0),
				vec2.All(1.5),
				0.0,
			),
		},
		[]ups.System{
			s2d.RenderRectangle2D{},

			s2d.AABB{},

			Controller2D{},

			Death{},
		},
	)

	WorldScene.New(
		"Obsticle",
		ups.Data{
			"Material": d2d.NewMaterial2D(
				load.NewTexture(data.Get(data.AssetPath)+"gnome.png"),
				vec3.New(1.0, 1.0, 1.0),
				1.0,
			),
			"Transform": d2d.NewTransform2D(
				vec2.New(0.0, 1.25),
				vec2.All(2.0), 0.0,
			),
			"Offset": float32(15.0),

			"Min":   float32(-10.0),
			"Max":   float32(10.0),
			"Speed": float32(90.0),
		},
		[]ups.System{
			s2d.RenderRectangle2D{},
			Obsticle{},
			FlipRotate{},
		},
	).Clone(
		ups.Data{
			"Offset": float32(30.0),
			"Speed":  float32(70.0),
		},

		ups.Data{
			"Offset": float32(45.0),
			"Speed":  float32(90.0),
		},

		ups.Data{
			"Offset": float32(56.5),
			"Speed":  float32(85.0),
		},

		ups.Data{
			"Offset": float32(66.5),
			"Speed":  float32(77.0),
		},
	)

	WorldScene.New(
		"Ground",
		ups.Data{
			"Material": d2d.NewMaterial2D(
				load.EmptyTexture(),
				vec3.New(0.0, 1.0, 0.0),
				5.0,
			),
			"Transform": d2d.NewTransform2D(
				vec2.New(0, -1.21),
				vec2.New(36.6, 3.0),
				0.0,
			),
		},
		[]ups.System{
			s2d.RenderRectangle2D{},
			s2d.AABB{},
			s2d.StaticAABB{},
		},
	)

	music := load.NewAudio(data.Get(data.AssetPath) + "Alla-Turca(chosic.com).mp3")
	music.Play(-1)
	music.SetVolume(20)

	ups.SceneManager.Set(WorldScene.ID)

	for !window.CloseEvent() {
		window.BeginDraw()

		ups.SceneManager.Update(window.GetDeltaTime())

		window.EndDraw()
	}
}
