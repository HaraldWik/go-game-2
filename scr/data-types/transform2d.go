package dt

import vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"

type Transform2D struct {
	Pos, Size vec2.Type
	Rot       float32
}

func NewTransform2D(pos, size vec2.Type, rot float32) Transform2D {
	return Transform2D{Pos: pos, Size: size, Rot: rot}
}
