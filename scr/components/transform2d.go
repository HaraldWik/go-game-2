package component

import (
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
)

type Transform2D struct {
	Pos, Size vec2.Type
	Rot       float32
}

func (transform *Transform2D) Update() {}

func (transform *Transform2D) Set() int {
	return 0
}
