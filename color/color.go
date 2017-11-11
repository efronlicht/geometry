package color

import "local/presentations/geometry/abstract"

type color string

func (c color) Match(other abstract.Color) bool {
	return c.RGB() == other.RGB()
}

func (c color) RGB() [3]byte {
	switch string(c) {
	case "RED":
		return [3]byte{255, 0, 0}
	case "GREEN":
		return [3]byte{0, 255, 0}
	case "BLUE":
		return [3]byte{0, 0, 255}
	default: //unknown color, so black
		return [3]byte{}
	}
}

const (
	RED   = "RED"
	GREEN = "GREEN"
	BLUE  = "BLUE"
	BLACK = "BLACK"
)

type coloredContainer struct {
	abstract.Container
	color
}

type coloredPoint struct {
	abstract.Point
	color
}

var _ abstract.ColoredContainer = coloredContainer{}
var _ abstract.ColoredPoint = coloredPoint{}

func Container(container abstract.Container, color color) abstract.ColoredContainer {
	return coloredContainer{container, color}
}

func Point(point abstract.Point, color color) abstract.ColoredPoint {
	return coloredPoint{point, color}
}
func (c coloredContainer) ContainsMatching(p coloredPoint) bool {
	return c.Contains(p) && c.Match(p)
}
