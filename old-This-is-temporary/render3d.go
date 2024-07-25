package mod

import (
	"github.com/HaraldWik/go-game-2/scr/abus"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
	"github.com/go-gl/gl/v2.1/gl"
)

type RenderCube3D struct {
	transform  Transform3D
	Color      vec3.Type
	Angle      float32
	hasStarted bool
}

func (obj *RenderCube3D) Update() {
	if !obj.hasStarted {
		curObj := abus.SceneManager.GetCurrentScene().CurObj
		curObj.AddProperty(obj)
		obj.hasStarted = true
	}

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(obj.transform.Pos.X, obj.transform.Pos.Y, obj.transform.Pos.Z)
	gl.Rotatef(obj.Angle, obj.transform.Rot.X/360, obj.transform.Rot.Y/360, obj.transform.Rot.Z/360)

	gl.Begin(gl.QUADS)

	gl.Color3f(obj.Color.X, obj.Color.Y, obj.Color.Z)

	// Front face
	gl.Vertex3f(-1.0, -1.0, 1.0)
	gl.Vertex3f(1.0, -1.0, 1.0)
	gl.Vertex3f(1.0, 1.0, 1.0)
	gl.Vertex3f(-1.0, 1.0, 1.0)

	// Back face
	gl.Vertex3f(-1.0, -1.0, -1.0)
	gl.Vertex3f(-1.0, 1.0, -1.0)
	gl.Vertex3f(1.0, 1.0, -1.0)
	gl.Vertex3f(1.0, -1.0, -1.0)

	// Top face
	gl.Vertex3f(-1.0, 1.0, -1.0)
	gl.Vertex3f(-1.0, 1.0, 1.0)
	gl.Vertex3f(1.0, 1.0, 1.0)
	gl.Vertex3f(1.0, 1.0, -1.0)

	// Bottom face
	gl.Vertex3f(-1.0, -1.0, -1.0)
	gl.Vertex3f(1.0, -1.0, -1.0)
	gl.Vertex3f(1.0, -1.0, 1.0)
	gl.Vertex3f(-1.0, -1.0, 1.0)

	// Right face
	gl.Vertex3f(1.0, -1.0, -1.0)
	gl.Vertex3f(1.0, 1.0, -1.0)
	gl.Vertex3f(1.0, 1.0, 1.0)
	gl.Vertex3f(1.0, -1.0, 1.0)

	// Left face
	gl.Vertex3f(-1.0, -1.0, -1.0)
	gl.Vertex3f(-1.0, -1.0, 1.0)
	gl.Vertex3f(-1.0, 1.0, 1.0)
	gl.Vertex3f(-1.0, 1.0, -1.0)

	gl.End()
}
