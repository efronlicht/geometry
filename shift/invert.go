package shift

import "local/presentations/geometry/abstract"

func Invert(c abstract.Container) abstract.Container {
	return invertedContainer{c}
}
