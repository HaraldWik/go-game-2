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
func (org Type) Add(vecs ...Type) Type {
	for _, vec := range vecs {
		org.X += vec.X
		org.Y += vec.Y
		org.Z += vec.Z
	}
	return org
}

// *Subtraction
func (org Type) Sub(vecs ...Type) Type {
	for _, vec := range vecs {
		org.X -= vec.X
		org.Y -= vec.Y
		org.Z -= vec.Z
	}
	return org
}

// *Multiplication
func (org Type) Mul(vecs ...Type) Type {
	for _, vec := range vecs {
		org.X *= vec.X
		org.Y *= vec.Y
		org.Z *= vec.Z
	}
	return org
}

// *Divition
func (org Type) Div(vecs ...Type) Type {
	for _, vec := range vecs {
		org.X *= vec.X
		org.Y *= vec.Y
		org.Z *= vec.Z
	}
	return org
}

// *Scale
func (org Type) Scale(scalar float32) Type {
	return New(
		org.X*scalar,
		org.Y*scalar,
		org.Z*scalar,
	)
}

// *Dot product
func (org Type) Dot(other Type) float32 {
	return float32(
		org.X*other.X +
			org.Y*other.Y +
			org.Z*other.Z,
	)
}

// *Cross product
func (org Type) Cross(other Type) Type {
	return New(
		org.Y*other.Z-org.Z*other.Y,
		org.Z*other.X-org.X*other.Z,
		org.X*other.Y-org.Y*other.X,
	)
}

// *Absolut
func (org Type) ABS() Type {
	return New(
		float32(math.Abs(float64(org.X))),
		float32(math.Abs(float64(org.Y))),
		float32(math.Abs(float64(org.Y))),
	)
}

// *Negativ
func (org Type) Neg() Type {
	return New(
		-float32(math.Abs(float64(org.X))),
		-float32(math.Abs(float64(org.Y))),
		-float32(math.Abs(float64(org.Z))),
	)
}

// *Length
func (org Type) Length() float32 {
	return float32(math.Sqrt(
		float64(org.X*org.X) +
			float64(org.Y*org.Y) +
			float64(org.Z*org.Z)),
	)
}

// *Normalize
func (org Type) Norm() Type {
	if org.Length() != 0 {
		return New(
			org.X/org.Length(),
			org.Y/org.Length(),
			org.Z/org.Length(),
		)
	} else {
		return Zero()
	}
}
