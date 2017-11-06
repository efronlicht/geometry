package combine

import "local/presentations/geometry/abstract"

//Union zero or more containers into a single container
func Union(containers ...abstract.Container) abstract.Container {
	return union(containers)
}

func Intersection(containers ...abstract.Container) abstract.Container {
	return intersection(containers)
}

//XOR two containers.
func XOR(c, d abstract.Container) abstract.Container {
	return xorContainer{c, d}
}

//Sub returns a container which contains a Point p if c contains the Point and d does not.
func Sub(c, d abstract.Container) abstract.Container {
	return subContainer{c, d}
}

//Invert a container, so that Invert(c).Contains(p) != c.Contains(p)
//Note: does NOT preserve density.
func Invert(c abstract.Container) abstract.Container {
	return invertedContainer{c}
}
