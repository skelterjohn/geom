package geom

type Triangle struct {
	A, B, C Point	
}

func (me *Triangle) Bounds() (bounds *Rect) {
	bounds = &Rect{me.A, me.A}
	bounds.ExpandToContainPoint(me.B)
	bounds.ExpandToContainPoint(me.C)
	return
}

func (me *Triangle) Vertices() (vertices []Point) {
	vertices = []Point{me.A, me.B, me.C}
	return
}

func (me *Triangle) ContainsPoint(p Point) bool {
	leftA := CrossProduct(me.B.Minus(me.A), p.Minus(me.A)) < 0
	leftB := CrossProduct(me.C.Minus(me.B), p.Minus(me.B)) < 0
	leftC := CrossProduct(me.A.Minus(me.C), p.Minus(me.C)) < 0
	return leftA == leftB && leftA == leftC
}

func (me *Triangle) HasVertex(v Point) bool {
	return v.Equals(me.A) || v.Equals(me.B) || v.Equals(me.C)	
}