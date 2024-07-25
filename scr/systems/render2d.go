package sys

import (
	"math"

	dt "github.com/HaraldWik/go-game-2/scr/data-types"
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
	"github.com/go-gl/gl/v2.1/gl"
)

type RenderRect2D struct{}

func (r RenderRect2D) Update() {
	var (
		obj       = ups.Manager.GetParent()
		color     = obj.Data.Get("Color").(vec3.Type)
		transform = obj.Data.Get("Transform").(dt.Transform2D)
	)

	gl.Begin(gl.QUADS)
	gl.Color3f(color.X, color.Y, color.Z)

	// Calculate vertices after rot
	sinR := float32(math.Sin(float64(transform.Rot / 360)))
	cosR := float32(math.Cos(float64(transform.Rot / 360)))

	x0 := -0.5*transform.Size.X*cosR - -0.5*transform.Size.Y*sinR + transform.Pos.X
	y0 := -0.5*transform.Size.X*sinR + -0.5*transform.Size.Y*cosR + transform.Pos.Y

	x1 := 0.5*transform.Size.X*cosR - -0.5*transform.Size.Y*sinR + transform.Pos.X
	y1 := 0.5*transform.Size.X*sinR + -0.5*transform.Size.Y*cosR + transform.Pos.Y

	x2 := 0.5*transform.Size.X*cosR - 0.5*transform.Size.Y*sinR + transform.Pos.X
	y2 := 0.5*transform.Size.X*sinR + 0.5*transform.Size.Y*cosR + transform.Pos.Y

	x3 := -0.5*transform.Size.X*cosR - 0.5*transform.Size.Y*sinR + transform.Pos.X
	y3 := -0.5*transform.Size.X*sinR + 0.5*transform.Size.Y*cosR + transform.Pos.Y

	// Draw vertices
	gl.Vertex2f(x0, y0) // Bottom-left
	gl.Vertex2f(x1, y1) // Bottom-right
	gl.Vertex2f(x2, y2) // Top-right
	gl.Vertex2f(x3, y3) // Top-left

	gl.End()
}

type RenderTriangle2D struct{}

func (t RenderTriangle2D) Update() {
	var (
		obj       = ups.Manager.GetParent()
		color     = obj.Data.Get("Color").(vec3.Type)
		transform = obj.Data.Get("Transform").(dt.Transform2D)
		flip      = obj.Data.Get("Flip").(bool)
	)

	gl.Begin(gl.TRIANGLES)
	gl.Color3f(color.X, color.Y, color.Z)

	// Calculate vertices after rotation
	sinR := float32(math.Sin(float64(transform.Rot / 360)))
	cosR := float32(math.Cos(float64(transform.Rot / 360)))

	x0 := transform.Pos.X
	y0 := transform.Pos.Y + transform.Size.Y/2

	var x1, y1, x2, y2 float32

	if flip {
		x1 = transform.Pos.X - transform.Size.X/2*cosR + transform.Size.Y/2*sinR
		y1 = transform.Pos.Y - transform.Size.X/2*sinR - transform.Size.Y/2*cosR

		x2 = transform.Pos.X + transform.Size.X/2*cosR + transform.Size.Y/2*sinR
		y2 = transform.Pos.Y + transform.Size.X/2*sinR - transform.Size.Y/2*cosR
	} else {
		x1 = transform.Pos.X + transform.Size.X/2*cosR - transform.Size.Y/2*sinR
		y1 = transform.Pos.Y + transform.Size.X/2*sinR + transform.Size.Y/2*cosR

		x2 = transform.Pos.X - transform.Size.X/2*cosR - transform.Size.Y/2*sinR
		y2 = transform.Pos.Y - transform.Size.X/2*sinR + transform.Size.Y/2*cosR
	}

	// Draw vertices
	gl.Vertex2f(x0, y0) // *Top
	gl.Vertex2f(x1, y1) // *Bottom-right
	gl.Vertex2f(x2, y2) // *Bottom-left

	gl.End()
}

type RenderCircle2D struct{}

func (c RenderCircle2D) Update() {
	var (
		obj       = ups.Manager.GetParent()
		color     = obj.Data.Get("Color").(vec3.Type)
		transform = obj.Data.Get("Transform").(dt.Transform2D)
		segments  = obj.Data.Get("Segments").(uint32)
	)

	gl.Begin(gl.LINE_LOOP)
	gl.Color3f(color.X, color.Y, color.Z)

	theta := 2.0 * math.Pi / float64(segments)

	// Draw circle points
	for i := 0; i < int(segments); i++ {
		x := float32(math.Cos(float64(i)*theta)) * transform.Size.X / 2.0
		y := float32(math.Sin(float64(i)*theta)) * transform.Size.Y / 2.0

		// Apply Pos & rotation
		cosR := float32(math.Cos(float64(transform.Rot)))
		sinR := float32(math.Sin(float64(transform.Rot)))
		rotatedX := x*cosR - y*sinR + transform.Pos.X
		rotatedY := x*sinR + y*cosR + transform.Pos.Y

		gl.Vertex2f(rotatedX, rotatedY)
	}

	gl.End()
}
