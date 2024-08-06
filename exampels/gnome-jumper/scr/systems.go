package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/HaraldWik/go-game-2/exampels/gnome-jumper/scr/data"
	d2d "github.com/HaraldWik/go-game-2/scr/2d/data"
	s2d "github.com/HaraldWik/go-game-2/scr/2d/systems"
	"github.com/HaraldWik/go-game-2/scr/input"
	load "github.com/HaraldWik/go-game-2/scr/loaders"
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
	vec4 "github.com/HaraldWik/go-game-2/scr/vector/4"
)

type ColorChange struct{}

func (c ColorChange) Start(obj *ups.Object) {
	obj.Data.Set("Time", float32(0.0))
	obj.Data.Set("Toggle", false)
}

func (c ColorChange) FixedUpdate(obj *ups.Object) {
	var (
		material = obj.Data.Get("Material").(d2d.Material2D)
		time     = obj.Data.Get("Time").(float32)
		toggle   = obj.Data.Get("Toggle").(bool)
	)

	if input.IsJustPressed(input.K_DELETE) {
		toggle = !toggle
	}

	if !toggle {
		material.Opacity = 0.0
	} else {
		time += 0.05

		r := float32((math.Sin(float64(time)) + 1) / 2)
		g := float32((math.Sin(float64(time)+2*math.Pi/3) + 1) / 2)
		b := float32((math.Sin(float64(time)+4*math.Pi/3) + 1) / 2)

		material.Color.X = r
		material.Color.Y = g
		material.Color.Z = b
		material.Opacity = float32(math.Abs(float64(r/2.0+g/2.0-b/4.0)) * math.Pi / 2.0)

		if material.Opacity > 0.985 {
			material.Opacity = 0.985
		}
		if material.Opacity < 0.5 {
			material.Opacity = 0.5
		}
	}

	obj.Data.Set("Material", material)
	obj.Data.Set("Time", time)
	obj.Data.Set("Toggle", toggle)
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
	var (
		transform = obj.Data.Get("Transform").(d2d.Transform2D)
		flip      = obj.Data.Get("Flip").(bool)
		min       = obj.Data.Get("Min").(float32)
		max       = obj.Data.Get("Max").(float32)
		speed     = obj.Data.Get("Speed").(float32)
	)

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

type Obsticle struct{}

func (o Obsticle) Start(obj *ups.Object) {
	var (
		transform = obj.Data.Get("Transform").(d2d.Transform2D)
		offset    = obj.Data.Get("Offset").(float32)
	)

	transform.Position.X = offset

	obj.Data.Set("Transform", transform)
	obj.Data.Set("HasBeenJumpedOver", false)

	obj.Tags.Add("Obsticle")
}

func (o Obsticle) Update(obj *ups.Object, deltaTime float32) {
	var (
		transform         = obj.Data.Get("Transform").(d2d.Transform2D)
		killables         = obj.Scene.GetByTag("Killable")
		hasBeenjumpedOver = obj.Data.Get("HasBeenJumpedOver").(bool)
	)

	if transform.Position.X < -20.0 {
		transform.Position.X = 20.0 + float32(rand.Intn(80))
		hasBeenjumpedOver = false
	} else {
		transform.Position.X -= 10 * deltaTime
	}

	// Check for collisions with killable objects
	for _, killable := range killables {
		killableTransform := killable.Data.Get("Transform").(d2d.Transform2D)

		for _, obstacle := range obj.Scene.GetByTag("Obsticle") {
			if obstacle.Name != obj.Name {
				obstacleTransform := obstacle.Data.Get("Transform").(d2d.Transform2D)

				// Calculate the distance between the obstacle and other obstacles
				distance := obstacleTransform.Position.X - transform.Position.X
				calculatedDistance := float32(math.Abs(float64(distance)))

				if calculatedDistance < 0.55 || (calculatedDistance > 1.0 && calculatedDistance < 6.0) {
					obstacleTransform.Position.X += float32(rand.Intn(90))
				}

				obstacle.Data.Set("Transform", obstacleTransform)
			}
		}

		// Check for collision between killable objects and the obstacle
		if killableTransform.Position.X < transform.Position.X+transform.Scale.X-0.55 &&
			killableTransform.Position.X+killableTransform.Scale.X-0.55 > transform.Position.X &&
			killableTransform.Position.Y < 1.8 {
			killable.Tags.Add("Die")
		} else if !hasBeenjumpedOver && !killable.Tags.Has("Die") {
			hasBeenjumpedOver = true

			score++
			data.Set(data.Score, score)

			if score > data.GetAsUint32(data.HighScore) {
				data.Set(data.HighScore, score)
			}
		}
	}
	obj.Data.Set("Transform", transform)
	obj.Data.Set("HasBeenJumpedOver", hasBeenjumpedOver)
}

type Death struct{}

func (c Death) Start(obj *ups.Object) {
	obj.Tags.Add("Killable")
}

func (d Death) Update(obj *ups.Object, deltaTime float32) {
	DeathScreen := ups.SceneManager.New()

	if input.IsJustPressed(input.K_R) {
		obj.Tags.Remove("Die")
		obj.Tags.Add("Killable")
		DeathScreen.Delete("YouDiedText")
		DeathScreen.Delete("Gnome")
		ups.SceneManager.Remove(DeathScreen.ID)
		ups.SceneManager.Add(WorldScene.ID)
	}

	if obj.Tags.Has("Die") && obj.Tags.Has("Killable") {
		score = 0
		obj.Tags.Remove("Killable")
		ups.SceneManager.Add(DeathScreen.ID)

		DeathScreen.New(
			"YouDiedText",
			ups.Data{
				"Material": d2d.NewMaterial2D(
					load.NewTexture(data.Get(data.AssetPath)+"you_died.png"),
					vec4.Zero(),
					vec3.All(1.0),
					1.0,
					10.0,
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

		DeathScreen.New(
			"Gnome",
			ups.Data{
				"Material": d2d.NewMaterial2D(
					load.NewTexture(data.Get(data.AssetPath)+"gnome.png"),
					vec4.Zero(),
					vec3.New(1.0, 0.5, 0.5),
					1.0,
					10.0,
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
	}
}

type Controller2D struct{}

func (c Controller2D) Start(obj *ups.Object) {
	obj.Data.Set("Jumped", true)
	obj.Data.Set("GoDown", false)
}

func (c Controller2D) Update(obj *ups.Object, deltaTime float32) {
	var (
		transform = obj.Data.Get("Transform").(d2d.Transform2D)

		jumped = obj.Data.Get("Jumped").(bool)
		goDown = obj.Data.Get("GoDown").(bool)
	)

	if transform.Position.Y > transform.Scale.Y/1.5 {
		if !goDown {
			transform.Position.Y -= 2.5 * deltaTime
		} else {
			transform.Position.Y -= 9.8 * deltaTime
		}
	}

	if input.IsPressed(input.K_SPACE) && !jumped {
		jumped = true
	}

	if jumped && !goDown {
		transform.Position.Y += 15.0 * deltaTime
	}

	if transform.Position.Y > 4.5 {
		goDown = true
	}

	if transform.Position.Y < 1.0 {
		jumped = false
		goDown = false
	}

	if input.IsPressed(input.K_0) {
		fmt.Println(transform.Position.Y)
	}

	obj.Data.Set("Transform", transform)
	obj.Data.Set("Jumped", jumped)
	obj.Data.Set("GoDown", goDown)
}

// Gui

// Helper functions!
func Lerp(a, b, t float32) float32 {
	return a + t*(b-a)
}
