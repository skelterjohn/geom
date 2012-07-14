package geom

import (
	"fmt"
	"testing"
)

type LineIntersectionTest struct {
	l1, l2 Line
	p      Coord
}

var itests = []LineIntersectionTest{
	{
		Line{
			Intersection: Coord{0, 0},
			Normal:       Coord{0, 1},
		},
		Line{
			Intersection: Coord{1, 1},
			Normal:       Coord{1, 0},
		},
		Coord{0, 1},
	},
}

func TestIntersect(t *testing.T) {
	for _, test := range itests {
		p := LineIntersection(test.l1, test.l2)
		if p != test.p {
			t.Error(fmt.Sprintf("should be %v, was %v", test.p, p))
		}
	}
}
