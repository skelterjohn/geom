// Copyright 2009 The geom Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//target:github.com/skelterjohn/geom
package geom

type Bounded interface {
	Bounds() (bounds *Rect)
}

type Transformable interface {
	Translate(offset Point)
	Rotate(rad float64)
	Scale(xfactor, yfactor float64)
}