package s2d

import (
	gfx "github.com/HaraldWik/go-game-2/scr/graphics"
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
)

type RenderMesh2D struct{}

func (t RenderMesh2D) FixedUpdate(obj *ups.Object) {
	gfx.GFX2D.AddObject(obj)
}

type RenderRectangle2D struct{}

func (r RenderRectangle2D) Start(obj *ups.Object) {
	vertices := []vec2.Type{
		vec2.New(-1.0, -1.0),
		vec2.New(1.0, -1.0),
		vec2.New(1.0, 1.0),
		vec2.New(-1.0, 1.0),
	}

	obj.Data.Set("Vertices", vertices)
}

func (r RenderRectangle2D) FixedUpdate(obj *ups.Object) {
	gfx.GFX2D.AddObject(obj)
}

type RenderTriangle2D struct{}

func (t RenderTriangle2D) Start(obj *ups.Object) {
	vertices := []vec2.Type{
		vec2.New(-1.0, -1.0),
		vec2.New(1.0, -1.0),
		vec2.New(1.0, 1.0),
		vec2.New(-1.0, 1.0),
	}

	obj.Data.Set("Vertices", vertices)
}

func (t RenderTriangle2D) FixedUpdate(obj *ups.Object) {
	gfx.GFX2D.AddObject(obj)
}
