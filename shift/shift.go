package shift

import (
	"local/presentations/geometry/abstract"
	"local/presentations/geometry/point"
)

type shiftedContainer struct {
	c     abstract.Container
	shift func(p abstract.Point) abstract.Point
}

type invertedContainer struct {
	c abstract.Container
}

func (ic invertedContainer) Contains(p abstract.Point) bool {
	return !ic.c.Contains(p)
}

func (sc shiftedContainer) Contains(p abstract.Point) bool {
	return sc.c.Contains(sc.shift(p))
}

func offset(q abstract.Point) func(p abstract.Point) abstract.Point {
	return func(p abstract.Point) abstract.Point {
		return point.Sub(p, q)
	}
}

//Shift a container by an offset.
func Shift(c abstract.Container, y, x float64) abstract.Container {
	return shiftedContainer{c, offset(point.Point(y, x))}
}

func FlipAcrossY(c Container) Container {
	return shiftedContainer{c, point.InvertY}
}

func FlipAcrossX(c Container) Container {
	return shiftedContainer{c, point.InvertX}
}
