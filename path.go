// Copyright 2009 The geom Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import (
	"math"
)

type Path struct {
	vertices []Point
	bounds   Rect
}

func (p *Path) Translate(offset Point) {
	p.bounds.Translate(offset)
	for i := range p.vertices {
		p.vertices[i].Translate(offset)
	}
}

func (p *Path) Rotate(rad float64) {
	for i := range p.vertices {
		p.vertices[i].Rotate(rad)
	}
	p.bounds = Rect{p.vertices[0], p.vertices[0]}
	p.bounds.ExpandToContain(PointChan(p.vertices[1:]))
}

func (p *Path) Scale(xf, yf float64) {
	
	for i := range p.vertices {
		p.vertices[i].Scale(xf, yf)
	}
	p.bounds.Scale(xf, yf)
}

func (p *Path) Clone() (op *Path) {
	op = &Path{}
	op.bounds = *p.bounds.Clone()
	op.vertices = append([]Point{}, p.vertices...)
	return
}

//uncomment to check interface fulfillment
//var _ Bounded = &Path{}

func (p *Path) Equals(oi interface{}) bool {
	o, ok := oi.(*Path)
	if !ok { return false }
	
	if len(p.vertices) != len(o.vertices) { return false }
	
	for i := range p.vertices {
		if !p.vertices[i].EqualsPoint(o.vertices[i])	{
			return false
		}
	}
	
	return true
}

func (p *Path) Register(op *Path) (offset Point, match bool ) {
	offset = p.bounds.Min.Minus(op.bounds.Min)
	if len(p.vertices) != len(op.vertices) {
		dbg("registure failure: wrong counts")
		return // with match = false
	}
	for i := range p.vertices {
		if !p.vertices[i].EqualsPoint(op.vertices[i].Plus(offset)) {
			dbg("register failure: v1=%v v2=%v offset=%v", p.vertices[i], op.vertices[i], offset)
			return // with match = false
		}
	}
	match = true
	return
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
