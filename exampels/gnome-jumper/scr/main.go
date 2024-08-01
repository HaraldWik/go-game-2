package main

import (
	"fmt"
	"math"
	"math/rand"

	d2d "github.com/HaraldWik/go-game-2/scr/2d/data"
	s2d "github.com/HaraldWik/go-game-2/scr/2d/systems"
	"github.com/HaraldWik/go-game-2/scr/app"
	"github.com/HaraldWik/go-game-2/scr/input"
	load "github.com/HaraldWik/go-game-2/scr/loaders"
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

var score uint32

func main() {
	app := app.New()
	window := app.NewWindow("Gnome Jumper 24", vec2.New(1920, 1075))
	window.Flags = window.FLAG_RESIZABLE
	window.MaxFPS = 55
	window.Open()

	sceneMain := ups.SceneManager.NewScene()
	secretScene := ups.SceneManager.NewScene()

	// Secret scene
	secretScene.NewObject(
		"Skybox",
		ups.Data{
			"Color": vec3.Zero(),
		},
		[]ups.System{
			s2d.Skybox2D{},
		},
	)

	secretScene.NewObject(
		"Secret",
		ups.Data{
			"Material":  d2d.NewMaterial2D(load.NewTexture("../assets/texture.png"), vec3.All(0.5), 1.0),
			"Transform": d2d.NewTransform2D(vec2.New(0.0, 0.0), vec2.New(36.75, 20.0), 0.0),
		},
		[]ups.System{
			s2d.RenderRectangle2D{},
		},
	)

	secretScene.NewObject(
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

	// MainScen

	sceneMain.NewObject(
		"BackroundGnome",
		ups.Data{
			"Material":  d2d.NewMaterial2D(load.NewTexture("../assets/gnome.png"), vec3.All(0.5), -9.0),
			"Transform": d2d.NewTransform2D(vec2.New(0.0, 0.0), vec2.New(16.0, 13.0), 25.0),

			"Min":   float32(-40.0),
			"Max":   float32(40.0),
			"Speed": float32(150.0),
		},
		[]ups.System{
			s2d.RenderRectangle2D{},
			FlipRotate{},
		},
	)

	sceneMain.NewObject(
		"Skybox",
		ups.Data{
			"Color": vec3.New(0.0, 0.25, 0.65),
		},
		[]ups.System{
			s2d.Skybox2D{},
		},
	)

	sceneMain.NewObject(
		"Camera2D",
		ups.Data{
			"Window":    window,
			"Transform": d2d.NewTransform2D(vec2.New(0.0, 7.5), vec2.All(1.0), 0.0),
			"Zoom":      float32(10),
		},
		[]ups.System{
			s2d.Camera2D{},
		},
	)

	sceneMain.NewObject(
		"Player",
		ups.Data{
			"Material": d2d.NewMaterial2D(
				load.EmptyTexture(),
				vec3.New(1.0, 0.0, 0.0),
				2.0,
			),
			"Transform": d2d.NewTransform2D(
				vec2.New(-6.0, 0.0),
				vec2.All(1.35), 0.0,
			),
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

			Controller2D{},
		},
	)

	sceneMain.NewObject(
		"Obsticle",
		ups.Data{
			"Material": d2d.NewMaterial2D(
				load.NewTexture("../assets/gnome.png"),
				vec3.New(1.0, 1.0, 1.0),
				1.0,
			),
			"Transform": d2d.NewTransform2D(vec2.New(0.0, 1.25), vec2.All(2.0), 0.0),
			"Offset":    float32(15.0),

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

	sceneMain.NewObject(
		"Ground",
		ups.Data{
			"Material": d2d.NewMaterial2D(
				load.EmptyTexture(),
				vec3.New(0.0, 1.0, 0.0),
				5.0,
			),
			"Transform": d2d.NewTransform2D(vec2.New(0, -1.21), vec2.New(40.0, 3.0), 0.0),
		},
		[]ups.System{
			s2d.RenderRectangle2D{},
		},
	)

	sceneMain.NewObject(
		"Manager",
		ups.Data{},
		[]ups.System{
			ObsticleManager{},
		},
	)

	music := load.NewAudio("../assets/Alla-Turca(chosic.com).mp3")
	music.Play(-1)
	music.SetVolume(20)

	ups.SceneManager.SetCurrentScenes(sceneMain.ID)

	var toggleMenu bool

	for !window.CloseEvent() {
		window.BeginDraw()

		if input.IsJustPressed(input.M_LMB) || !toggleMenu {
			ups.SceneManager.SetCurrentScenes(sceneMain.ID)
			toggleMenu = false
		}
		if input.IsJustPressed(input.M_RMB) || toggleMenu {
			ups.SceneManager.SetCurrentScenes(secretScene.ID)
			toggleMenu = true
		}
		if input.IsJustPressed(input.K_ESC) {
			toggleMenu = !toggleMenu
		}

		ups.SceneManager.Update(window.GetDeltaTime())

		window.EndDraw()
	}
}

type Rotate struct{}

func (r Rotate) Start(obj *ups.Object) {}

func (r Rotate) Update(obj *ups.Object, deltaTime float32) {
	transform := obj.Data.Get("Transform").(d2d.Transform2D)
	transform.Rotation += 200 * deltaTime
	obj.Data.Set("Transform", transform)
}

type FlipRotate struct{}

func (r FlipRotate) Start(obj *ups.Object) {
	obj.Data.Set("Flip", true)
}

func (r FlipRotate) Update(obj *ups.Object, deltaTime float32) {
	transform := obj.Data.Get("Transform").(d2d.Transform2D)
	flip := obj.Data.Get("Flip").(bool)
	min := obj.Data.Get("Min").(float32)
	max := obj.Data.Get("Max").(float32)
	speed := obj.Data.Get("Speed").(float32)

	if transform.Rotation >= max {
		flip = false
	} else if transform.Rotation <= min {
		flip = true
	}

	if flip {
		transform.Rotation += speed * deltaTime
	} else {
		transform.Rotation -= speed * deltaTime
	}

	if transform.Rotation >= 360 {
		transform.Rotation -= 360
	} else if transform.Rotation <= -360 {
		transform.Rotation += 360
	}

	obj.Data.Set("Transform", transform)
	obj.Data.Set("Flip", flip)
}

type ObsticleManager struct{}

func (m ObsticleManager) Start(obj *ups.Object) {
	var (
		player = obj.Scene.FindByTag("Player")[0]
	)

	obj.Data.Set("Obsticles", obj.Scene.FindByTag("Obsticle"))
	obj.Data.Set("Player", player)
	obj.Data.Set("HasDied", false)
	obj.Data.Set("CanDie", true)
}

func (m ObsticleManager) Update(obj *ups.Object, deltaTime float32) {
	var (
		obsticles = obj.Data.Get("Obsticles").([]*ups.Object)
		player    = obj.Data.Get("Player").(*ups.Object)
		hasDied   = obj.Data.Get("HasDied").(bool)
		canDie    = obj.Data.Get("CanDie").(bool)
	)

	if input.IsJustPressed(input.K_R) {
		obj.Data.Set("HasDied", false)
		obj.Data.Set("CanDie", true)
		obj.Scene.DeleteObject("You-died")
		obj.Scene.DeleteObject("Gnome")
	}

	for _, obs := range obsticles {
		playerTransform := player.Data.Get("Transform").(d2d.Transform2D)
		obsTransform := obs.Data.Get("Transform").(d2d.Transform2D)

		for _, otherObs := range obsticles {
			if otherObs.Name != obs.Name {
				otherObsTransform := otherObs.Data.Get("Transform").(d2d.Transform2D)

				distance := obsTransform.Position.X - otherObsTransform.Position.X

				calculatedDistance := float32(math.Abs(float64(distance)))

				if calculatedDistance < 0.55 || calculatedDistance > 1.0 && calculatedDistance < 6.0 {
					otherObsTransform.Position.X += float32(rand.Intn(10.0))
				}

				otherObs.Data.Set("Transform", otherObsTransform)
			}
		}

		if playerTransform.Position.X < obsTransform.Position.X+obsTransform.Scale.X-0.55 &&
			playerTransform.Position.X+playerTransform.Scale.X-0.55 > obsTransform.Position.X {
			if playerTransform.Position.Y < 1.8 {
				hasDied = true
				obj.Data.Set("HadDied", hasDied)
			} else {
				if !obs.Data.Get("hasBeenjumpedOver").(bool) && canDie {
					score++
					obs.Data.Set("hasBeenjumpedOver", true)
				}
			}
		}
	}

	if hasDied && canDie {
		fmt.Println(score)
		score = 0

		obj.Scene.NewObject(
			"You-died",
			ups.Data{
				"Material": d2d.NewMaterial2D(
					load.NewTexture("../assets/you_died.png"),
					vec3.All(1.0), 9.0,
				),
				"Transform": d2d.NewTransform2D(
					vec2.New(0.0, 7.0),
					vec2.All(37.5), 0.0,
				),
			},
			[]ups.System{
				s2d.RenderRectangle2D{},
			},
		)

		obj.Scene.NewObject(
			"Gnome",
			ups.Data{
				"Material": d2d.NewMaterial2D(
					load.NewTexture("../assets/gnome.png"),
					vec3.New(1.0, 0.5, 0.5),
					8.0,
				),
				"Transform": d2d.NewTransform2D(
					vec2.New(0.0, 13.0),
					vec2.New(8.0, 5.0),
					25.0,
				),
			},
			[]ups.System{
				s2d.RenderRectangle2D{},
				Rotate{},
			},
		)

		obj.Data.Set("CanDie", false)
	}
}

type Obsticle struct{}

func (o Obsticle) Start(obj *ups.Object) {
	var (
		transform = obj.Data.Get("Transform").(d2d.Transform2D)
		offset    = obj.Data.Get("Offset").(float32)
	)

	transform.Position.X = offset

	obj.Data.Set("Transform", transform)
	obj.Data.Set("hasBeenjumpedOver", false)

	obj.Tags.Add("Obsticle")
}

func (o Obsticle) Update(obj *ups.Object, deltaTime float32) {
	var (
		transform = obj.Data.Get("Transform").(d2d.Transform2D)
	)

	if transform.Position.X < -20.0 {
		transform.Position.X = 20.0 + float32(rand.Intn(50))
		obj.Data.Set("hasBeenjumpedOver", false)
	} else {
		transform.Position.X -= 10 * deltaTime
	}

	obj.Data.Set("Transform", transform)
}

type Controller2D struct{}

func (c Controller2D) Start(obj *ups.Object) {
	obj.Data.Set("Jumped", true)
}

func (c Controller2D) Update(obj *ups.Object, deltaTime float32) {
	var (
		transform = obj.Data.Get("Transform").(d2d.Transform2D)

		jumped = obj.Data.Get("Jumped").(bool)
	)

	const (
		GRAVITY         = 8.5
		JUMP_VELOCITY   = 20
		MAX_JUMP_HEIGHT = 5.0
		GROUND_LEVEL    = 1.0
	)

	if input.IsPressed(input.K_SPACE) || input.IsPressed(input.K_UP) || input.IsPressed(input.K_W) && transform.Position.Y+transform.Scale.Y > GROUND_LEVEL+0.4 {
		if !jumped {
			transform.Position.Y += JUMP_VELOCITY * deltaTime
		}
	}

	if input.IsJustReleased(input.K_SPACE) || input.IsJustReleased(input.K_UP) || input.IsJustReleased(input.K_W) && transform.Position.Y > GROUND_LEVEL+0.4 {
		jumped = true
	}

	if transform.Position.Y > MAX_JUMP_HEIGHT {
		jumped = true
	}

	if transform.Position.Y < GROUND_LEVEL+0.4 {
		jumped = false
	}

	if transform.Position.Y < MAX_JUMP_HEIGHT-GROUND_LEVEL && jumped {
		transform.Position.Y -= transform.Position.Y/15 - GROUND_LEVEL*deltaTime
	}

	if transform.Position.Y > transform.Scale.Y && transform.Position.Y > 1 {
		transform.Position.Y -= GRAVITY * deltaTime
	} else {
		transform.Position.Y = 1.0
	}

	obj.Data.Set("Transform", transform)
	obj.Data.Set("Jumped", jumped)
}

// Helper functions!
func Lerp(a, b, t float32) float32 {
	return a + t*(b-a)
}
