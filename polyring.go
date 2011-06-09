package geom

import (

)

type PolyRing struct {
	Next, Prev *PolyRing
	Vertex Point
}

func NewPolyRing(poly *Polygon) (pr *PolyRing) {
	pr = &PolyRing {
		Next : pr,
		Prev : pr,
		Vertex : poly.Vertices()[0],
	}
	for _, v := range poly.Vertices()[1:] {
		pr.InsertBefore(v)
	}
	return
}

func (pr *PolyRing) InsertBefore(v Vertex) {
	nr := &PolyRing {
		Next : pr,
		Prev : pr.Prev,
		Vertex : v,
	}
	nr.Next.Prev = nr
	nr.Prev.Next = nr
}

func (pr *PolyRing) InsertAfter(v Vertex) {
	nr := &PolyRing {
		Next : pr.Next,
		Prev : pr,
		Vertex : v,
	}
	nr.Next.Prev = nr
	nr.Prev.Next = nr
}

func (pr* PolyRing) ExtractPolygon(opr *PolyRing) (prs []*PolyRing) {
	return
}
