package component

import (
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
	"github.com/go-gl/gl/v2.1/gl"
)

type RenderCube3D struct {
	Transform Transform3D
	Color     vec3.Type
}

func (cube *RenderCube3D) Update(angle float32) {
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(0, 0, -10) // Move the cube further back
	gl.Rotatef(angle, cube.Transform.Rot.X/360, cube.Transform.Rot.Y/360, cube.Transform.Rot.Z/360)

	gl.Begin(gl.QUADS)

	// Front face
	gl.Color3f(1.0, 0.0, 0.0)
	gl.Vertex3f(-1.0, -1.0, 1.0)
	gl.Vertex3f(1.0, -1.0, 1.0)
	gl.Vertex3f(1.0, 1.0, 1.0)
	gl.Vertex3f(-1.0, 1.0, 1.0)

	// Back face
	gl.Color3f(0.0, 1.0, 0.0)
	gl.Vertex3f(-1.0, -1.0, -1.0)
	gl.Vertex3f(-1.0, 1.0, -1.0)
	gl.Vertex3f(1.0, 1.0, -1.0)
	gl.Vertex3f(1.0, -1.0, -1.0)

	// Top face
	gl.Color3f(0.0, 0.0, 1.0)
	gl.Vertex3f(-1.0, 1.0, -1.0)
	gl.Vertex3f(-1.0, 1.0, 1.0)
	gl.Vertex3f(1.0, 1.0, 1.0)
	gl.Vertex3f(1.0, 1.0, -1.0)

	// Bottom face
	gl.Color3f(1.0, 1.0, 0.0)
	gl.Vertex3f(-1.0, -1.0, -1.0)
	gl.Vertex3f(1.0, -1.0, -1.0)
	gl.Vertex3f(1.0, -1.0, 1.0)
	gl.Vertex3f(-1.0, -1.0, 1.0)

	// Right face
	gl.Color3f(1.0, 0.0, 1.0)
	gl.Vertex3f(1.0, -1.0, -1.0)
	gl.Vertex3f(1.0, 1.0, -1.0)
	gl.Vertex3f(1.0, 1.0, 1.0)
	gl.Vertex3f(1.0, -1.0, 1.0)

	// Left face
	gl.Color3f(0.0, 1.0, 1.0)
	gl.Vertex3f(-1.0, -1.0, -1.0)
	gl.Vertex3f(-1.0, -1.0, 1.0)
	gl.Vertex3f(-1.0, 1.0, 1.0)
	gl.Vertex3f(-1.0, 1.0, -1.0)

	gl.End()
}
