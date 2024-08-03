package s2d

import (
	"github.com/HaraldWik/go-game-2/scr/ups"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
	"github.com/go-gl/gl/v2.1/gl"
)

type Skybox2D struct{}

func (s Skybox2D) Update(obj *ups.Object, deltaTime float32) {
	var (
		color = obj.Data.Get("Color").(vec3.Type)
	)

	gl.ClearColor(color.X, color.Y, color.Z, 1.0)
}
