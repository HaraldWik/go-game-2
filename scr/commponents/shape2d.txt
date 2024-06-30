package object

import (
	"math"

	component "github.com/HaraldWik/go-game/scr/components"
	"github.com/go-gl/gl/v2.1/gl"
)

type Shape2D struct {
	Transform component.Transform2D
	Color     component.RGB
}

func NewShape2DAdvanced(transform component.Transform2D, rgb component.RGB) Shape2D {
	return Shape2D{
		Transform: transform,
		Color:     rgb,
	}
}

func (shape Shape2D) DrawRectangle() {
	gl.Begin(gl.QUADS)
	gl.Color3f(shape.Color.R, shape.Color.G, shape.Color.B)

	// *Calculate vertices after rot
	sinR := float32(math.Sin(float64(shape.Transform.Rotation)))
	cosR := float32(math.Cos(float64(shape.Transform.Rotation)))

	x0 := -0.5*shape.Transform.Size.X*cosR - -0.5*shape.Transform.Size.Y*sinR + shape.Transform.Position.X
	y0 := -0.5*shape.Transform.Size.X*sinR + -0.5*shape.Transform.Size.Y*cosR + shape.Transform.Position.Y

	x1 := 0.5*shape.Transform.Size.X*cosR - -0.5*shape.Transform.Size.Y*sinR + shape.Transform.Position.X
	y1 := 0.5*shape.Transform.Size.X*sinR + -0.5*shape.Transform.Size.Y*cosR + shape.Transform.Position.Y

	x2 := 0.5*shape.Transform.Size.X*cosR - 0.5*shape.Transform.Size.Y*sinR + shape.Transform.Position.X
	y2 := 0.5*shape.Transform.Size.X*sinR + 0.5*shape.Transform.Size.Y*cosR + shape.Transform.Position.Y

	x3 := -0.5*shape.Transform.Size.X*cosR - 0.5*shape.Transform.Size.Y*sinR + shape.Transform.Position.X
	y3 := -0.5*shape.Transform.Size.X*sinR + 0.5*shape.Transform.Size.Y*cosR + shape.Transform.Position.Y

	// *Draw vertices
	gl.Vertex2f(x0, y0) // *Bottom-left
	gl.Vertex2f(x1, y1) // *Bottom-right
	gl.Vertex2f(x2, y2) // *Top-right
	gl.Vertex2f(x3, y3) // *Top-left

	gl.End()
}

func (shape Shape2D) DrawTriangle(flip bool) {
	gl.Begin(gl.TRIANGLES)
	gl.Color3f(shape.Color.R, shape.Color.G, shape.Color.B)

	// *Calculate vertices after rotation
	sinR := float32(math.Sin(float64(shape.Transform.Rotation)))
	cosR := float32(math.Cos(float64(shape.Transform.Rotation)))

	x0 := shape.Transform.Position.X
	y0 := shape.Transform.Position.Y + shape.Transform.Size.Y/2

	var x1, y1, x2, y2 float32

	if flip {
		x1 = shape.Transform.Position.X - shape.Transform.Size.X/2*cosR + shape.Transform.Size.Y/2*sinR
		y1 = shape.Transform.Position.Y - shape.Transform.Size.X/2*sinR - shape.Transform.Size.Y/2*cosR

		x2 = shape.Transform.Position.X + shape.Transform.Size.X/2*cosR + shape.Transform.Size.Y/2*sinR
		y2 = shape.Transform.Position.Y + shape.Transform.Size.X/2*sinR - shape.Transform.Size.Y/2*cosR
	} else {
		x1 = shape.Transform.Position.X + shape.Transform.Size.X/2*cosR - shape.Transform.Size.Y/2*sinR
		y1 = shape.Transform.Position.Y + shape.Transform.Size.X/2*sinR + shape.Transform.Size.Y/2*cosR

		x2 = shape.Transform.Position.X - shape.Transform.Size.X/2*cosR - shape.Transform.Size.Y/2*sinR
		y2 = shape.Transform.Position.Y - shape.Transform.Size.X/2*sinR + shape.Transform.Size.Y/2*cosR
	}

	// *Draw vertices
	gl.Vertex2f(x0, y0) // *Top
	gl.Vertex2f(x1, y1) // *Bottom-right
	gl.Vertex2f(x2, y2) // *Bottom-left

	gl.End()
}

func (shape Shape2D) DrawCircle(segments int) {
	gl.Begin(gl.LINE_LOOP)
	gl.Color3f(shape.Color.R, shape.Color.G, shape.Color.B)

	theta := 2.0 * math.Pi / float64(segments)

	// *Draw circle points
	for i := 0; i < segments; i++ {
		x := float32(math.Cos(float64(i)*theta)) * shape.Transform.Size.X / 2.0
		y := float32(math.Sin(float64(i)*theta)) * shape.Transform.Size.Y / 2.0

		// *Apply position & rotation
		cosR := float32(math.Cos(float64(shape.Transform.Rotation)))
		sinR := float32(math.Sin(float64(shape.Transform.Rotation)))
		rotatedX := x*cosR - y*sinR + shape.Transform.Position.X
		rotatedY := x*sinR + y*cosR + shape.Transform.Position.Y

		gl.Vertex2f(rotatedX, rotatedY)
	}

	gl.End()
}
