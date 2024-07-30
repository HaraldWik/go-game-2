package main

import (
	"math"
	"math/rand"

	"github.com/HaraldWik/go-game-2/scr/app"
	"github.com/HaraldWik/go-game-2/scr/audio"
	dt "github.com/HaraldWik/go-game-2/scr/data-types"
	"github.com/HaraldWik/go-game-2/scr/input"
	sys "github.com/HaraldWik/go-game-2/scr/systems"
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

func main() {
	app := app.New()
	window := app.NewWindow("Gnome Jumper 24", vec2.New(1920, 1075))
	window.Flags = window.FLAG_RESIZABLE
	window.Open()

	ups.NewObject(
		"Camera2D",
		ups.Data{
			"Window":    window,
			"Transform": dt.NewTransform2D(vec2.New(0.0, 7.5), vec2.All(1.0), 0.0),
			"Zoom":      float32(10),
		},
		[]ups.System{
			sys.Camera2D{},
		},
	)

	ups.NewObject(
		"Player",
		ups.Data{
			"Color":     vec3.New(1.0, 0.5, 0.0),
			"Transform": dt.NewTransform2D(vec2.New(-10.0, 5), vec2.All(1.0), 0.0),
			"Segments":  uint32(30),
			"Speed":     float32(900.0),
		},
		[]ups.System{
			sys.RenderCircle2D{},

			Controller2D{},
		},
		"Player",
	)

	ups.NewObject(
		"Obstical",
		ups.Data{
			"Material":  dt.NewMaterial(dt.LoadTexture("../assets/gnome.png"), vec3.New(1.0, 1.0, 1.0)),
			"Transform": dt.NewTransform2D(vec2.New(0.0, 1.25), vec2.All(2.0), 0.0),
			"Offset":    float32(15.0),
		},
		[]ups.System{
			sys.RenderTexture2D{},
			Obstical{},
		},
	).Clone(
		4,
		ups.Data{
			"Material":  dt.NewMaterial(dt.LoadTexture("../assets/gnome.png"), vec3.New(1.0, 1.0, 1.0)),
			"Transform": dt.NewTransform2D(vec2.New(0.0, 1.25), vec2.All(2.0), 0.0),
			"Offset":    float32(30.0),
		},

		ups.Data{
			"Material":  dt.NewMaterial(dt.LoadTexture("../assets/gnome.png"), vec3.New(1.0, 1.0, 1.0)),
			"Transform": dt.NewTransform2D(vec2.New(0.0, 1.25), vec2.All(2.0), 0.0),
			"Offset":    float32(45.0),
		},

		ups.Data{
			"Material":  dt.NewMaterial(dt.LoadTexture("../assets/gnome.png"), vec3.New(1.0, 1.0, 1.0)),
			"Transform": dt.NewTransform2D(vec2.New(0.0, 1.25), vec2.All(2.0), 0.0),
			"Offset":    float32(56.5),
		},

		ups.Data{
			"Material":  dt.NewMaterial(dt.LoadTexture("../assets/gnome.png"), vec3.New(1.0, 1.0, 1.0)),
			"Transform": dt.NewTransform2D(vec2.New(0.0, 1.25), vec2.All(2.0), 0.0),
			"Offset":    float32(66.5),
		},
	)

	ups.NewObject(
		"Ground",
		ups.Data{
			"Color":     vec3.New(0.0, 1.0, 0.0),
			"Transform": dt.NewTransform2D(vec2.New(0, -1.25), vec2.New(40.0, 3.0), 0.0),
		},
		[]ups.System{
			sys.RenderRectangle2D{},
		},
	)

	ups.NewObject(
		"Manager",
		ups.Data{},
		[]ups.System{
			ObsticalManager{},
		},
	)

	music := audio.LoadMP3("../assets/Alla-Turca(chosic.com).mp3")
	music.Play(-1)
	music.SetVolume(20)

	for !window.CloseEvent() {
		window.BeginDraw(vec3.New(0.0, 0.144, 0.856))

		ups.Engine.Update(window.GetDeltaTime())

		window.EndDraw(60)
	}
}

type ObsticalManager struct{}

func (m ObsticalManager) Start() {
	var (
		obj    = ups.Engine.GetParent()
		player = ups.Engine.FindTag("Player")[0]
	)

	obj.Data.Set("Obsticals", ups.Engine.FindTag("Obstical"))
	obj.Data.Set("Player", player)
	obj.Data.Set("HasDied", false)
	obj.Data.Set("CanDie", true)
}

func (m ObsticalManager) Update(deltaTime float32) {
	var (
		obj = ups.Engine.GetParent()

		obsticals = obj.Data.Get("Obsticals").([]*ups.Object)
		player    = obj.Data.Get("Player").(*ups.Object)
		hasDied   = obj.Data.Get("HasDied").(bool)
		canDie    = obj.Data.Get("CanDie").(bool)
	)

	for _, obs := range obsticals {
		playerTransform := player.Data.Get("Transform").(dt.Transform2D)
		obsTransform := obs.Data.Get("Transform").(dt.Transform2D)

		for _, otherObs := range obsticals {
			if otherObs.Name != obs.Name {
				otherObsTransform := otherObs.Data.Get("Transform").(dt.Transform2D)

				distance := obsTransform.Pos.X - otherObsTransform.Pos.X

				calculatedDistance := float32(math.Abs(float64(distance)))

				if calculatedDistance < 0.55 || calculatedDistance > 1.0 && calculatedDistance < 6.0 {
					otherObsTransform.Pos.X += float32(rand.Intn(10.0))
				}

				otherObs.Data.Set("Transform", otherObsTransform)
			}
		}

		if playerTransform.Pos.X < obsTransform.Pos.X+obsTransform.Size.X-0.55 &&
			playerTransform.Pos.X+playerTransform.Size.X-0.55 > obsTransform.Pos.X &&
			playerTransform.Pos.Y < 1.8 {
			hasDied = true
			obj.Data.Set("HadDied", hasDied)
		}
	}

	if hasDied && canDie {
		ups.NewObject(
			"You-died",
			ups.Data{
				"Material":  dt.NewMaterial(dt.LoadTexture("../assets/you_died.png"), vec3.All(1.0)),
				"Transform": dt.NewTransform2D(vec2.New(0.0, 7.0), vec2.All(37.5), 0.0),
			},
			[]ups.System{
				sys.RenderTexture2D{},
			},
		)

		ups.NewObject(
			"Gnome",
			ups.Data{
				"Material":  dt.NewMaterial(dt.LoadTexture("../assets/gnome.png"), vec3.New(1.0, 0.5, 0.5)),
				"Transform": dt.NewTransform2D(vec2.New(0.0, 13.0), vec2.New(8.0, 5.0), 25.0),
			},
			[]ups.System{
				sys.RenderTexture2D{},
				Rotate{},
			},
		)

		obj.Data.Set("CanDie", false)
	}
}

type Rotate struct{}

func (r Rotate) Start() {}

func (o Rotate) Update(deltaTime float32) {
	var (
		obj       = ups.Engine.GetParent()
		transform = obj.Data.Get("Transform").(dt.Transform2D)
	)

	transform.Rot += 200 * deltaTime

	obj.Data.Set("Transform", transform)
}

type Obstical struct{}

func (o Obstical) Start() {
	var (
		obj       = ups.Engine.GetParent()
		transform = obj.Data.Get("Transform").(dt.Transform2D)
		offset    = obj.Data.Get("Offset").(float32)
	)

	transform.Pos.X += offset

	obj.Data.Set("Transform", transform)
}

func (o Obstical) Update(deltaTime float32) {
	var (
		obj       = ups.Engine.GetParent()
		transform = obj.Data.Get("Transform").(dt.Transform2D)
	)

	if transform.Pos.X < -20.0 {
		transform.Pos.X = 20.0 + float32(rand.Intn(50))
	} else {
		transform.Pos.X -= 10 * deltaTime
	}

	obj.Data.Set("Transform", transform)
}

type Controller2D struct{}

func (c Controller2D) Start() {
	var (
		obj = ups.Engine.GetParent()
	)

	obj.Data.Set("Jumped", true)
}

func (c Controller2D) Update(deltaTime float32) {
	var (
		obj       = ups.Engine.GetParent()
		transform = obj.Data.Get("Transform").(dt.Transform2D)
		speed     = obj.Data.Get("Speed").(float32)

		jumped = obj.Data.Get("Jumped").(bool)
	)

	const (
		GRAVITY         = 8.5
		JUMP_VELOCITY   = 20
		MAX_JUMP_HEIGHT = 5.0
		GROUND_LEVEL    = 1.0
	)

	if input.IsPressed(input.K_SPACE) || input.IsPressed(input.K_UP) || input.IsPressed(input.K_W) && transform.Pos.Y+transform.Size.Y > GROUND_LEVEL+0.1 {
		if !jumped {
			transform.Pos.Y += JUMP_VELOCITY * deltaTime
		}
	}

	if input.IsJustReleased(input.K_SPACE) || input.IsJustReleased(input.K_UP) || input.IsJustReleased(input.K_W) && transform.Pos.Y > GROUND_LEVEL+0.1 {
		jumped = true
	}

	if transform.Pos.Y > MAX_JUMP_HEIGHT {
		jumped = true
	}

	if transform.Pos.Y < GROUND_LEVEL+0.1 {
		jumped = false
	}

	if transform.Pos.Y < MAX_JUMP_HEIGHT-GROUND_LEVEL && jumped {
		transform.Pos.Y -= transform.Pos.Y/15 - GROUND_LEVEL*deltaTime
	}

	if transform.Pos.Y > transform.Size.Y && transform.Pos.Y > 1 {
		transform.Pos.Y -= GRAVITY * deltaTime
	} else {
		transform.Pos.Y = 1.0
	}

	obj.Data.Set("Speed", speed)
	obj.Data.Set("Transform", transform)
	obj.Data.Set("Jumped", jumped)
}

// Helper functions!
func Lerp(a, b, t float32) float32 {
	return a + t*(b-a)
}
