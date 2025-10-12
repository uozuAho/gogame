package game

import "math"

// Really just a 2D coord. Could be a point or a vector.
type Point2D struct {
	X float64
	Y float64
}

func (p Point2D) Len() float64 {
	return math.Hypot(p.X, p.Y)
}

func (p Point2D) UnitVec() Point2D {
	len := math.Hypot(p.X, p.Y)
	if len == 0 {
		return Point2D{0, 0}
	}
	return Point2D{p.X / len, p.Y / len}
}

func (p1 *Point2D) Copy() Point2D {
	return Point2D{p1.X, p1.Y}
}

func (p1 *Point2D) Multiply(factor float64) {
	p1.X *= factor
	p1.Y *= factor
}

func (p1 *Point2D) Add(p2 Point2D) {
	p1.X += p2.X
	p1.Y += p2.Y
}

func (p1 *Point2D) Subtract(p2 Point2D) {
	p1.X -= p2.X
	p1.Y -= p2.Y
}
