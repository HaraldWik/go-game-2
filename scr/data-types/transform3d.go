package dt

import (
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

type Transform3D struct {
	Pos, Size, Rot vec3.Type
}

func NewTransform3D(pos, size, rot vec3.Type) Transform3D {
	return Transform3D{
		Pos:  pos,
		Size: size,
		Rot:  rot,
	}
}
