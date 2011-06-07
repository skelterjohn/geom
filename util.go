package geom

func minf(x, y float64) (r float64) {
	r = x
	if r > y {
		r = y
	}
	return
}

func maxf(x, y float64) (r float64) {
	r = x
	if r < y {
		r = y
	}
	return
}
