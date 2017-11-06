package point

import "local/presentations/geometry/abstract"

type point struct {
	y, x float64
}

func (p point) Y() float64 {
	return p.y
}
func (p point) X() float64 {
	return p.x
}

func Point(y, x float64) abstract.Point {
	return point{y, x}
}
func Add(p, q abstract.Point) abstract.Point {
	return point{
		y: p.Y() + q.Y(),
		x: p.X() + q.X()}
}

func Invert(p abstract.Point) abstract.Point {
	return point{y: -p.Y(), x: -p.X()}
}
func Sub(p, q abstract.Point) abstract.Point {
	return point{y: p.Y() - q.Y(), x: p.X() - q.X()}
}

func InvertY(p abstract.Point) abstract.Point {
	return point{y: -p.Y(), x: p.X()}
}

func InvertX(p abstract.Point) abstract.Point {
	return point{y: p.Y(), x: p.X()}
}
