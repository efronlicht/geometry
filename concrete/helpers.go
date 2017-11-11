package concrete

import (
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
