package concrete

import (
	"github.com/efronlicht/geometry/abstract"
	"math"
)

type Circle struct {
	abstract.Point
	Radius float64
}

var _ abstract.Container = Circle{}

func (c Circle) Contains(p abstract.Point) bool {
	r := c.Radius
	if r < 0 {
		return false
	}
	x, y := p.X(), p.Y()
	dx, dy := math.Abs(c.X()-x), math.Abs(c.Y()-y)
	return dx < r && dy < r
}

type SemiCircle struct {
	Circle
	Start, Stop float64
}

func (sc SemiCircle) Contains(p abstract.Point) bool {
	if !sc.Circle.Contains(p) {
		return false
	}

	angle := atan2(p.Y()-sc.Y(), p.X()-sc.X())
	if sc.Start <= angle && angle <= sc.Stop {
		return true
	}
	return false
}
