package geom

import (
	"container/ring"
)

type Polygon struct {
	Path
}

func (me *Polygon) Vertex(index int) (v Point) {
	index = index % len(me.vertices)
	if index < 0 {
		index = len(me.vertices)+index
	}
	v = me.vertices[index]
	return
}

func (me *Polygon) VertexAngle(index int) (r float64) {
	a := me.Vertex(index-1)
	b := me.Vertex(index)
	c := me.Vertex(index+1)
	r = VertexAngle(a, b, c)
	return
}

func (me *Polygon) WindingOrder() (winding float64) {
	for i := 0; i < len(me.vertices); i++ {
		winding += me.VertexAngle(i)
	}
	return	
}

func (me *Polygon) Triangles() (tris []Triangle) {
	
	vr := ring.New(len(me.vertices))
	
	for _, v := range me.vertices {
		vr.Value = v
		vr = vr.Next()
	}
	
	leftMost := func(vr *ring.Ring) (lvr *ring.Ring) {
		lvr = vr
		for cr := vr.Next(); cr != vr; cr = cr.Next() {
			if cr.Value.(Point).X < lvr.Value.(Point).X {
				lvr = cr
			}
		}
		return
	}
	
	for vr.Len() > 2 {
		lvr := leftMost(vr)
		
		pvr := lvr.Prev()
		nvr := lvr.Next()
		//a left-most point is a convex vertex
		lp := lvr.Value.(Point)
		pp := pvr.Value.(Point)
		np := nvr.Value.(Point)
		
		//candidate triangle - check if it contains any other points
		ctri := Triangle{pp, lp, np}
		//leftmost contained vertex - nil initially
		var lcv *ring.Ring
		for lcv == nil {
		
			for cr := lvr.Next(); cr != lvr.Prev(); cr = cr.Next() {
				cv := cr.Value.(Point)
				if ctri.HasVertex(cv) {
					continue
				}
				if ctri.ContainsPoint(cv) {
					if lcv == nil || cv.X < lcv.Value.(Point).X {
						lcv = cr	
					}
				}
			}
			
			if lcv == nvr {
				
			}
		}
		
		
		//otherwise
	}
	
	return	
}