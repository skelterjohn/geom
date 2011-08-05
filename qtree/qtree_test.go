package qtree

import (
	"testing"
	"fmt"
	"github.com/skelterjohn/geom"
)

func TestInsertCollect(t *testing.T) {
	cfg := ConfigDefault()
	qt := New(cfg, &geom.Rect{geom.Point{0, 0}, geom.Point{100, 100}})
	
	r := &geom.Rect{geom.Point{20, 20}, geom.Point{40, 40}}
	qt.Insert(r)

	collection := make(map[Item]bool)
	qt.CollectIntersect(r, collection)
	
	fmt.Printf("%v\n", collection)
}