package mod

import (
	"github.com/HaraldWik/go-game-2/scr/app"
	"github.com/go-gl/gl/v2.1/gl"
)

type Cam2D struct {
	Win       app.Win
	Transform Transform2D
	Zoom      float32
}

func (cam *Cam2D) Update() {
	gl.Viewport(0, 0, int32(cam.Win.GetSize().X), int32(cam.Win.GetSize().Y))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	// Calculate aspect ratio
	aspect := cam.Win.GetSize().X / cam.Win.GetSize().Y
	left := -1*cam.Zoom*aspect + cam.Transform.Pos.X
	right := 1*cam.Zoom*aspect + cam.Transform.Pos.X
	top := 1*cam.Zoom + cam.Transform.Pos.Y
	bottom := -1*cam.Zoom + cam.Transform.Pos.Y
	gl.Ortho(float64(left), float64(right), float64(bottom), float64(top), -3, 3)

	// Rotation & position
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(float32(cam.Transform.Pos.X), float32(cam.Transform.Pos.Y), 0)
	gl.Rotatef(float32(cam.Transform.Rot), 0, 0, 1)
	gl.Translatef(-float32(cam.Transform.Pos.X), -float32(cam.Transform.Pos.Y), 0)
}
