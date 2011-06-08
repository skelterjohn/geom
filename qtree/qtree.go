package qtree

import (
	"github.com/skelterjohn/geom"
)

type Tree struct {
	height int

	Bounds    *geom.Rect
	Partition geom.Point

	Subtrees [4]*Tree

	Elements map[geom.Bounded]bool
}

func New(height int, bounds *geom.Rect) (me *Tree) {
	me = &Tree{
		height:    height,
		Bounds:    bounds,
		Partition: bounds.Min.Plus(bounds.Max).Times(0.5),
	}
	return
}

func (me *Tree) Insert(element geom.Bounded) (inserted bool) {
	if !geom.RectsIntersect(me.Bounds, element.Bounds()) {
		return
	}

	me.Elements[element] = true

	//no subtrees if height is zero
	if me.height == 0 {
		return
	}

	for i, t := range me.Subtrees {
		if t == nil {
			subbounds := *me.Bounds
			switch i {
			case 0:
				subbounds.Min.X = me.Partition.X
				subbounds.Min.Y = me.Partition.Y
			case 1:
				subbounds.Min.X = me.Partition.X
				subbounds.Max.Y = me.Partition.Y
			case 2:
				subbounds.Max.X = me.Partition.X
				subbounds.Min.Y = me.Partition.Y
			case 3:
				subbounds.Max.X = me.Partition.X
				subbounds.Max.Y = me.Partition.Y
			}
			t = New(me.height-1, &subbounds)
			me.Subtrees[i] = t
		}

		t.Insert(element)
	}

	return
}

func (me *Tree) Remove(element geom.Bounded) {
	if !geom.RectsIntersect(me.Bounds, element.Bounds()) {
		return
	}
	me.Elements[element] = false, false
	for _, t := range me.Subtrees {
		if t == nil {
			continue
		}
		t.Remove(element)
	}
}

func (me *Tree) Collect(bounds *geom.Rect, collection map[geom.Bounded]bool) {
	if !geom.RectsIntersect(bounds, me.Bounds) {
		return
	}

	//only check at height zero
	if me.height == 0 {
		for elem, ok := range me.Elements {
			if !ok {
				panic("forgot to delete element properly")
			}
			if geom.RectsIntersect(bounds, elem.Bounds()) {
				collection[elem] = true
			}
		}
		return
	}

	for _, t := range me.Subtrees {
		if t == nil {
			continue
		}
		t.Collect(bounds, collection)
	}
}
