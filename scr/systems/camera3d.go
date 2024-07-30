package sys

import (
	"math"

	"github.com/HaraldWik/go-game-2/scr/app"
	dt "github.com/HaraldWik/go-game-2/scr/data-types"
	"github.com/HaraldWik/go-game-2/scr/ups"
	"github.com/go-gl/gl/v2.1/gl"
)

type Camera3D struct{} /*
# Data:

| ¤ 'Window' <app.Window>
| ¤ 'Transform' <dt.Transform2D>
| ¤ 'Fov' <float32>

# Description:

<Camera3D> is a 3 dimentinal OpenGL camera that supports position, rotation and fov.

# Exampel:

```
	ups.NewObjectOptimal(
		"Camera",
		ups.Data{
			"Window":    myWindow,
			"Transform": dt.NewTransform2D(vec2.Zero(), vec2.All(1.0), 0.0),
			"Fov":      float32(45.0),
		},
		sys.Camera3D{},
	)
```

# Info:

'app', 'dt', 'ups' & 'sys' are packages.

Declaration of the data 'Fov' is needed with 'float32()' or a variable of type <float32>

*/

func (c Camera3D) Start() {}

func (c Camera3D) Update(deltaTime float32) {
	var (
		obj       = ups.Engine.GetParent()
		window    = obj.Data.Get("Window").(app.Window)
		transform = obj.Data.Get("Transform").(dt.Transform3D)
		fov       = obj.Data.Get("Fov").(float32)
	)

	gl.Viewport(0, 0, int32(window.GetSize().X), int32(window.GetSize().Y))

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	gluPerspective(float64(fov), float64(window.GetSize().X)/float64(window.GetSize().Y), 0.1, 100.0)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Rotatef(transform.Rot.X, 1, 0, 0)
	gl.Rotatef(transform.Rot.Y, 0, 1, 0)
	gl.Rotatef(transform.Rot.Z, 0, 0, 1)
	gl.Translatef(-transform.Rot.X, -transform.Rot.Y, -transform.Rot.Z)
}

func gluPerspective(fovY, aspect, zNear, zFar float64) {
	fH := math.Tan(fovY/360*math.Pi) * zNear
	fW := fH * aspect
	gl.Frustum(-fW, fW, -fH, fH, zNear, zFar)
}
