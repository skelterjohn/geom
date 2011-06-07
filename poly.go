package geom

type Polygon struct {
	vertices []Point
	bounds    Rect
}

func (p *Polygon) AddVertex(v Point) {
	if len(p.vertices) == 0 {
		p.bounds = Rect{
			Min: v,
			Max: v,
		}
	} else {
		p.bounds.ExpandToContain(v)
	}
	p.vertices = append(p.vertices, v)
}

func (p *Polygon) Bounds() (bounds *Rect) {
	return &p.bounds
}

func (p *Polygon) Vertices() (v []Point) {
	v = p.vertices
	return
}