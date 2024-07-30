package sys

import (
	dt "github.com/HaraldWik/go-game-2/scr/data-types"
	"github.com/HaraldWik/go-game-2/scr/ups"
)

type Physics2D struct{} /*
TEMPPORARY

# Data:

| ¤ 'Transform' <dt.Transform2D>

# Description:

<Physics2D> is a 3 dimentinal OpenGL camera that supports position, rotation and fov.

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

func (p Physics2D) Start() {}

func (p Physics2D) Update(deltaTime float32) {
	var (
		obj       = ups.Engine.GetParent()
		transform = obj.Data.Get("Transform").(dt.Transform2D)
	)

	//Collider

	if transform.Pos.Y-transform.Size.Y > 0.0 {
		transform.Pos.Y -= 9.8 * deltaTime
	} else {
		transform.Pos.Y = transform.Size.Y
	}
	obj.Data.Set("Transform", transform)
}
