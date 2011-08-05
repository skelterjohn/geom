// Copyright 2009 The geom Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geom

import "fmt"

var Debug = false

func dbg(format string, args ...interface{}) {
	if Debug {
		fmt.Printf(format+"\n", args...)
	}
}