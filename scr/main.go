/*package main

import (
	"log"
	"math"
	"runtime"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
)

type Vec2 struct {
	X, Y float32
}

type Vec3 struct {
	X, Y, Z float32
}

type Camera struct {
	Position Vec3
	Rotation Vec3
	FOV      float64
	Aspect   float64
	Near     float64
	Far      float64
}

type Win struct {
	Name  string
	Size  Vec2
	Flags uint32
	sdl   *sdl.Window
}

func (win *Win) Open() {
	runtime.LockOSThread()

	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Fatalf("Failed to initialize SDL: %v", err)
	}
	defer sdl.Quit()

	var err error
	win.sdl, err = sdl.CreateWindow(win.Name, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(win.Size.X), int32(win.Size.Y), win.Flags|sdl.WINDOW_OPENGL|sdl.WINDOW_RESIZABLE)
	if err != nil {
		log.Fatalf("Failed to create window: %v", err)
	}
	defer win.sdl.Destroy()

	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 2)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)

	_, err = win.sdl.GLCreateContext()
	if err != nil {
		log.Fatalf("Failed to create OpenGL context: %v", err)
	}

	if err := gl.Init(); err != nil {
		log.Fatalf("Failed to initialize OpenGL: %v", err)
	}

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.ClearColor(0.1, 0.2, 0.3, 1.0)

	win.mainLoop()
}

func (win *Win) mainLoop() {
	lastTime := time.Now()
	angle := 0.0

	camera := Camera{
		Position: Vec3{0, 0, 5},
		Rotation: Vec3{0, 0, 0},
		FOV:      45,
		Aspect:   float64(win.Size.X) / float64(win.Size.Y),
		Near:     0.1,
		Far:      100,
	}

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.WindowEvent:
				if e.Event == sdl.WINDOWEVENT_RESIZED {
					width, height := win.sdl.GetSize()
					camera.Aspect = float64(width) / float64(height)
					gl.Viewport(0, 0, width, height)
				}
			}
		}

		now := time.Now()
		elapsed := now.Sub(lastTime).Seconds()
		lastTime = now

		angle += 90.0 * elapsed // Rotate at 90 degrees per second

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		camera.Setup()

		win.renderCube(float32(angle))

		win.sdl.GLSwap()
	}
}

func (camera *Camera) Setup() {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gluPerspective(camera.FOV, camera.Aspect, camera.Near, camera.Far)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Rotatef(camera.Rotation.X, 1, 0, 0)
	gl.Rotatef(camera.Rotation.Y, 0, 1, 0)
	gl.Rotatef(camera.Rotation.Z, 0, 0, 1)
	gl.Translatef(-camera.Position.X, -camera.Position.Y, -camera.Position.Z)
}

func (win *Win) renderCube(angle float32) {
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(0, 0, -10) // Move the cube further back
	gl.Rotatef(angle, 1, 1, 1)

	gl.Begin(gl.QUADS)

	// Front face
	gl.Color3f(1.0, 0.0, 0.0)
	gl.Vertex3f(-1.0, -1.0, 1.0)
	gl.Vertex3f(1.0, -1.0, 1.0)
	gl.Vertex3f(1.0, 1.0, 1.0)
	gl.Vertex3f(-1.0, 1.0, 1.0)

	// Back face
	gl.Color3f(0.0, 1.0, 0.0)
	gl.Vertex3f(-1.0, -1.0, -1.0)
	gl.Vertex3f(-1.0, 1.0, -1.0)
	gl.Vertex3f(1.0, 1.0, -1.0)
	gl.Vertex3f(1.0, -1.0, -1.0)

	// Top face
	gl.Color3f(0.0, 0.0, 1.0)
	gl.Vertex3f(-1.0, 1.0, -1.0)
	gl.Vertex3f(-1.0, 1.0, 1.0)
	gl.Vertex3f(1.0, 1.0, 1.0)
	gl.Vertex3f(1.0, 1.0, -1.0)

	// Bottom face
	gl.Color3f(1.0, 1.0, 0.0)
	gl.Vertex3f(-1.0, -1.0, -1.0)
	gl.Vertex3f(1.0, -1.0, -1.0)
	gl.Vertex3f(1.0, -1.0, 1.0)
	gl.Vertex3f(-1.0, -1.0, 1.0)

	// Right face
	gl.Color3f(1.0, 0.0, 1.0)
	gl.Vertex3f(1.0, -1.0, -1.0)
	gl.Vertex3f(1.0, 1.0, -1.0)
	gl.Vertex3f(1.0, 1.0, 1.0)
	gl.Vertex3f(1.0, -1.0, 1.0)

	// Left face
	gl.Color3f(0.0, 1.0, 1.0)
	gl.Vertex3f(-1.0, -1.0, -1.0)
	gl.Vertex3f(-1.0, -1.0, 1.0)
	gl.Vertex3f(-1.0, 1.0, 1.0)
	gl.Vertex3f(-1.0, 1.0, -1.0)

	gl.End()
}

func gluPerspective(fovY, aspect, zNear, zFar float64) {
	fH := math.Tan(fovY/360*math.Pi) * zNear
	fW := fH * aspect
	gl.Frustum(-fW, fW, -fH, fH, zNear, zFar)
}

func main() {
	win := Win{
		Name:  "OpenGL with SDL2 in Go",
		Size:  Vec2{X: 800, Y: 600},
		Flags: sdl.WINDOW_SHOWN,
	}
	win.Open()
}
*/

package main

import (
	debug "github.com/HaraldWik/debug/scr"
	"github.com/HaraldWik/go-game-2/scr/app"
	component "github.com/HaraldWik/go-game-2/scr/components"

	"github.com/HaraldWik/go-game-2/scr/abus"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

func main() {
	app := app.NewApp()
	win := app.NewWindow("Window-1", vec2.New(500, 400))
	win.Flags = win.FLAG_RESIZABLE
	win.Open()

	scene := abus.NewScene()

	scene.Create(component.Cam2D.New(component.Cam2D{}, win, component.Transform2D{}, 25.0))

	for i := 0; i < 300; i++ {
		scene.Create(&component.RenderRect2D{component.Transform2D{Pos: vec2.New(float32(i-50), float32(i-25)), Size: vec2.All(1.0), Rot: float32(i * 10)}, vec3.New(float32(i)*0.01, 0.0, float32(i)*0.1)})
		scene.Create(&component.RenderRect2D{component.Transform2D{Pos: vec2.New(float32(i-20), float32(i-25)), Size: vec2.All(1.0), Rot: float32(i * 15)}, vec3.New(float32(i)*0.01, 0.0, 0.0)})
		scene.Create(&component.RenderRect2D{component.Transform2D{Pos: vec2.New(float32(i-70), float32(i-25)), Size: vec2.All(1.0), Rot: float32(i * 15)}, vec3.New(0.0, float32(i)*0.01, 0.0)})
		//scene.Create(&component.RenderCircle2D{component.Transform2D{Pos: vec2.New(float32(i-7), float32(i-5)), Size: vec2.All(1.0), Rot: float32(i * 15)}, vec3.New(0.0, float32(i)*0.1, 0.0)})
	}

	t := vec2.New(0.0, 20.0)
	scene.Create(&component.RenderRect2D{Transform: component.Transform2D{Pos: t}, Color: vec3.Zero()})

	debug.Info("Hello")

	for !win.CloseEvent() {
		win.Update()
		win.BeginDraw(0.1, 0.2, 0.7)

		scene.Update()

		win.EndDraw(60)
	}
}
