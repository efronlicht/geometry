package density

import "local/presentations/geometry/abstract"

//SetDensity uniformly sets the density of a container to n.
func Set(c abstract.Container, n int) abstract.DenseContainer {
	return denseContainer{c, n}
}

//AdjustDensity maps the function f across the Density of the DenseContainer
func Adjust(dc abstract.DenseContainer, f func(int) int) abstract.DenseContainer {
	return adjuster{dc, f}
}

//Neg returns a abstract.DenseContainer whose density inverts the density of it's argument.
func Neg(c abstract.DenseContainer) abstract.DenseContainer {
	return adjuster{c, func(a int) int { return -a }}
}

type denseContainer struct {
	abstract.Container
	density int
}

type denseUnion []abstract.DenseContainer

func (du denseUnion) Contains(p abstract.Point) bool {
	for _, c := range du {
		if c.Contains(p) {
			return true
		}
	}
	return false
}

func (du denseUnion) Density(p abstract.Point) int {
	var density int
	for _, c := range du {
		density += c.Density(p)
	}
	return density
}

type denseIntersection []abstract.DenseContainer

func (di denseIntersection) Contains(p abstract.Point) bool {
	for _, c := range di {
		if !c.Contains(p) {
			return false
		}
	}
	return len(di) > 0
}

func (di denseIntersection) Density(p abstract.Point) int {
	var density int
	for _, c := range di {
		if !c.Contains(p) {
			return 0
		}
		density += c.Density(p)
	}

	return density
}
func (d denseContainer) Density(p abstract.Point) int {
	if d.Contains(p) {
		return d.density
	}
	return 0
}

type adjuster struct {
	abstract.DenseContainer
	f func(int) int
}

func (a adjuster) Density(p abstract.Point) int {
	return a.f(a.DenseContainer.Density(p))

}
