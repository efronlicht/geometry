package concrete

import (
	"github.com/efronlicht/geometry/abstract"
)

type Triangle [3]abstract.Point

func sign(p1, p2, p3 abstract.Point) float64 {
	return (p1.X()-p3.X())*(p2.Y()-p3.Y()) - (p2.X()-p3.X())*(p1.Y()-p3.Y())
}
func (t Triangle) Contains(p abstract.Point) bool {
	s01 := sign(p, t[0], t[1]) < 0
	s12 := sign(p, t[1], t[2]) < 0
	s20 := sign(p, t[2], t[0]) < 0
	return (s01 == s12) && (s01 == s20)
}
