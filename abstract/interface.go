package abstract

//Container represents a 2d shape that can contain a point.
type Container interface {
	Contains(Point) bool
}

type Point interface {
	Y() float64
	X() float64
}

//DenseContainer represents a 2d shape that can contain a point, with a "density" of containment. This density can be negative.
type DenseContainer interface {
	Container
	Density(Point) int
}
