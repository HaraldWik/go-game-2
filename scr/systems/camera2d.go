package sys

import (
	"github.com/HaraldWik/go-game-2/scr/app"
	dt "github.com/HaraldWik/go-game-2/scr/data-types"
	"github.com/HaraldWik/go-game-2/scr/ups"
	"github.com/go-gl/gl/v2.1/gl"
)

type Camera2D struct{} /*
# Data:

| ¤ 'Window' <app.Window>
| ¤ 'Transform' <dt.Transform2D>
| ¤ 'Zoom' <float32>

# Description:

<Camera2D> is a 2 dimentinal OpenGL camera that supports position, rotation and zoom.

# Exampel:

```
	ups.NewObjectOptimal(
		"Camera",
		ups.Data{
			"Window":    myWindow,
			"Transform": dt.NewTransform2D(vec2.Zero(), vec2.All(1.0), 0.0),
			"Zoom":      float32(10.0),
		},
		sys.Camera2D{},
	)
```

# Info:

'app', 'dt', 'ups' & 'sys' are packages.

Declaration of the data 'Zoom' is needed with 'float32()' or a variable of type <float32>

*/

func (c Camera2D) Start() {}

func (c Camera2D) Update(deltaTime float32) {
	var (
		obj       = ups.Engine.GetParent()
		window    = obj.Data.Get("Window").(app.Window)
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
