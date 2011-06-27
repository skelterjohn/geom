package geom

import (
	"testing"
	"fmt"
)

func TestVertexAngle(t *testing.T) {
	A := Point{0, 0}
	B := Point{1, 0}
	C := Point{1, 1}
	D := Point{1, -1}
	r1 := VertexAngle(A, B, C)
	r2 := VertexAngle(A, B, D)
	if r1 != -r2 {

	}

	p1 := Point{1, 2}
	p2 := Point{0, 3}
	p3 := Point{0, 0}
	p4 := Point{1, 1}

	rp := VertexAngle(p1, p2, p3)
	rn := VertexAngle(p4, p2, p3)
	fmt.Println(rp, rn)
}

func TestVectorAngle(t *testing.T) {
	v1 := Point{0, -1}
	v2 := Point{-1, 0}
	v3 := Point{1, -1}
	v4 := Point{-1, 0}

	a12 := VectorAngle(v1, v2)
	a34 := VectorAngle(v3, v4)

	fmt.Println(a12, a34)
}
