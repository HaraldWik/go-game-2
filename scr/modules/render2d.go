package mod

import (
	"math"

	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
	"github.com/go-gl/gl/v2.1/gl"
)

type RenderRect2D struct {
	Transform Transform2D
	Color     vec3.Type
}

func (rect *RenderRect2D) Update() {
	gl.Begin(gl.QUADS)
	gl.Color3f(rect.Color.X, rect.Color.Y, rect.Color.Z)

	// Calculate vertices after rot
	sinR := float32(math.Sin(float64(rect.Transform.Rot / 360)))
	cosR := float32(math.Cos(float64(rect.Transform.Rot / 360)))

	x0 := -0.5*rect.Transform.Size.X*cosR - -0.5*rect.Transform.Size.Y*sinR + rect.Transform.Pos.X
	y0 := -0.5*rect.Transform.Size.X*sinR + -0.5*rect.Transform.Size.Y*cosR + rect.Transform.Pos.Y

	x1 := 0.5*rect.Transform.Size.X*cosR - -0.5*rect.Transform.Size.Y*sinR + rect.Transform.Pos.X
	y1 := 0.5*rect.Transform.Size.X*sinR + -0.5*rect.Transform.Size.Y*cosR + rect.Transform.Pos.Y

	x2 := 0.5*rect.Transform.Size.X*cosR - 0.5*rect.Transform.Size.Y*sinR + rect.Transform.Pos.X
	y2 := 0.5*rect.Transform.Size.X*sinR + 0.5*rect.Transform.Size.Y*cosR + rect.Transform.Pos.Y

	x3 := -0.5*rect.Transform.Size.X*cosR - 0.5*rect.Transform.Size.Y*sinR + rect.Transform.Pos.X
	y3 := -0.5*rect.Transform.Size.X*sinR + 0.5*rect.Transform.Size.Y*cosR + rect.Transform.Pos.Y

	// Draw vertices
	gl.Vertex2f(x0, y0) // Bottom-left
	gl.Vertex2f(x1, y1) // Bottom-right
	gl.Vertex2f(x2, y2) // Top-right
	gl.Vertex2f(x3, y3) // Top-left

	gl.End()
}

type RenderTriangle2D struct {
	Transform Transform2D
	Color     vec3.Type
	Flip      bool
}

func (triangle *RenderTriangle2D) Update() {
	gl.Begin(gl.TRIANGLES)
	gl.Color3f(triangle.Color.X, triangle.Color.Y, triangle.Color.Z)

	// *Calculate vertices after rotation
	sinR := float32(math.Sin(float64(triangle.Transform.Rot / 360)))
	cosR := float32(math.Cos(float64(triangle.Transform.Rot / 360)))

	x0 := triangle.Transform.Pos.X
	y0 := triangle.Transform.Pos.Y + triangle.Transform.Size.Y/2

	var x1, y1, x2, y2 float32

	if triangle.Flip {
		x1 = triangle.Transform.Pos.X - triangle.Transform.Size.X/2*cosR + triangle.Transform.Size.Y/2*sinR
		y1 = triangle.Transform.Pos.Y - triangle.Transform.Size.X/2*sinR - triangle.Transform.Size.Y/2*cosR

		x2 = triangle.Transform.Pos.X + triangle.Transform.Size.X/2*cosR + triangle.Transform.Size.Y/2*sinR
		y2 = triangle.Transform.Pos.Y + triangle.Transform.Size.X/2*sinR - triangle.Transform.Size.Y/2*cosR
	} else {
		x1 = triangle.Transform.Pos.X + triangle.Transform.Size.X/2*cosR - triangle.Transform.Size.Y/2*sinR
		y1 = triangle.Transform.Pos.Y + triangle.Transform.Size.X/2*sinR + triangle.Transform.Size.Y/2*cosR

		x2 = triangle.Transform.Pos.X - triangle.Transform.Size.X/2*cosR - triangle.Transform.Size.Y/2*sinR
		y2 = triangle.Transform.Pos.Y - triangle.Transform.Size.X/2*sinR + triangle.Transform.Size.Y/2*cosR
	}

	// *Draw vertices
	gl.Vertex2f(x0, y0) // *Top
	gl.Vertex2f(x1, y1) // *Bottom-right
	gl.Vertex2f(x2, y2) // *Bottom-left

	gl.End()
}

type RenderCircle2D struct {
	Transform Transform2D
	Color     vec3.Type
	Segments  int32
}

func (circle RenderCircle2D) Update() {
	gl.Begin(gl.LINE_LOOP)
	gl.Color3f(circle.Color.X, circle.Color.Y, circle.Color.Z)

	theta := 2.0 * math.Pi / float64(circle.Segments)

	// *Draw circle points
	for i := 0; i < int(circle.Segments); i++ {
		x := float32(math.Cos(float64(i)*theta)) * circle.Transform.Size.X / 2.0
		y := float32(math.Sin(float64(i)*theta)) * circle.Transform.Size.Y / 2.0

		// *Apply Pos & rotation
		cosR := float32(math.Cos(float64(circle.Transform.Rot)))
		sinR := float32(math.Sin(float64(circle.Transform.Rot)))
		rotatedX := x*cosR - y*sinR + circle.Transform.Pos.X
		rotatedY := x*sinR + y*cosR + circle.Transform.Pos.Y

		gl.Vertex2f(rotatedX, rotatedY)
	}

	gl.End()
}