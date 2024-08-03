package main

import (
	"math"

	d2d "github.com/HaraldWik/go-game-2/scr/2d/data"
	s2d "github.com/HaraldWik/go-game-2/scr/2d/systems"
	"github.com/HaraldWik/go-game-2/scr/app"
	"github.com/HaraldWik/go-game-2/scr/input"
	load "github.com/HaraldWik/go-game-2/scr/loaders"
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

var (
	MainScene = ups.SceneManager.New()
)

func main() {
	app := app.New()
	window := app.NewWindow("Tired Tires", vec2.New(1920, 1075))
	window.SetFlags(window.FLAG_RESIZABLE)
	window.SetMaxFPS(100)
	window.Open()

	MainScene.New(
		"Camera2D",
		ups.Data{
			"Window": window,
			"Transform": d2d.NewTransform2D(
				vec2.Zero(),
				vec2.Zero(),
				0.0,
			),
			"Zoom": float32(10),
		},
		[]ups.System{
			s2d.Camera2D{},
			FollowPlayer{},
		},
	)

	MainScene.New(
		"Skybox",
		ups.Data{
			"Color": vec3.New(0.0, 0.144, 0.856),
		},
		[]ups.System{
			s2d.Skybox2D{},
		},
	)

	MainScene.New(
		"Player",
		ups.Data{
			"Material": d2d.NewMaterial2D(
				load.EmptyTexture(),
				vec3.All(1.0),
				1.0,
			),
			"Transform": d2d.NewTransform2D(vec2.New(-10.0, 5), vec2.All(1.0), 0.0),
		},
		[]ups.System{
			s2d.RenderTriangle2D{},
			CarController{},
			s2d.AABB{},
		},
	)

	MainScene.New(
		"Thing",
		ups.Data{
			"Material": d2d.NewMaterial2D(
				load.EmptyTexture(),
				vec3.New(1.0, 0.0, 0.0),
				1.0,
			),
			"Transform": d2d.NewTransform2D(vec2.New(0.0, 6.0), vec2.All(1.0), 0.0),
		},
		[]ups.System{
			s2d.RenderRectangle2D{},
			s2d.AABB{},
			s2d.StaticAABB{},
		},
	)

	MainScene.New(
		"Thing2",
		ups.Data{
			"Material": d2d.NewMaterial2D(
				load.EmptyTexture(),
				vec3.New(1.0, 0.0, 0.0),
				1.0,
			),
			"Transform": d2d.NewTransform2D(vec2.New(-2.0, 6.0), vec2.All(1.0), 0.0),
			"Vertices": []vec2.Type{
				vec2.New(1.000000, 0.000000),
				vec2.New(0.923880, 0.382683),
				vec2.New(0.707107, 0.707107),
				vec2.New(0.382683, 0.923880),
				vec2.New(0.000000, 1.000000),
				vec2.New(-0.382683, 0.923880),
				vec2.New(-0.707107, 0.707107),
				vec2.New(-0.923880, 0.382683),
				vec2.New(-1.000000, 0.000000),
				vec2.New(-0.923880, -0.382683),
				vec2.New(-0.707107, -0.707107),
				vec2.New(-0.382683, -0.923880),
				vec2.New(0.000000, -1.000000),
				vec2.New(0.382683, -0.923880),
				vec2.New(0.707107, -0.707107),
				vec2.New(0.923880, -0.382683),
			},
		},
		[]ups.System{
			s2d.RenderMesh2D{},
			s2d.AABB{},
			s2d.StaticAABB{},
		},
	)

	ups.SceneManager.Set(MainScene.ID)

	for !window.CloseEvent() {
		window.BeginDraw()

		ups.SceneManager.Update(window.GetDeltaTime())

		window.EndDraw()
	}
}

type CarController struct{}

func (c CarController) Start(obj *ups.Object) {
	var (
		transform = obj.Data.Get("Transform").(d2d.Transform2D)
	)

	obj.Data.Set("Velocity", transform)
	obj.Data.Set("Acceleration", float32(0))

	engineSound := load.NewAudio("../assets/car_engine.wav")
	engineSound.SetVolume(5)
	engineSound.Play(-1)
}

func (c CarController) Update(obj *ups.Object, deltaTime float32) {
	var (
		transform    = obj.Data.Get("Transform").(d2d.Transform2D)
		velocity     = obj.Data.Get("Velocity").(d2d.Transform2D)
		acceleration = obj.Data.Get("Acceleration").(float32)
	)

	const (
		ACCELERATION_SPEED = 2.0
		DECELERATION_SPEED = 6.5
		MAX_ACCELERATION   = 15.0

		ROTATION_SPEED = 180.0
	)

	// Movement
	if input.IsPressed(input.K_W) {
		if acceleration < MAX_ACCELERATION {
			acceleration += ACCELERATION_SPEED * deltaTime
		}
	} else {
		if acceleration > 0.0 {
			acceleration -= DECELERATION_SPEED * deltaTime
		}
	}

	if input.IsPressed(input.K_S) && acceleration > 0.1 {
		acceleration -= DECELERATION_SPEED * deltaTime
	}

	velocity.Translate(vec2.New(0.0, acceleration*deltaTime))

	// Movement rounding
	if !input.IsPressed(input.K_W) && !input.IsPressed(input.K_S) {
		if acceleration < 0.2 && acceleration > -0.2 {
			acceleration = 0.0
		}
	}

	// Rotation

	if input.IsPressed(input.K_A) && !input.IsPressed(input.K_D) && acceleration != 0 {
		velocity.Rotation -= ROTATION_SPEED * deltaTime

		velocity.Translate(vec2.New(-acceleration*1.5*deltaTime, acceleration/3*deltaTime))

		acceleration -= DECELERATION_SPEED / 5 * deltaTime
	}
	if input.IsPressed(input.K_D) && !input.IsPressed(input.K_A) && acceleration != 0 {
		velocity.Rotation += ROTATION_SPEED * deltaTime

		velocity.Translate(vec2.New(acceleration*1.5*deltaTime, acceleration/3*deltaTime))

		acceleration -= DECELERATION_SPEED / 5 * deltaTime
	}

	lerpStrength := float32(3)

	transform.Position.X = Lerp(transform.Position.X, velocity.Position.X, lerpStrength*deltaTime)
	transform.Position.Y = Lerp(transform.Position.Y, velocity.Position.Y, lerpStrength*deltaTime)
	transform.Rotation = velocity.Rotation

	obj.Data.Set("Transform", transform)
	obj.Data.Set("Velocity", velocity)
	obj.Data.Set("Acceleration", acceleration)
}

type FollowPlayer struct{}

func (f FollowPlayer) Start(obj *ups.Object) {}

func (f FollowPlayer) Update(obj *ups.Object, deltaTime float32) {
	var (
		transform = obj.Data.Get("Transform").(d2d.Transform2D)

		player       = obj.Scene.GetByTag("Player")[0]
		target       = player.Data.Get("Transform").(d2d.Transform2D)
		acceleration = player.Data.Get("Acceleration").(float32)
	)

	target.Translate(vec2.New(0.0, float32(math.Abs(float64(acceleration)))/3))

	transform = target

	obj.Data.Set("Transform", transform)
}

func Lerp(start, end, t float32) float32 {
	return start + t*(end-start)
}
