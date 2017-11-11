package geometry

import (
	"github.com/efronlicht/geometry/abstract"

	"github.com/efronlicht/geometry/concrete"
	"math"
)

func Donut(p abstract.Point, r, R float64) (abstract.Container, bool) {

	r, R = minmax(r, R)
	inner := concrete.Circle{p, r}
	outer := concrete.Circle{p, R}

	if negInfOrNan(r) || negInfOrNan(R) || invalid(p) {
		return nil, false
	}
	return Sub(outer, inner), true
}

func Circle(p abstract.Point, r float64) (container abstract.Container, ok bool) {
	if negInfOrNan(r) || invalid(p) {
		return
	}
	return concrete.Circle{p, r}, true
}

func Triangle(a, b, c abstract.Point) (container abstract.Container, ok bool) {
	if invalid(a) || invalid(b) || invalid(c) {
		return
	} else if a.X() == b.X() && b.X() == c.X() {
		return
	} else if a.Y() == b.Y() && b.Y() == c.Y() {
		return
	}
	return concrete.Triangle{a, b, c}, true
}

//Rectangle creates a rectangle from it's corners
func Rectangle(a, b abstract.Point) (abstract.Container, bool) {
	y0, y1 := minmax(a.Y(), b.Y())
	x0, x1 := minmax(a.X(), b.X())
	if y0 == y1 || x0 == x1 {
		return nil, false
	}
	return concrete.Rectangle{Y1: y1, Y0: y0, X1: x1, X0: x0}, true
}

func SemiCircle(p abstract.Point, r, start, stop float64) (container abstract.Container, ok bool) {
	circle, ok := Circle(p, r)
	if !ok {
		return
	} else if math.IsNaN(start) || math.IsInf(start, 0) {
		return
	} else if math.IsNaN(stop) || math.IsInf(stop, 0) {
		return
	}
	start, stop = math.Mod(start, 2*pi), math.Mod(stop, 2*pi)
	if start > stop {
		return
	}
	return concrete.SemiCircle{circle.(concrete.Circle), start, stop}, true
}
