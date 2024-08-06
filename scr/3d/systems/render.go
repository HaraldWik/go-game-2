package s3d

import (
	"log"

	d3d "github.com/HaraldWik/go-game-2/scr/3d/data"
	load "github.com/HaraldWik/go-game-2/scr/loaders"
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
	"github.com/go-gl/gl/v2.1/gl"
)

type RenderCube3D struct{}

func (c RenderCube3D) Update(obj *ups.Object, deltaTime float32) {
	var (
		color     = obj.Data.Get("Color").(vec3.Type)
		transform = obj.Data.Get("Transform").(d3d.Transform3D)
	)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(transform.Position.X, transform.Position.Y, transform.Position.Z)
	gl.Rotatef(0.0, transform.Rotation.X, transform.Rotation.Y, transform.Rotation.Z)

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

type RenderObj3D struct{}

func (o RenderObj3D) Start(obj *ups.Object) {}

func (o RenderObj3D) Update(obj *ups.Object, deltaTime float32) {
	var (
		material  = obj.Data.Get("Material").(d3d.Material3D)
		model     = obj.Data.Get("Model").(load.Obj)
		transform = obj.Data.Get("Transform").(d3d.Transform3D)
	)

	if len(model.Vertices) == 0 || len(model.UVs) == 0 || len(model.Indices) == 0 {
		log.Fatalf("Model data is incomplete")
	}

	gl.MatrixMode(gl.MODELVIEW)
	gl.PushMatrix()

	// Apply transformation
	gl.LoadIdentity()
	gl.Translatef(transform.Position.X, transform.Position.Y, transform.Position.Z)
	gl.Rotatef(-transform.Rotation.X, 1, 0, 0)
	gl.Rotatef(-transform.Rotation.Y, 0, 1, 0)
	gl.Rotatef(-transform.Rotation.Z, 0, 0, 1)

	// Bind the texture
	gl.Enable(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, material.Texture.Image)

	// Set color with alpha
	gl.Color3f(material.Alpha.X, material.Alpha.Y, material.Alpha.Z)

	// Enable blending for textures
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	// Render model
	gl.Begin(gl.TRIANGLES)

	// Iterate over indices and draw the model
	for i := 0; i < len(model.Indices); i += 3 {
		if i+2 >= len(model.Indices) {
			log.Fatalf("Index out of range: %d", i)
		}

		// Fetch indices
		i0, i1, i2 := model.Indices[i], model.Indices[i+1], model.Indices[i+2]

		// Validate indices
		if int(i0) >= len(model.Vertices) || int(i1) >= len(model.Vertices) || int(i2) >= len(model.Vertices) {
			log.Fatalf("Vertex index out of bounds: %d, %d, %d", i0, i1, i2)
		}
		if int(i0) >= len(model.UVs) || int(i1) >= len(model.UVs) || int(i2) >= len(model.UVs) {
			log.Fatalf("UV index out of bounds: %d, %d, %d", i0, i1, i2)
		}

		// Draw each vertex of the triangle
		for _, index := range []uint32{i0, i1, i2} {
			if int(index) >= len(model.Vertices) || int(index) >= len(model.UVs) {
				log.Fatalf("Index out of bounds: %d", index)
			}

			vert := model.Vertices[index]
			uv := model.UVs[index]

			// Set texture coordinates
			gl.TexCoord2f(uv.X, uv.Y)
			// Set vertex position
			gl.Vertex3f(vert.X, vert.Y, vert.Z)
		}
	}

	gl.End()

	// Disable texture and blending after rendering
	gl.Disable(gl.TEXTURE_2D)
	gl.Disable(gl.BLEND)

	gl.PopMatrix()

	if err := gl.GetError(); err != gl.NO_ERROR {
		log.Fatalf("OpenGL Error: %v\n", err)
	}
}
