package sys

import (
	"log"
	"math"

	dt "github.com/HaraldWik/go-game-2/scr/data-types"
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
	"github.com/go-gl/gl/v2.1/gl"
)

type RenderMesh2D struct{}

func (t RenderMesh2D) Start() {}

func (t RenderMesh2D) Update(deltaTime float32) {
	var (
		obj      = ups.Engine.GetParent()
		material = obj.Data.Get("Material").(dt.Material)
		vertices = obj.Data.Get("Vertices").([]vec2.Type)
	)

	if len(vertices) < 3 {
		log.Fatalf("RenderMesh2D requires at least 3 vertices")
	}

	gl.Color3f(material.Alpha.X, material.Alpha.Y, material.Alpha.Z)
	gl.Enable(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, material.Texture.Image)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	// Calculate the bounding box of the vertices
	minX, minY := vertices[0].X, vertices[0].Y
	maxX, maxY := vertices[0].X, vertices[0].Y

	for _, v := range vertices {
		if v.X < minX {
			minX = v.X
		}
		if v.X > maxX {
			maxX = v.X
		}
		if v.Y < minY {
			minY = v.Y
		}
		if v.Y > maxY {
			maxY = v.Y
		}
	}

	width := maxX - minX
	height := maxY - minY

	// Render the mesh
	gl.Begin(gl.TRIANGLE_FAN) // Use TRIANGLE_FAN to handle arbitrary convex polygons
	for _, vertex := range vertices {
		// Map texture coordinates based on the bounding box
		tx := (vertex.X - minX) / width
		ty := (vertex.Y - minY) / height
		gl.TexCoord2f(tx, ty)
		gl.Vertex2f(vertex.X, vertex.Y)
	}
	gl.End()

	// Disable texturing and blending after rendering
	gl.Disable(gl.TEXTURE_2D)
	gl.Disable(gl.BLEND)

	// Check for OpenGL errors
	if err := gl.GetError(); err != gl.NO_ERROR {
		log.Fatalf("OpenGL Error: %v\n", err)
	}
}

type RenderTexture2D struct{}

func (t RenderTexture2D) Start() {}

func (t RenderTexture2D) Update(deltaTime float32) {
	var (
		obj       = ups.Engine.GetParent()
		material  = obj.Data.Get("Material").(dt.Material)
		transform = obj.Data.Get("Transform").(dt.Transform2D)
	)

	gl.MatrixMode(gl.MODELVIEW)
	gl.PushMatrix()

	gl.LoadIdentity()
	gl.Translatef(float32(transform.Pos.X), float32(transform.Pos.Y), 0)
	gl.Rotatef(float32(-transform.Rot), 0, 0, 1)

	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.BindTexture(gl.TEXTURE_2D, material.Texture.Image)

	gl.Color3f(material.Alpha.X, material.Alpha.Y, material.Alpha.Z)

	gl.Begin(gl.QUADS)
	gl.TexCoord2f(0, 1)
	gl.Vertex2f(-float32(transform.Size.X)/2, -float32(transform.Size.Y)/2)
	gl.TexCoord2f(1, 1)
	gl.Vertex2f(float32(transform.Size.X)/2, -float32(transform.Size.Y)/2)
	gl.TexCoord2f(1, 0)
	gl.Vertex2f(float32(transform.Size.X)/2, float32(transform.Size.Y)/2)
	gl.TexCoord2f(0, 0)
	gl.Vertex2f(-float32(transform.Size.X)/2, float32(transform.Size.Y)/2)
	gl.End()

	gl.Disable(gl.TEXTURE_2D)
	gl.Disable(gl.BLEND)

	gl.PopMatrix()

	if err := gl.GetError(); err != gl.NO_ERROR {
		log.Fatalf("OpenGL Error: %v\n", err)
	}
}

type RenderRectangle2D struct{} /*
# Data:

| ¤ 'Color' <vec3.Type>
| ¤ 'Transform' <dt.Transform2D>

# Description:

<RenderRectangle2D> is a 2 dimentinal OpenGL rectangle that supports position, rotation and size.

# Exampel:

```
	ups.NewObjectOptimal(
		"Rectangle",
		ups.Data{
			"Color":    vec3.New(1.0, 0.0, 0.0),
			"Transform": dt.NewTransform2D(vec2.Zero(), vec2.All(1.0), 0.0),
		},
		sys.RenderRectangle2D{},
	)
```

# Info:

'dt', 'ups' & 'sys' are packages.
*/

func (r RenderRectangle2D) Start() {}

func (r RenderRectangle2D) Update(deltaTime float32) {
	var (
		obj       = ups.Engine.GetParent()
		color     = obj.Data.Get("Color").(vec3.Type)
		transform = obj.Data.Get("Transform").(dt.Transform2D)
	)

	gl.MatrixMode(gl.MODELVIEW)
	gl.PushMatrix()

	gl.LoadIdentity()
	gl.Translatef(float32(transform.Pos.X), float32(transform.Pos.Y), 0)
	gl.Rotatef(float32(-transform.Rot), 0, 0, 1)

	gl.Color3f(color.X, color.Y, color.Z)

	gl.Begin(gl.QUADS)
	gl.Vertex2f(-float32(transform.Size.X)/2, -float32(transform.Size.Y)/2)
	gl.Vertex2f(float32(transform.Size.X)/2, -float32(transform.Size.Y)/2)
	gl.Vertex2f(float32(transform.Size.X)/2, float32(transform.Size.Y)/2)
	gl.Vertex2f(-float32(transform.Size.X)/2, float32(transform.Size.Y)/2)
	gl.End()

	gl.Disable(gl.TEXTURE_2D)

	gl.PopMatrix()

	if err := gl.GetError(); err != gl.NO_ERROR {
		log.Fatalf("OpenGL Error: %v\n", err)
	}
}

type RenderTriangle2D struct{} /*
# Data:

| ¤ 'Color' <vec3.Type>
| ¤ 'Transform' <dt.Transform2D>

# Description:

<RenderTriangle2D> is a 2 dimentinal OpenGL triangle that supports position, rotation and size.

# Exampel:

```
	ups.NewObjectOptimal(
		"Triangle",
		ups.Data{
			"Color":    vec3.New(0.0, 1.0, 0.0),
			"Transform": dt.NewTransform2D(vec2.Zero(), vec2.All(1.0), 0.0),
		},
		sys.RenderTriangle2D{},
	)
```

# Info:

'dt', 'ups' & 'sys' are packages.
*/

func (t RenderTriangle2D) Start() {}

func (t RenderTriangle2D) Update(deltaTime float32) {
	var (
		obj       = ups.Engine.GetParent()
		color     = obj.Data.Get("Color").(vec3.Type)
		transform = obj.Data.Get("Transform").(dt.Transform2D)
	)

	gl.Begin(gl.TRIANGLES)
	gl.Color3f(color.X, color.Y, color.Z)

	x0 := transform.Pos.X
	y0 := transform.Pos.Y + transform.Size.Y/2

	var x1, y1, x2, y2 float32

	x1 = transform.Pos.X - transform.Size.X/2 + transform.Size.Y/2
	y1 = transform.Pos.Y - transform.Size.X/2 - transform.Size.Y/2

	x2 = transform.Pos.X + transform.Size.X/2 + transform.Size.Y/2
	y2 = transform.Pos.Y + transform.Size.X/2 - transform.Size.Y/2

	// Draw vertices
	gl.Vertex2f(x0, y0) // *Top
	gl.Vertex2f(x1, y1) // *Bottom-right
	gl.Vertex2f(x2, y2) // *Bottom-left

	gl.End()
}

type RenderCircle2D struct{} /*
# Data:

| ¤ 'Color' <vec3.Type>
| ¤ 'Transform' <dt.Transform2D>
| ¤ 'Segments' <uint32>

# Description:

<RenderCircle2D> is a 2 dimentinal OpenGL circle that supports position, rotation, size and segments.

# Exampel:

```
	ups.NewObjectOptimal(
		"Circle",
		ups.Data{
			"Color":    vec3.New(1.0, 0.0, 0.0),
			"Transform": dt.NewTransform2D(vec2.Zero(), vec2.All(1.0), 0.0),
			"Segments": uint32(16),
		},
		sys.RenderCircle2D{},
	)
```

# Info:

'dt', 'ups' & 'sys' are packages.
*/

func (c RenderCircle2D) Start() {}

func (c RenderCircle2D) Update(deltaTime float32) {
	var (
		obj       = ups.Engine.GetParent()
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
		rotatedX := x - y + transform.Pos.X
		rotatedY := x + y + transform.Pos.Y

		gl.Vertex2f(rotatedX, rotatedY)
	}

	gl.End()
}
