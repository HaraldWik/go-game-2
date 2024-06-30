package object

import (
	component "github.com/HaraldWik/go-game/scr/components"
	win "github.com/HaraldWik/go-game/scr/window"
	"github.com/go-gl/gl/v2.1/gl"
)

type Camera2D struct {
	window      win.Window
	Transform2D component.Transform2D
	zoom        float32
}

func NewCamera2D() Camera2D {
	return Camera2D{}
}

func NewCamera2DAdvanced(window win.Window, transform component.Transform2D, zoom float32) Camera2D {
	return Camera2D{
		window:      window,
		Transform2D: transform,
		zoom:        zoom,
	}
}

// *Draw
func (camera Camera2D) Draw() {
	gl.Viewport(0, 0, int32(camera.window.GetSizeWidth()), int32(camera.window.GetSizeHeight()))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	// *Calculate aspect ratio
	aspect := float32(camera.window.GetSizeWidth()) / float32(camera.window.GetSizeHeight())
	left := -1*camera.zoom*aspect + camera.Transform2D.Position.X
	right := 1*camera.zoom*aspect + camera.Transform2D.Position.X
	top := 1*camera.zoom + camera.Transform2D.Position.Y
	bottom := -1*camera.zoom + camera.Transform2D.Position.Y
	gl.Ortho(float64(left), float64(right), float64(bottom), float64(top), -3, 3)

	// *Rotation & position
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(float32(camera.Transform2D.Position.X), float32(camera.Transform2D.Position.Y), 0)
	gl.Rotatef(float32(camera.Transform2D.Rotation), 0, 0, 1)
	gl.Translatef(-float32(camera.Transform2D.Position.X), -float32(camera.Transform2D.Position.Y), 0)
}
