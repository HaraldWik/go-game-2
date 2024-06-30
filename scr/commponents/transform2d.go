package commponent

import (
	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
)

type Transform2D struct {
	Pos, Size vec2.Type
	Rot       float32
}

func (t *Transform2D) Update() {

}
