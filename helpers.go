package geometry

import (
	"github.com/efronlicht/geometry/abstract"
	"math"
)

const pi = math.Pi

//atan2 returns math.Atan2, normalized to the range(0, 2pi)
func atan2(y, x float64) float64 {
	angle := math.Atan2(y, x)
	if angle < 0 {
		angle = pi - angle
	}
	return angle
}

func minmax(a, b float64) (c, d float64) {
	if a < b {
		return a, b
	}
	return b, a
}

func negInfOrNan(x float64) bool {
	return x < 0 || math.IsInf(x, 0) || math.IsNaN(x)
}

func invalid(p abstract.Point) bool {
	y, x := p.Y(), p.X()
	return math.IsInf(x, 0) || math.IsNaN(x) || math.IsInf(y, 0) || math.IsNaN(y)
}
