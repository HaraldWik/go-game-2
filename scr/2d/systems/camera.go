package s2d

import (
	d2d "github.com/HaraldWik/go-game-2/scr/2d/data"
	"github.com/HaraldWik/go-game-2/scr/app"

	"github.com/HaraldWik/go-game-2/scr/ups"
	"github.com/go-gl/gl/v2.1/gl"
)

type Camera2D struct{}

func (c Camera2D) Update(obj *ups.Object, deltaTime float32) {
	var (
		window    = obj.Data.Get("Window").(app.Window)
		transform = obj.Data.Get("Transform").(d2d.Transform2D)
		zoom      = obj.Data.Get("Zoom").(float32)
	)

	gl.Viewport(0, 0, int32(window.GetSize().X), int32(window.GetSize().Y))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	aspect := window.GetSize().X / window.GetSize().Y
	left := -zoom*aspect + transform.Position.X
	right := zoom*aspect + transform.Position.X
	top := zoom + transform.Position.Y
	bottom := -zoom + transform.Position.Y
	gl.Ortho(float64(left), float64(right), float64(bottom), float64(top), -10, 10)
}
