//concrete.go contains structs that represent concrete geometry

package shapes

import (
	"local/presentations/geometry/concrete"
	"math"
)

const pi = math.Pi

func Square(p Point, sidelen float64) (Container, bool) {
	if negInfOrNan(sidelen) || p.invalid() {
		return Container(nil), false
	}
	w := sidelen / 2
	return rectangle{
		y1: p.y + w, y0: p.y - w,
		x1: p.x + w, x0: p.x - w,
	}, true
}

func Donut(p Point, r, R float64) (Container, bool) {

	r, R = minmax(r, R)
	inner := concrete.Circle{p, r}
	outer := concrete.Circle{p, R}

	if negInfOrNan(r) || negInfOrNan(R) || p.invalid() {
		return Container(nil), false
	}
	return Sub(outer, inner), true
}

func Circle(p Point, r float64) (container Container, ok bool) {
	if negInfOrNan(r) || p.invalid() {
		return
	}
	return concrete.Circle{Point: p, radius: r}, true
}

func Triangle(a, b, c Point) (container Container, ok bool) {
	if a.invalid() || b.invalid() || c.invalid() {
		return
	} else if a.x == b.x && b.x == c.x {
		return
	} else if a.y == b.y && b.y == c.y {
		return
	}
	return concrete.Triangle{a, b, c}, true
}

//Rectangle creates a rectangle from it's corners
func Rectangle(a, b Point) (Container, bool) {
	y0, y1 := minmax(a.y, b.y)
	x0, x1 := minmax(a.x, b.x)
	if y0 == y1 || x0 == x1 {
		return Container(nil), false
	}
	return concrete.Rectangle{y1: y1, y0: y0, x1: x1, x0: x0}, true
}

func SemiCircle(p Point, r, start, stop float64) (container Container, ok bool) {
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
	return concrete.SemiCircle{Circle: circle, start: start, stop: stop}, true
}

func NewPoint(y, x float64) Point {
	return concrete.Point{y: y, x: x}
}
