package d2d

import (
	_ "image/jpeg"
	_ "image/png"

	load "github.com/HaraldWik/go-game-2/scr/loaders"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

type Material2D struct {
	Texture load.Texture

	Alpha vec3.Type

	Z float32
}

func NewMaterial2D(texture load.Texture, alpha vec3.Type, z float32) Material2D {
	return Material2D{
		Texture: texture,

		Alpha: alpha,

		Z: z,
	}
}
