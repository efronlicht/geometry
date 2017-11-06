package concrete

import (
	"local/presentations/geometry/abstract"
	"math"
)

type Circle struct {
	abstract.Point
	radius float64
}

func (c Circle) Contains(p abstract.Point) bool {
	r := c.radius
	if r < 0 {
		return false
	}
	x, y := p.X(), p.Y()
	dx, dy := math.Abs(c.X()-x), math.Abs(c.Y()-y)
	return dx < r && dy < r
}

type SemiCircle struct {
	Circle
	start, stop float64
}

func (sc SemiCircle) Contains(p abstract.Point) bool {
	if !sc.Circle.Contains(p) {
		return false
	}

	angle := atan2(p.Y()-sc.Y(), p.X()-sc.X())
	if sc.start <= angle && angle <= sc.stop {
		return true
	}
	return false
}
