package mod

import (
	"math"

	"github.com/HaraldWik/go-game-2/scr/abus"
	"github.com/HaraldWik/go-game-2/scr/app"
	"github.com/go-gl/gl/v2.1/gl"
)

type Cam3D struct {
	Win        app.Win
	transform  Transform3D
	Fov        float32
	hasStarted bool
}

func (obj *Cam3D) Update() {
	if !obj.hasStarted {
		curObj := abus.SceneManager.GetCurrentScene().CurObj
		curObj.AddProperty(obj)
		obj.hasStarted = true
	}

	gl.Viewport(0, 0, int32(obj.Win.GetSize().X), int32(obj.Win.GetSize().Y))

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	gluPerspective(float64(obj.Fov), float64(obj.Win.GetSize().X)/float64(obj.Win.GetSize().Y), 0.1, 100.0)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Rotatef(obj.transform.Rot.X, 1, 0, 0)
	gl.Rotatef(obj.transform.Rot.Y, 0, 1, 0)
	gl.Rotatef(obj.transform.Rot.Z, 0, 0, 1)
	gl.Translatef(-obj.transform.Rot.X, -obj.transform.Rot.Y, -obj.transform.Rot.Z)
}

func gluPerspective(fovY, aspect, zNear, zFar float64) {
	fH := math.Tan(fovY/360*math.Pi) * zNear
	fW := fH * aspect
	gl.Frustum(-fW, fW, -fH, fH, zNear, zFar)
}
