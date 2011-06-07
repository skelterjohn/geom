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
	dx := p.X - q.X
	dy := p.Y - q.Y
	ss := dx*dx + dy*dy
	d = math.Sqrt(ss)
	return
}

func (p Point) Minux(q Point) (r Point) {
	r.X = p.X - q.X
	r.Y = p.Y - q.Y
	return
}

func (p Point) Plus(q Point) (r Point) {
	r.X = p.X + q.X
	r.Y = p.Y + q.Y
	return
}

func (p Point) Dot(q Point) (r float64) {
	r = p.X*q.X + p.Y*q.Y
	return
}

func (p Point) Times(s float64) (r Point) {
	r.X = p.X * s
	r.Y = p.Y * s
	return
}

func (p Point) PPOf(q Point) bool {
	return q.X >= p.X && q.Y >= p.Y
}

func (p Point) PMOf(q Point) bool {
	return q.X >= p.X && q.Y <= p.Y
}

func (p Point) MPOf(q Point) bool {
	return q.X <= p.X && q.Y >= p.Y
}

func (p Point) MMOf(q Point) bool {
	return q.X <= p.X && q.Y <= p.Y
}
