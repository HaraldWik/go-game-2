package d3d

import (
	load "github.com/HaraldWik/go-game-2/scr/loaders"
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

type Material3D struct {
	Texture load.Texture

	Alpha vec3.Type
}

func NewMaterial3D(texture load.Texture, alpha vec3.Type) Material3D {
	return Material3D{
		Texture: texture,

		Alpha: alpha,
	}
}
