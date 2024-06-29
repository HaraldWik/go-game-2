package vec3

import (
	"math"
)

type Type struct {
	X, Y, Z float32
}

func New(x, y, z float32) Type {
	return Type{
		X: x,
		Y: y,
		Z: z,
	}
}

func Zero() Type {
	return Type{}
}

func All(number float32) Type {
	return Type{
		X: number,
		Y: number,
		Z: number,
	}
}

// *Addition
func (original Type) Add(vectors ...Type) Type {
	for _, vector := range vectors {
		original.X += vector.X
		original.Y += vector.Y
		original.Z += vector.Z
	}
	return original
}

// *Subtraction
func (original Type) Sub(vectors ...Type) Type {
	for _, vector := range vectors {
		original.X -= vector.X
		original.Y -= vector.Y
		original.Z -= vector.Z
	}
	return original
}

// *Multiplication
func (original Type) Mul(vectors ...Type) Type {
	for _, vector := range vectors {
		original.X *= vector.X
		original.Y *= vector.Y
		original.Z *= vector.Z
	}
	return original
}

// *Divition
func (original Type) Div(vectors ...Type) Type {
	for _, vector := range vectors {
		original.X *= vector.X
		original.Y *= vector.Y
		original.Z *= vector.Z
	}
	return original
}

// *Scale
func (original Type) Scale(scalar float32) Type {
	return New(
		original.X*scalar,
		original.Y*scalar,
		original.Z*scalar,
	)
}

// *Dot product
func (original Type) Dot(other Type) float32 {
	return float32(
		original.X*other.X +
			original.Y*other.Y +
			original.Z*other.Z,
	)
}

// *Cross product
func (original Type) Cross(other Type) Type {
	return New(
		original.Y*other.Z-original.Z*other.Y,
		original.Z*other.X-original.X*other.Z,
		original.X*other.Y-original.Y*other.X,
	)
}

// *Absolut
func (original Type) ABS() Type {
	return New(
		float32(math.Abs(float64(original.X))),
		float32(math.Abs(float64(original.Y))),
		float32(math.Abs(float64(original.Y))),
	)
}

// *Negativ
func (original Type) Neg() Type {
	return New(
		-float32(math.Abs(float64(original.X))),
		-float32(math.Abs(float64(original.Y))),
		-float32(math.Abs(float64(original.Z))),
	)
}

// *Length
func (original Type) Length() float32 {
	return float32(math.Sqrt(
		float64(original.X*original.X) +
			float64(original.Y*original.Y) +
			float64(original.Z*original.Z)),
	)
}

// *Normalize
func (original Type) Norm() Type {
	if original.Length() != 0 {
		return New(
			original.X/original.Length(),
			original.Y/original.Length(),
			original.Z/original.Length(),
		)
	} else {
		return Zero()
	}
}
