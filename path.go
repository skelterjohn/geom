package geom


type Path struct {
	vertices []Point
	bounds   Rect
}

//uncomment to check interface fulfillment
//var _ Bounded = &Path{}

func (p *Path) AddVertex(v Point) {
	if len(p.vertices) == 0 {
		p.bounds = Rect{
			Min: v,
			Max: v,
		}
	} else {
		p.bounds.ExpandToContainPoint(v)
	}
	p.vertices = append(p.vertices, v)
}

func (p *Path) Bounds() (bounds *Rect) {
	return &p.bounds
}

func (p *Path) Vertices() (v []Point) {
	v = p.vertices
	return
}

func (p *Path) Translate(offset Point) {
	p.bounds.Translate(offset)
	for i, v := range p.vertices {
		p.vertices[i] = v.Plus(offset)
	}
}

func (p *Path) Scale(factor float64) {
	p.bounds.Scale(factor)
	for i, v := range p.vertices {
		p.vertices[i] = v.Times(factor)
	}
}
