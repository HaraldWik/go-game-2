package d3d

import (
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

type Transform3D struct {
	Pos, Scale, Rot vec3.Type
}

func NewTransform3D(pos, scale, rot vec3.Type) Transform3D {
	return Transform3D{
		Pos:   pos,
		Scale: scale,
		Rot:   rot,
	}
}
