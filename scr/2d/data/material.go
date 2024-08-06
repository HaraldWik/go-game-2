package d2d

import (
	_ "image/jpeg"
	_ "image/png"

	load "github.com/HaraldWik/go-game-2/scr/loaders"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
	vec4 "github.com/HaraldWik/go-game-2/scr/vector/4"
)

type Material2D struct {
	Texture load.Texture
	Region  vec4.Type

	Color   vec3.Type
	Opacity float32

	Z float32
}

func NewMaterial2D(texture load.Texture, region vec4.Type, color vec3.Type, opacity, z float32) Material2D {
	if region == vec4.Zero() {
		region = vec4.New(
			0,
			0,
			texture.Size.X,
			texture.Size.Y,
		)
	}

	return Material2D{
		Texture: texture,
		Region:  region,

		Color:   color,
		Opacity: opacity,

		Z: z,
	}
}
