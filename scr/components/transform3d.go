package component

import (
	vec3 "github.com/HaraldWik/go-game-2/scr/vector/3"
)

type Transform3D struct {
	Pos, Size, Rot vec3.Type
}

func (transform *Transform3D) Update() {}
