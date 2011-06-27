package geom

import "fmt"

var Debug = false

func dbg(format string, args ...interface{}) {
	if Debug {
		fmt.Printf(format+"\n", args...)
	}
}