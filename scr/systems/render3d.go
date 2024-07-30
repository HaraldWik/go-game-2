package sys

import (
	dt "github.com/HaraldWik/go-game-2/scr/data-types"
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
	"github.com/go-gl/gl/v2.1/gl"
)

type RenderCube3D struct{}

func (c RenderCube3D) Start() {}

func (c RenderCube3D) Update(deltaTime float32) {
	var (
		obj       = ups.Engine.GetParent()
		color     = obj.Data.Get("Color").(vec3.Type)
		transform = obj.Data.Get("Transform").(dt.Transform3D)
	)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(transform.Pos.X, transform.Pos.Y, transform.Pos.Z)
	gl.Rotatef(0.0, transform.Rot.X/360, transform.Rot.Y/360, transform.Rot.Z/360)

	gl.Begin(gl.QUADS)

	gl.Color3f(color.X, color.Y, color.Z)

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
