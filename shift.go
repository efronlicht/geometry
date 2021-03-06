package geometry

import (
	"github.com/efronlicht/geometry/abstract"
	"github.com/efronlicht/geometry/point"
)

//Shift a container by an offset.
func Shift(c abstract.Container, y, x float64) abstract.Container {
	return shiftedContainer{c, offset(point.Point(y, x))}
}

func FlipAcrossY(c abstract.Container) abstract.Container {
	return shiftedContainer{c, point.InvertY}
}

func FlipAcrossX(c abstract.Container) abstract.Container {
	return shiftedContainer{c, point.InvertX}
}

type shiftedContainer struct {
	c     abstract.Container
	shift func(p abstract.Point) abstract.Point
}

func (sc shiftedContainer) Contains(p abstract.Point) bool {
	return sc.c.Contains(sc.shift(p))
}

func offset(q abstract.Point) func(p abstract.Point) abstract.Point {
	return func(p abstract.Point) abstract.Point {
		return point.Sub(p, q)
	}
}
