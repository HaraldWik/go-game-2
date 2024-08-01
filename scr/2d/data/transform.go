package d2d

import (
	"math"

	vec2 "github.com/HaraldWik/go-game-2/scr/vector/2"
)

type Transform2D struct {
	Position, Scale vec2.Type
	Rotation        float32
}

func NewTransform2D(pos, scale vec2.Type, rot float32) Transform2D {
	return Transform2D{
		Position: pos,
		Scale:    scale,
		Rotation: rot,
	}
}

func (t *Transform2D) Translate(add vec2.Type) {
	angleRad := float32(t.Rotation * math.Pi / 180.0)

	newPosition := vec2.New(
		t.Position.X+add.X,
		t.Position.Y+add.Y,
	)

	relativeTranslation := vec2.New(
		newPosition.X-t.Position.X,
		newPosition.Y-t.Position.Y,
	)

	cosAngle := float32(math.Cos(float64(-angleRad)))
	sinAngle := float32(math.Sin(float64(-angleRad)))
	rotatedTranslation := vec2.Type{
		X: relativeTranslation.X*cosAngle - relativeTranslation.Y*sinAngle,
		Y: relativeTranslation.X*sinAngle + relativeTranslation.Y*cosAngle,
	}

	t.Position = vec2.New(
		t.Position.X+rotatedTranslation.X,
		t.Position.Y+rotatedTranslation.Y,
	)
}
