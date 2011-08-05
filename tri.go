// Copyright 2009 The geom Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

func (me *Triangle) Equals(oi interface{}) bool {
	ot, ok := oi.(*Triangle)
	if !ok { return false }
	if me.A.EqualsPoint(ot.A) {
		if me.B.EqualsPoint(ot.B) {
			return me.C.EqualsPoint(ot.C)
		}
		if me.B.EqualsPoint(ot.C) {
			return me.C.EqualsPoint(ot.B)
		}
	}
	if me.A.EqualsPoint(ot.B) {
		if me.B.EqualsPoint(ot.A) {
			return me.C.EqualsPoint(ot.C)
		}
		if me.B.EqualsPoint(ot.C) {
			return me.C.EqualsPoint(ot.A)
		}
	}
	if me.A.EqualsPoint(ot.C) {
		if me.B.EqualsPoint(ot.B) {
			return me.C.EqualsPoint(ot.A)
		}
		if me.B.EqualsPoint(ot.A) {
			return me.C.EqualsPoint(ot.B)
		}
	}
	return false
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
	return v.EqualsPoint(me.A) || v.EqualsPoint(me.B) || v.EqualsPoint(me.C)	
}