package geometry

import "github.com/efronlicht/geometry/abstract"

func Invert(c abstract.Container) abstract.Container {
	return invertedContainer{c}
}

type invertedContainer struct {
	c abstract.Container
}

func (ic invertedContainer) Contains(p abstract.Point) bool {
	return !ic.c.Contains(p)
}
