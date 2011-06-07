package geom

type Rect struct {
	Min, Max Point
}

func (r *Rect) Contains(p Point) bool {
	return p.PPOf(r.Min) && p.MMOf(r.Max)
}

func (r *Rect) Translate(offset Point) {
	r.Min = r.Min.Plus(offset)
	r.Max = r.Max.Plus(offset)
}

func (r *Rect) ExpandToContain(p Point) {
	r.Min.X = minf(r.Min.X, p.X)
	r.Min.Y = minf(r.Min.Y, p.Y)
	r.Max.X = maxf(r.Max.X, p.X)
	r.Max.Y = maxf(r.Max.Y, p.Y)
}

func RectsIntersect(r1, r2 *Rect) bool {
	return r1.Contains(r2.Min) || r1.Contains(r2.Max) ||
		r2.Contains(r1.Min) || r2.Contains(r1.Max)
}
