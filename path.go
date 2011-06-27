package geom

import (
	"math"
)

type Path struct {
	vertices []Point
	bounds   Rect
}

//uncomment to check interface fulfillment
//var _ Bounded = &Path{}

func (p *Path) Equals(oi interface{}) bool {
	o, ok := oi.(*Path)
	if !ok { return false }
	
	if len(p.vertices) != len(o.vertices) { return false }
	
	for i := range p.vertices {
		if !p.vertices[i].Equals(o.vertices[i])	{
			return false
		}
	}
	
	return true
}

func (p *Path) Length() int {
	return len(p.vertices)
}

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

func (p *Path) InsertVertexAfter(v Point, index int) {
	p.vertices = append(p.vertices, v)
	copy(p.vertices[index+1:], p.vertices[index:len(p.vertices)-1])
	p.vertices[index] = v
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

func (me *Path) Error(other *Path) (offset Point, error float64) {
	
	meCenter := me.bounds.Center()
	oCenter := other.bounds.Center()
	
	offset = meCenter.Minus(oCenter)
	if len(me.vertices) != len(other.vertices) {
		error = math.Inf(1)
		return
	}
	
	for i, mv := range me.vertices {
		ov := other.vertices[i]
		offsetMe := mv.Minus(meCenter)
		offsetOther := ov.Minus(oCenter)
		error += offsetMe.DistanceFrom(offsetOther)
	}
	
	return
}