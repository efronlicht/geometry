package point

import (
	"local/presentations/geometry/abstract"
	"math"
)

type point struct {
	y, x float64
}

type polar struct {
	r, theta float64
}

func (p polar) Y() float64 {
	return p.r * math.Sin(p.theta)
}

func (p polar) X() float64 {
	return p.r * math.Cos(p.theta)
}

func (p point) Y() float64 {
	return p.y
}
func (p point) X() float64 {
	return p.x
}

func toPolar(p abstract.Point) polar {
	if p, ok := p.(polar); ok {
		return p
	}
	y, x := p.Y(), p.X()
	r := math.Sqrt(y*y + x*x)
	theta := atan2(y, x)
	return polar{r, theta}
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

func Rotate(p abstract.Point, theta float64) (abstract.Point, bool) {
	if negInfOrNan(theta) {
		return nil, false
	}
	polar := toPolar(p)
	polar.theta = math.Mod(theta+polar.theta, 2*math.Pi)
	return polar, true
}

func InvertY(p abstract.Point) abstract.Point {
	return point{y: -p.Y(), x: p.X()}
}

func InvertX(p abstract.Point) abstract.Point {
	return point{y: p.Y(), x: p.X()}
}

//atan2 returns math.Atan2, normalized to the range(0, 2pi)
func atan2(y, x float64) float64 {
	angle := math.Atan2(y, x)
	if angle < 0 {
		angle = math.Pi - angle
	}
	return angle
}

func negInfOrNan(x float64) bool {
	return x < 0 || math.IsInf(x, 0) || math.IsNaN(x)
}
