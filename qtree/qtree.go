package qtree

import (
	"fmt"
	"github.com/skelterjohn/geom"
)

var Debug = false
var Indent = 0

func dbg(format string, args ...interface{}) {
	if Debug {
		buf := make([]byte, Indent)
		for i := range buf {
			buf[i] = ' '
		}
		fmt.Printf(string(buf)+format+"\n", args...)
	}
}

type Config struct {
	SplitCount int
	SplitSizeRatio float64
	Height int
}

func ConfigDefault() (cfg Config) {
	cfg.SplitCount = 5
	cfg.SplitSizeRatio = 0.5
	cfg.Height = 10
	return	
}

type Tree struct {
	height int
	cfg Config

	Bounds    *geom.Rect
	Partition geom.Point

	Subtrees [4]*Tree

	Elements map[geom.Bounded]bool
	BigElements map[geom.Bounded]bool
}

func New(cfg Config, bounds *geom.Rect) (me *Tree) {
	me = &Tree{
		cfg:       cfg,
		Bounds:    bounds,
		Partition: bounds.Min.Plus(bounds.Max).Times(0.5),
	}
	return
}

func (me *Tree) IsBig(bounds *geom.Rect) bool {
	return bounds.Width() >= me.cfg.SplitSizeRatio * me.Bounds.Width() ||
	       bounds.Height() >= me.cfg.SplitSizeRatio * me.Bounds.Height()
}

func (me *Tree) insertSubTrees(element geom.Bounded) (inserted bool) {
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
			cfg := me.cfg
			cfg.Height--
			t = New(cfg, &subbounds)
			me.Subtrees[i] = t
		}

		inserted = inserted || t.Insert(element)
	}
	
	return
}

func (me *Tree) Insert(element geom.Bounded) (inserted bool) {
	str := ""
	if geom.RectsIntersect(element.Bounds(), me.Bounds) {
		str = "*"
	}
	dbg("inserting in %v%s", *me.Bounds, str)
	
	if !geom.RectsIntersect(me.Bounds, element.Bounds()) {
		return
	}
	
	inserted = true
	
	//if this element is too big, stop here
	if me.IsBig(element.Bounds()) {
		if me.BigElements == nil {
			me.BigElements = make(map[geom.Bounded]bool)
		}
		me.BigElements[element] = true
		return	
	}

	//if we're at the bottom, stop here
	if me.cfg.Height == 0 {
		if me.Elements == nil {
			me.Elements = make(map[geom.Bounded]bool)
		}
		me.Elements[element] = true
		return
	}
	
	//if we've got enough at this level, break into subtrees
	if me.Elements != nil && len(me.Elements) == me.cfg.SplitCount {
		for elem := range me.Elements {
			me.insertSubTrees(elem)	
		}
		me.Elements = nil
	}
	
	//if we already have subtrees, insert into them
	if me.Subtrees[0] != nil {
		me.insertSubTrees(element)
		return
	}
	
	//no subtrees, stop here
	if me.Elements == nil {
		me.Elements = make(map[geom.Bounded]bool)
	}
	me.Elements[element] = true
	
	return
}

func (me *Tree) Remove(element geom.Bounded) (removed bool) {
	dbg("removing %v", element)
	if Debug {
		println(element)	
	}
	Indent++
	defer func() {Indent--}()
	
	if !geom.RectsIntersect(me.Bounds, element.Bounds()) {
		dbg("out of bounds")
		return
	}
	
	dbg("BigElements: %v", me.BigElements)
	dbg("Elements: %v", me.Elements)
	
	if me.BigElements != nil {
		if _, ok := me.BigElements[element]; ok {
			removed = true
			me.BigElements[element] = false, false
			return
		}
	}
	if me.Elements != nil {
		if _, ok := me.Elements[element]; ok {
			removed = true
			me.Elements[element] = false, false
			return
		}
	}
	for _, t := range me.Subtrees {
		if t == nil {
			continue
		}
		if t.Remove(element) {
			removed = true
		}
	}
	
	return
}

func (me *Tree) Enumerate(collection map[geom.Bounded]int) {
	if me.Elements != nil {
		for elem := range me.Elements {
			collection[elem] = collection[elem]+1
		}
	}
	
	if me.BigElements != nil {
		for elem := range me.BigElements {
			collection[elem] = collection[elem]+1
		}
	}
	
	for _, t := range me.Subtrees {
		if t != nil {
			t.Enumerate(collection)
		}
	}
}

func (me *Tree) Do(foo func(x geom.Bounded)) {
	if me.Elements != nil {
		for elem := range me.Elements {
			foo(elem)
		}
	}
	
	if me.BigElements != nil {
		for elem := range me.BigElements {
			foo(elem)
		}
	}
	
	for _, t := range me.Subtrees {
		if t != nil {
			t.Do(foo)
		}
	}
}

func (me *Tree) CollectInside(bounds *geom.Rect, collection map[geom.Bounded]bool) {
	if !geom.RectsIntersect(bounds, me.Bounds) {
		return
	}
	if me.BigElements != nil {
		for elem := range me.BigElements {
			if bounds.ContainsRect(elem.Bounds()) {
				collection[elem] = true
				return
			}
		}
	}
	if me.Elements != nil {
		for elem := range me.Elements {
			if bounds.ContainsRect(elem.Bounds()) {
				collection[elem] = true
				return
			}
		}
	}
	
	for _, t := range me.Subtrees {
		if t == nil {
			continue
		}
		t.CollectInside(bounds, collection)
	}
	
	return
}

func (me *Tree) CollectIntersect(bounds *geom.Rect, collection map[geom.Bounded]bool) (found bool) {
	str := ""
	if geom.RectsIntersect(bounds, me.Bounds) {
		str = "*"
	}
	dbg("looking in %v%s", *me.Bounds, str)
	Indent++
	defer func() {
		if found {
			dbg("found in %v", *me.Bounds)
		}
		Indent--
	}()
	
	
	if !geom.RectsIntersect(bounds, me.Bounds) {
		return
	}
	if me.BigElements != nil {
		for elem := range me.BigElements {
			str = ""
			if geom.RectsIntersect(elem.Bounds(), bounds) {
				collection[elem] = true
				found = true
				str = "*"
			}
			dbg("big: %v%s", *elem.Bounds(), str)
		}
	}

	if me.Elements != nil {
		for elem, ok := range me.Elements {
			if !ok {
				panic("forgot to delete element properly")
			}
			str = ""
			if geom.RectsIntersect(bounds, elem.Bounds()) {
				collection[elem] = true
				found = true
				str = "*"
			}
			dbg("elems: %v%s", *elem.Bounds(), str)
		}
	}

	if me.Subtrees[0] == nil {
		dbg("no subtrees")
	}

	for _, t := range me.Subtrees {
		if t == nil {
			continue
		}
		found = t.CollectIntersect(bounds, collection) || found
	}
	
	return
}

func (me *Tree) String() string {
	str := "[]"
	if me.Subtrees[0] != nil {
		str = fmt.Sprintf("%v", me.Subtrees)
	}
	return fmt.Sprintf("QTree{%v, %v, %v, %s}", me.Bounds, me.Elements, me.BigElements, str)
}
