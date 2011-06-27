package geom

import (
	"testing"
)

func TestRectsIntersect(t *testing.T) {
	r1 := &Rect{Point{0, 0}, Point{650, 650}}
	r2 := &Rect{Point{200, 500}, Point{450, 750}}
	if !r1.Min.QuadPP(r2.Min) {
		t.Error("QuadPP")
	}
	if !r1.Max.QuadMM(r2.Min) {
		t.Error("QuadMM")
	}
	if !r1.Contains(r2.Min) {
		t.Error("contains")
	}
	if !RectsIntersect(r1, r2) {
		t.Error("intersect")
	}
	
	r1 = &Rect{Point{325, 325}, Point{650, 650}}
	r2 = &Rect{Point{200, 500}, Point{450, 750}}
	
	Debug = true
	
	if !r1.Min.QuadPP(r2.Min) {
		t.Error("QuadPP2")
	}
	if !r1.Max.QuadMM(r2.Min) {
		t.Error("QuadMM2")
	}
	if !r1.Contains(r2.Min) {
		t.Error("contains2")
	}
	if !RectsIntersect(r1, r2) {
		t.Error("intersect2")
	}
	
}