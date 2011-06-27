package geom

import (
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Equals(q Point) bool {
	return p.X == q.X && p.Y == q.Y
}

func (p Point) DistanceFrom(q Point) (d float64) {
	return p.Minus(q).Magnitude()
}

func (p Point) Magnitude() (m float64) {
	ss := p.X*p.X + p.Y*p.Y
	m = math.Sqrt(ss)
	return
}

func (p Point) Minus(q Point) (r Point) {
	r.X = p.X - q.X
	r.Y = p.Y - q.Y
	return
}

func (p Point) Plus(q Point) (r Point) {
	r.X = p.X + q.X
	r.Y = p.Y + q.Y
	return
}

func (p Point) Times(s float64) (r Point) {
	r.X = p.X * s
	r.Y = p.Y * s
	return
}

func (p Point) QuadPP(q Point) bool {
	return q.X >= p.X && q.Y >= p.Y
}

func (p Point) QuadPM(q Point) bool {
	return q.X >= p.X && q.Y <= p.Y
}

func (p Point) QuadMP(q Point) bool {
	return q.X <= p.X && q.Y >= p.Y
}

func (p Point) QuadMM(q Point) bool {
	return q.X <= p.X && q.Y <= p.Y
}

func DotProduct(p, q Point) (r float64) {
	r = p.X*q.X + p.Y*q.Y
	return
}

func CrossProduct(p, q Point) (z float64) {
	z = p.X*q.Y - p.Y*q.X
	return
}

func VectorAngle(X, Y Point) (r float64) {
	XdotY := DotProduct(X, Y)
	mXmY := X.Magnitude() * Y.Magnitude()
	r = math.Acos(XdotY / mXmY)
	z := CrossProduct(X, Y)
	if z < 0 {
		r *= -1
	}
	return
}

func VertexAngle(A, B, C Point) (r float64) {
	X := A.Minus(B)
	Y := C.Minus(B)
	r = VectorAngle(X, Y)
	if r < 0 {
		r *= -1
	}
	return
}
