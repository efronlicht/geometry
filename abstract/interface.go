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

type aDenseContainer struct {
	DenseContainer
}

func (adc aDenseContainer) DoubleDensity(p Point) int {
	return 2 * adc.Density(p)
}

type denseContainer2 interface {
	Contains(Point) bool
	Density(Point) int
}

func satisfiesInterface(a DenseContainer) bool {
	return true
}

type universe struct{}

func (u universe) Contains(_ Point) bool {
	return true
}

func (u universe) Density(_ Point) int {
	return 500
}

var (
	u universe
	_ Container       = u
	_ denseContainer2 = u
	_ DenseContainer  = u

	someDenseContainer    denseContainer2
	anotherDenseContainer DenseContainer
	_                     = satisfiesInterface(someDenseContainer)
	_                     = satisfiesInterface(anotherDenseContainer)
)
