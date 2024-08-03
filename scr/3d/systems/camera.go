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

	gl.Viewport(0, 0, int32(window.GetSize().X), int32(window.GetSize().Y))

	// Set up the projection matrix
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	// Compute perspective projection matrix
	aspectRatio := float64(window.GetSize().X) / float64(window.GetSize().Y)
	gluPerspective(float64(fov), aspectRatio, 0.1, 100.0)

	// Set up the modelview matrix
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	// Apply transformations
	gl.Translatef(transform.Pos.X, transform.Pos.Y, transform.Pos.Z)
	gl.Rotatef(transform.Rot.X, 1, 0, 0) // Rotations are typically around axes, so adjust accordingly
	gl.Rotatef(transform.Rot.Y, 0, 1, 0)
	gl.Rotatef(transform.Rot.Z, 0, 0, 1)

	gl.PopMatrix()
}

func gluPerspective(fovY, aspect, zNear, zFar float64) {
	fH := math.Tan(fovY*math.Pi/360) * zNear
	fW := fH * aspect
	gl.Frustum(-fW, fW, -fH, fH, zNear, zFar)
}
