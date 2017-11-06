package concrete

import "local/presentations/geometry/abstract"

type Triangle [3]Point

func sign(p1, p2, p3 Point) float64 {
	return (p1.x-p3.x)*(p2.y-p3.y) - (p2.x-p3.x)*(p1.y-p3.y)
}
func (t Triangle) Contains(p abstract.Point) bool {

	var q = Point{p.Y(), p.X()}

	s01 := sign(q, t[0], t[1])
	s12 := sign(q, t[1], t[2])
	s20 := sign(q, t[2], t[0])
	return (s01 == s12) && (s01 == s20)
}

type Point struct {
	y, x float64
}
