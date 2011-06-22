package geom

type Rect struct {
	Min, Max Point
}

func (r *Rect) Width() float64 {
	return r.Max.X - r.Min.X
}

func (r *Rect) Height() float64 {
	return r.Max.Y - r.Min.Y
}

func (r *Rect) Size() (w, h float64) {
	return r.Max.X - r.Min.X, r.Max.Y - r.Min.Y
}

func (r *Rect) Center() (center Point) {
	center.X = 0.5 * (r.Min.X + r.Max.X)
	center.Y = 0.5 * (r.Min.Y + r.Max.Y)
	return
}

func (r *Rect) Contains(p Point) bool {
	return r.Min.QuadPP(p) && r.Max.QuadMM(p)
}

func (r *Rect) Translate(offset Point) {
	r.Min = r.Min.Plus(offset)
	r.Max = r.Max.Plus(offset)
}

func (r *Rect) Scale(factor float64) {
	r.Min = r.Min.Times(factor)
	r.Max = r.Max.Times(factor)
}

func (r *Rect) ExpandToContainPoint(p Point) {
	r.Min.X = minf(r.Min.X, p.X)
	r.Min.Y = minf(r.Min.Y, p.Y)
	r.Max.X = maxf(r.Max.X, p.X)
	r.Max.Y = maxf(r.Max.Y, p.Y)
}

func (r *Rect) ExpandToContainRect(q *Rect) {
	r.ExpandToContainPoint(q.Min)
	r.ExpandToContainPoint(q.Max)
}

func (r *Rect) Bounds() (bounds *Rect) {
	bounds = r
	return
}

func RectsIntersect(r1, r2 *Rect) bool {
	ov := func(min1, max1, min2, max2 float64) (overlap bool) {
		if min1 <= min2 && max1 >= min2 {
			return true
		}
		if min1 <= max2 && max1 >= max2 {
			return true
		}
		if min2 <= min1 && max2 >= min1 {
			return true
		}
		if min2 <= max1 && max2 >= max1 {
			return true
		}
		return false
	}
	dbg("RI(%v, %v)", r1, r2)
	xoverlap := ov(r1.Min.X, r1.Max.X, r2.Min.X, r2.Max.X)
	yoverlap := ov(r1.Min.Y, r1.Max.Y, r2.Min.Y, r2.Max.Y)
	dbg("%v %v", xoverlap, yoverlap)
	return xoverlap && yoverlap
}
