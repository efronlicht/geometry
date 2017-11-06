package concrete

import "local/presentations/geometry/abstract"

type Rectangle struct {
	y1, y0 float64
	x1, x0 float64
}

func (r Rectangle) Contains(p abstract.Point) bool {
	y, x := p.Y(), p.X()
	return r.y1 > y &&
		r.y0 < y &&
		r.x1 > x &&
		r.x0 < x
}
