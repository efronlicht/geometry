package concrete

import "github.com/efronlicht/geometry/abstract"

type Rectangle struct {
	Y1, Y0 float64
	X1, X0 float64
}

var _ abstract.Container = Rectangle{}

func (r Rectangle) Contains(p abstract.Point) bool {
	y, x := p.Y(), p.X()
	return r.Y1 > y &&
		r.Y0 < y &&
		r.X1 > x &&
		r.X0 < x
}
