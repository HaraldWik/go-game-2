package s3d

import (
	"math"

	d3d "github.com/HaraldWik/go-game-2/scr/3d/data"
	"github.com/HaraldWik/go-game-2/scr/app"
	"github.com/HaraldWik/go-game-2/scr/ups"
	"github.com/go-gl/gl/v2.1/gl"
)

type Camera3D struct{}

func (c Camera3D) Update(obj *ups.Object, deltaTime float32) {
	var (
		window    = obj.Data.Get("Window").(app.Window)
		transform = obj.Data.Get("Transform").(d3d.Transform3D)
		fov       = obj.Data.Get("Fov").(float32)
	)

	near := 0.1
	far := 100.0

	// Perspective projection matrix
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	f := 1.0 / math.Tan(float64(fov)/2.0*math.Pi/180.0)
	gl.Frustum(
		-float64(window.GetSize().X)*f, float64(window.GetSize().Y)*f,
		-f, f,
		near, far,
	)

	// View matrix
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	// Apply camera rotations (pitch, yaw, roll)
	gl.Rotatef(float32(-transform.Rotation.X), 1.0, 0.0, 0.0)
	gl.Rotatef(float32(-transform.Rotation.Y), 0.0, 1.0, 0.0)
	gl.Rotatef(float32(-transform.Rotation.Z), 0.0, 0.0, 1.0)

	// Apply camera translation
	gl.Translatef(float32(-transform.Position.X), float32(-transform.Position.Y), float32(-transform.Position.Z))
}
