package geom

import (
	"testing"
	"fmt"
)

func TestInsert(t *testing.T) {
	p := &Polygon{}
	p.AddVertex(Point{0, 0})
	p.AddVertex(Point{0, 1})
	p.AddVertex(Point{1, 1})
	p.AddVertex(Point{1, 0})
	p.InsertVertexAfter(Point{5, 0}, 2)
	
	p2 := &Polygon{}
	p2.AddVertex(Point{0, 0})
	p2.AddVertex(Point{0, 1})
	p2.AddVertex(Point{5, 0})
	p2.AddVertex(Point{1, 1})
	p2.AddVertex(Point{1, 0})
	
	if !p.Equals(p2) {
		t.Fail()
	}
}

func TestPolyTriangularize(t *testing.T) {
	poly := new(Polygon)
	poly.AddVertex(Point{0, 0})
	poly.AddVertex(Point{0, 1})
	poly.AddVertex(Point{1, 1})
	poly.AddVertex(Point{1, 0})
	tris, ok := poly.Triangles()
	if ok {
		fmt.Println()
		for _, tri := range tris {
			fmt.Printf("triangle: %v\n", tri)	
		}
	} else {
		fmt.Printf("No triangles for %v\n", poly)	
	}
	
	poly = new(Polygon)
	poly.AddVertex(Point{0, 0})
	poly.AddVertex(Point{1, 1})
	poly.AddVertex(Point{2, 0})
	poly.AddVertex(Point{2, 3})
	poly.AddVertex(Point{1, 2})
	poly.AddVertex(Point{0, 3})
	tris, ok = poly.Triangles()
	if ok {
		fmt.Println()
		for _, tri := range tris {
			fmt.Printf("triangle: %v\n", tri)	
		}
	} else {
		fmt.Printf("No triangles for %v\n", poly)	
	}
	
	poly = new(Polygon)
	poly.AddVertex(Point{2, 1})
	poly.AddVertex(Point{2, 2})
	poly.AddVertex(Point{1, 2})
	poly.AddVertex(Point{1, 3})
	tris, ok = poly.Triangles()
	if ok {
		fmt.Println()
		for _, tri := range tris {
			fmt.Printf("triangle: %v\n", tri)	
		}
	} else {
		fmt.Printf("No triangles for %v\n", poly)	
	}
}
//{44 736} {44 848} {88 848} {88 1044} {44 1044} {44 1244} {68 1244} {68 1068} {112 1068} {112 824} {68 824} {68 736}
func TestPiece(t *testing.T) {
	vertices := []Point{
		Point{1, 1},
		Point{1, 6},
		Point{2, 6},
		Point{2, 3},
		Point{4, 3},
		Point{4, 6},
		Point{5, 6},
		Point{5, 2},
		Point{2, 2},
		Point{2, 1},
	}

	poly := new(Polygon)
	for _, v := range vertices {
		poly.AddVertex(v)
	}
	tris, ok := poly.Triangles()
	if ok {
		fmt.Println()
		for _, tri := range tris {
			fmt.Printf("triangle: %v\n", tri)	
		}
	} else {
		fmt.Printf("No triangles for %v\n", poly)	
	}
	
	vertices = []Point{
		Point{44, 736},
		Point{44, 848},
		Point{88, 848},
		Point{88, 1044},
		Point{44, 1044},
		Point{44, 1244},
		Point{68, 1244},
		Point{68, 1068},
		Point{112, 1068},
		Point{112, 824},
		Point{68, 824},
		Point{68, 736},
	}

	poly = new(Polygon)
	for _, v := range vertices {
		poly.AddVertex(v)
	}
	tris, ok = poly.Triangles()
	if ok {
		fmt.Println()
		for _, tri := range tris {
			fmt.Printf("triangle: %v\n", tri)	
		}
	} else {
		fmt.Printf("No triangles for %v\n", poly)	
	}
}