package s2d

import (
	d2d "github.com/HaraldWik/go-game-2/scr/2d/data"
	gfx "github.com/HaraldWik/go-game-2/scr/graphics"
	load "github.com/HaraldWik/go-game-2/scr/loaders"
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

type RenderMesh2D struct{}

func (t RenderMesh2D) Update(obj *ups.Object, deltaTime float32) {
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

func (r RenderRectangle2D) Update(obj *ups.Object, deltaTime float32) {
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

func (t RenderTriangle2D) Update(obj *ups.Object, deltaTime float32) {
	gfx.GFX2D.AddObject(obj)
}

type RenderText2D struct{}

func (t RenderText2D) Start(obj *ups.Object) {
	vertices := []vec2.Type{
		vec2.New(-1.0, -1.0),
		vec2.New(1.0, -1.0),
		vec2.New(0.0, 1.0),
	}

	obj.Data.Set("Vertices", vertices)
}

func (t RenderText2D) Update(obj *ups.Object, deltaTime float32) {
	var (
		font = obj.Data.Get("Font").(*load.Font)
		text = obj.Data.Get("Text").(string)
	)

	obj.Data.Set("Material", d2d.NewMaterial2D(
		font.RenderTextToTexture(text),
		vec3.New(1.0, 1.0, 1.0),
		15.0,
	))

	gfx.GFX2D.AddObject(obj)
}
