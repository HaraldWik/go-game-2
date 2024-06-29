package vec4

import "math"

type Type struct {
	X, Y, Z, W float32
}

func New(x, y, z, w float32) Type {
	return Type{
		X: x,
		Y: y,
		Z: z,
		W: w,
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
		W: number,
	}
}

// *Addition
func (original Type) Add(vectors ...Type) Type {
	for _, vector := range vectors {
		original.X += vector.X
		original.Y += vector.Y
		original.Z += vector.Z
		original.W += vector.W
	}
	return original
}

// *Subtraction
func (original Type) Sub(vectors ...Type) Type {
	for _, vector := range vectors {
		original.X -= vector.X
		original.Y -= vector.Y
		original.Z -= vector.Z
		original.W -= vector.W
	}
	return original
}

// *Multiplication
func (original Type) Mul(vectors ...Type) Type {
	for _, vector := range vectors {
		original.X *= vector.X
		original.Y *= vector.Y
		original.Z *= vector.Z
		original.W *= vector.W
	}
	return original
}

// *Divition
func (original Type) Div(vectors ...Type) Type {
	for _, vector := range vectors {
		original.X *= vector.X
		original.Y *= vector.Y
		original.Z *= vector.Z
		original.W *= vector.W
	}
	return original
}

// *Scale
func (original Type) Scale(scalar float32) Type {
	return New(
		original.X*scalar,
		original.Y*scalar,
		original.Z*scalar,
		original.W*scalar,
	)
}

// *Dot product
func (original Type) Dot(other Type) float32 {
	return float32(
		original.X*other.X +
			original.Y*other.Y +
			original.Z*other.Z +
			original.W*other.W,
	)
}

// *Cross product
func (original Type) Cross(other Type) Type {
	return New(
		original.Y*other.Z-original.Z*other.Y,
		original.Z*other.X-original.X*other.Z,
		original.X*other.Y-original.Y*other.X,
		original.W*other.Z-original.Z*other.W,
	)
}

// *Absolut
func (original Type) ABS() Type {
	return New(
		float32(math.Abs(float64(original.X))),
		float32(math.Abs(float64(original.Y))),
		float32(math.Abs(float64(original.Y))),
		float32(math.Abs(float64(original.W))),
	)
}

// *Negativ
func (original Type) Neg() Type {
	return New(
		-float32(math.Abs(float64(original.X))),
		-float32(math.Abs(float64(original.Y))),
		-float32(math.Abs(float64(original.Z))),
		-float32(math.Abs(float64(original.W))),
	)
}

// *Length
func (original Type) Length() float32 {
	return float32(math.Sqrt(
		float64(original.X*original.X) +
			float64(original.Y*original.Y) +
			float64(original.Z*original.Z) +
			float64(original.W*original.W)),
	)
}

// *Normalize
func (original Type) Norm() Type {
	if original.Length() != 0 {
		return New(
			original.X/original.Length(),
			original.Y/original.Length(),
			original.Z/original.Length(),
			original.W/original.Length(),
		)
	} else {
		return Zero()
	}
}
