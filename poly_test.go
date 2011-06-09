package geom

import (
	"testing"
	"fmt"
)

func aTestPolyTriangularize(t *testing.T) {
	poly := new(Polygon)
	poly.AddVertex(Point{0, 0})
	poly.AddVertex(Point{0, 1})
	poly.AddVertex(Point{1, 1})
	poly.AddVertex(Point{1, 0})
	tris := poly.Triangles()
	for _, tri := range tris {
		fmt.Printf("triangle: %v\n", tri)	
	}
	
	poly = new(Polygon)
	poly.AddVertex(Point{0, 0})
	poly.AddVertex(Point{1, 1})
	poly.AddVertex(Point{2, 0})
	poly.AddVertex(Point{2, 3})
	poly.AddVertex(Point{1, 2})
	poly.AddVertex(Point{0, 3})
	tris = poly.Triangles()
	fmt.Println()
	for _, tri := range tris {
		fmt.Printf("triangle: %v\n", tri)	
	}
}

func TestPiece(t *testing.T) {
	vertices := []Point{
		Point{2408, 2332},
		Point{2408, 2516},
		Point{2432, 2516},
		Point{2432, 2400},
		Point{2496, 2400},
		Point{2496, 2516},
		Point{2520, 2516},
		Point{2520, 2376},
		Point{2432, 2376},
		Point{2432, 2332},
	}

	poly := new(Polygon)
	for _, v := range vertices {
		poly.AddVertex(v)
	}
	tris := poly.Triangles()
	fmt.Println()
	for _, tri := range tris {
		fmt.Printf("triangle: %v\n", tri)	
	}
}