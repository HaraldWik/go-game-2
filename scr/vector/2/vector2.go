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
func (original Type) Add(vectors ...Type) Type {
	for _, vector := range vectors {
		original.X += vector.X
		original.Y += vector.Y
	}
	return original
}

// *Subtraction
func (original Type) Sub(vectors ...Type) Type {
	for _, vector := range vectors {
		original.X -= vector.X
		original.Y -= vector.Y
	}
	return original
}

// *Multiplication
func (original Type) Mul(vectors ...Type) Type {
	for _, vector := range vectors {
		original.X *= vector.X
		original.Y *= vector.Y
	}
	return original
}

// *Divition
func (original Type) Div(vectors ...Type) Type {
	for _, vector := range vectors {
		original.X *= vector.X
		original.Y *= vector.Y
	}
	return original
}

// *Scale
func (original Type) Scale(scalar float32) Type {
	return New(
		original.X*scalar,
		original.Y*scalar,
	)
}

// *Dot product
func (original Type) Dot(other Type) float32 {
	return float32(
		original.X*other.X +
			original.Y*other.Y,
	)
}

// *Cross product
func (original Type) Cross(other Type) Type {
	return New(
		original.X*other.Y-original.Y*other.X,
		original.X*other.Y-original.Y*other.X,
	)
}

// *Absolut
func (original Type) ABS() Type {
	return New(
		float32(math.Abs(float64(original.X))),
		float32(math.Abs(float64(original.Y))),
	)
}

// *Negativ
func (original Type) Neg() Type {
	return New(
		-float32(math.Abs(float64(original.X))),
		-float32(math.Abs(float64(original.Y))),
	)
}

// *Length
func (original Type) Length() float32 {
	return float32(math.Sqrt(
		float64(original.X*original.X) +
			float64(original.Y*original.Y)),
	)
}

// *Normalize
func (original Type) Norm() Type {
	if original.Length() != 0 {
		return New(
			original.X/original.Length(),
			original.Y/original.Length(),
		)
	} else {
		return Zero()
	}
}
