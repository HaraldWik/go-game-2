package vec2

import "math"

type Type struct {
	X, Y float32
}

func New(x, y float32) Type {
	return Type{
		X: x,
		Y: y,
	}
}

func Zero() Type {
	return Type{}
}

func All(number float32) Type {
	return Type{
		X: number,
		Y: number,
	}
}

// *Addition
func (org Type) Add(vecs ...Type) Type {
	for _, vec := range vecs {
		org.X += vec.X
		org.Y += vec.Y
	}
	return org
}

// *Subtraction
func (org Type) Sub(vecs ...Type) Type {
	for _, vec := range vecs {
		org.X -= vec.X
		org.Y -= vec.Y
	}
	return org
}

// *Multiplication
func (org Type) Mul(vecs ...Type) Type {
	for _, vec := range vecs {
		org.X *= vec.X
		org.Y *= vec.Y
	}
	return org
}

// *Divition
func (org Type) Div(vecs ...Type) Type {
	for _, vec := range vecs {
		org.X *= vec.X
		org.Y *= vec.Y
	}
	return org
}

// *Scale
func (org Type) Scale(scalar float32) Type {
	return New(
		org.X*scalar,
		org.Y*scalar,
	)
}

// *Dot product
func (org Type) Dot(other Type) float32 {
	return float32(
		org.X*other.X +
			org.Y*other.Y,
	)
}

// *Cross product
func (org Type) Cross(other Type) Type {
	return New(
		org.X*other.Y-org.Y*other.X,
		org.X*other.Y-org.Y*other.X,
	)
}

// *Absolut
func (org Type) ABS() Type {
	return New(
		float32(math.Abs(float64(org.X))),
		float32(math.Abs(float64(org.Y))),
	)
}

// *Negativ
func (org Type) Neg() Type {
	return New(
		-float32(math.Abs(float64(org.X))),
		-float32(math.Abs(float64(org.Y))),
	)
}

// *Length
func (org Type) Length() float32 {
	return float32(math.Sqrt(
		float64(org.X*org.X) +
			float64(org.Y*org.Y)),
	)
}

// *Normalize
func (org Type) Norm() Type {
	if org.Length() != 0 {
		return New(
			org.X/org.Length(),
			org.Y/org.Length(),
		)
	} else {
		return Zero()
	}
}
