package sys

import (
	"github.com/HaraldWik/go-game-2/scr/app"
	dt "github.com/HaraldWik/go-game-2/scr/data-types"
	"github.com/HaraldWik/go-game-2/scr/ups"
	"github.com/go-gl/gl/v2.1/gl"
)

type Camera2D struct{}

func (c Camera2D) Update() {
	var (
		obj       = ups.Manager.GetParent()
		window    = obj.Data.Get("Window").(app.Win)
		transform = obj.Data.Get("Transform").(dt.Transform2D)
		zoom      = obj.Data.Get("Zoom").(float32)
	)

	gl.Viewport(0, 0, int32(window.GetSize().X), int32(window.GetSize().Y))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	// Calculate aspect ratio
	aspect := window.GetSize().X / window.GetSize().Y
	left := -1*zoom*aspect + transform.Pos.X
	right := 1*zoom*aspect + transform.Pos.X
	top := 1*zoom + transform.Pos.Y
	bottom := -1*zoom + transform.Pos.Y
	gl.Ortho(float64(left), float64(right), float64(bottom), float64(top), -3, 3)

	// Rotation & position
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(float32(transform.Pos.X), float32(transform.Pos.Y), 0)
	gl.Rotatef(float32(transform.Rot), 0, 0, 1)
	gl.Translatef(-float32(transform.Pos.X), -float32(transform.Pos.Y), 0)
}
