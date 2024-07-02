package component

import (
	"github.com/HaraldWik/go-game-2/scr/app"
	"github.com/go-gl/gl/v2.1/gl"
)

type Camera2D struct {
	Window    app.Win
	Transform Transform2D
	Zoom      float32
}

func (camera *Camera2D) Update() {
	gl.Viewport(0, 0, int32(camera.Window.Size.X), int32(camera.Window.Size.Y))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	x, y := camera.Window.SDL.GetSize()
	// Calculate aspect ratio
	aspect := float32(x) / float32(y)
	left := -1*camera.Zoom*aspect + camera.Transform.Pos.X
	right := 1*camera.Zoom*aspect + camera.Transform.Pos.X
	top := 1*camera.Zoom + camera.Transform.Pos.Y
	bottom := -1*camera.Zoom + camera.Transform.Pos.Y
	gl.Ortho(float64(left), float64(right), float64(bottom), float64(top), -3, 3)

	// Rotation & position
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(float32(camera.Transform.Pos.X), float32(camera.Transform.Pos.Y), 0)
	gl.Rotatef(float32(camera.Transform.Rot), 0, 0, 1)
	gl.Translatef(-float32(camera.Transform.Pos.X), -float32(camera.Transform.Pos.Y), 0)
}
