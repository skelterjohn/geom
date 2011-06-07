package geom

type QTree struct {
	height int

	Bounds    Rect
	Partition Point

	Subtrees [4]*QTree

	Elements map[Bounded]bool
}

func NewQTree(height int, bounds Rect) (me *QTree) {
	me = &QTree{
		height:    height,
		Bounds:    bounds,
		Partition: bounds.Min.Plus(bounds.Max).Times(0.5),
	}
	return
}

func (me *QTree) Insert(element Bounded) (inserted bool) {
	if !RectsIntersect(&me.Bounds, element.Bounds()) {
		return
	}

	me.Elements[element] = true

	//no subtrees if height is zero
	if me.height == 0 {
		return
	}

	for i, t := range me.Subtrees {
		if t == nil {
			subbounds := me.Bounds
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
			t = NewQTree(me.height-1, subbounds)
			me.Subtrees[i] = t
		}

		t.Insert(element)
	}

	return
}

func (me *QTree) Remove(element Bounded) {
	if !RectsIntersect(&me.Bounds, element.Bounds()) {
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

func (me *QTree) Collect(bounds *Rect, collection map[Bounded]bool) {
	if !RectsIntersect(bounds, &me.Bounds) {
		return
	}

	//only check at height zero
	if me.height == 0 {
		for elem, ok := range me.Elements {
			if !ok {
				panic("forgot to delete element properly")
			}
			if RectsIntersect(bounds, elem.Bounds()) {
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
