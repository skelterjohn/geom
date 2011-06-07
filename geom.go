//target:github.com/skelterjohn/geom
package geom

type Bounded interface {
	Bounds() (bounds *Rect)
}
