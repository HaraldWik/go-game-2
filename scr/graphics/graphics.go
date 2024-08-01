package gfx

import (
	"log"
	"sort"

	d2d "github.com/HaraldWik/go-game-2/scr/2d/data"
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	"github.com/go-gl/gl/v2.1/gl"
)

type gfx2D struct {
	Objects []*ups.Object
}

var GFX2D gfx2D

func (g *gfx2D) AddObject(obj *ups.Object) {
	g.Objects = append(g.Objects, obj)
}

func (g *gfx2D) DrawCycle() {
	sort.Slice(g.Objects, func(i, j int) bool {
		return g.Objects[i].Data.Get("Material").(d2d.Material2D).Z < g.Objects[j].Data.Get("Material").(d2d.Material2D).Z
	})

	for _, obj := range g.Objects {
		var (
			transform = obj.Data.Get("Transform").(d2d.Transform2D)
			material  = obj.Data.Get("Material").(d2d.Material2D)
			vertices  = obj.Data.Get("Vertices").([]vec2.Type)
		)

		if len(vertices) < 3 {
			log.Fatalf("RenderMesh2D requires at least 3 vertices")
		}

		gl.MatrixMode(gl.MODELVIEW)
		gl.PushMatrix()

		gl.LoadIdentity()
		gl.Translatef(transform.Position.X, transform.Position.Y, 0.0)
		gl.Rotatef(-transform.Rotation, 0.0, 0.0, 1.0)
		gl.Scalef(transform.Scale.X/2.0, transform.Scale.Y/2.0, 1.0)

		gl.Enable(gl.TEXTURE_2D)
		gl.Enable(gl.BLEND)
		gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
		gl.BindTexture(gl.TEXTURE_2D, material.Texture.Image)

		gl.Color3f(material.Alpha.X, material.Alpha.Y, material.Alpha.Z)

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

		centerX := (minX + maxX) / 2
		centerY := (minY + maxY) / 2

		gl.Begin(gl.TRIANGLE_FAN)
		for _, vertex := range vertices {
			adjustedX := vertex.X - centerX
			adjustedY := vertex.Y - centerY

			tx := (vertex.X - minX) / width
			ty := 1.0 - (vertex.Y-minY)/height
			gl.TexCoord2f(tx, ty)
			gl.Vertex2f(adjustedX, adjustedY)
		}
		gl.End()

		gl.Disable(gl.TEXTURE_2D)
		gl.Disable(gl.BLEND)

		gl.PopMatrix()

		if err := gl.GetError(); err != gl.NO_ERROR {
			log.Fatalf("OpenGL Error: %v\n", err)
		}
	}
	g.Objects = g.Objects[:0]
}
