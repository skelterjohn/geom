package qtree

/*
A number of tests to test qtree's ability to analyze the differences between
collections of shapes
*/

import (
	"github.com/skelterjohn/geom"
)

func col1() (items []Item) {
	items = append(items, Item(&geom.Rect{geom.Point{0, 0}, geom.Point{1, 1}}))
	items = append(items, Item(&geom.Rect{geom.Point{0, 2}, geom.Point{1, 3}}))
	items = append(items, Item(&geom.Rect{geom.Point{2, 2}, geom.Point{3, 3}}))
	items = append(items, Item(&geom.Rect{geom.Point{2, 0}, geom.Point{3, 1}}))
	return
}

func col2() (items []Item) {
	items = append(items, Item(&geom.Rect{geom.Point{10, 0}, geom.Point{11, 1}}))
	items = append(items, Item(&geom.Rect{geom.Point{10, 2}, geom.Point{11, 3}}))
	items = append(items, Item(&geom.Rect{geom.Point{12, 2}, geom.Point{13, 3}}))
	items = append(items, Item(&geom.Rect{geom.Point{12, 0}, geom.Point{13, 1}}))
	return
}

func col3() (items []Item) {
	items = append(items, Item(&geom.Rect{geom.Point{0, 10}, geom.Point{1, 11}}))
	items = append(items, Item(&geom.Rect{geom.Point{0, 12}, geom.Point{1, 13}}))
	items = append(items, Item(&geom.Rect{geom.Point{2, 12}, geom.Point{3, 13}}))
	items = append(items, Item(&geom.Rect{geom.Point{2, 10}, geom.Point{3, 11}}))
	return
}
