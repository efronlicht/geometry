package geometry

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

//union combines zero or more Containers
type union []abstract.Container

//Contains returns true if any of the Containers within the mc Contain the abstract.Point
func (u union) Contains(p abstract.Point) bool {
	for _, c := range u {
		if c.Contains(p) {
			return true
		}
	}
	return false
}

type intersection []abstract.Container

func (i intersection) Contains(p abstract.Point) bool {
	if len(i) == 0 {
		return false
	}
	for _, c := range i {
		if !c.Contains(p) {
			return false
		}
	}
	return true
}

type xorContainer [2]abstract.Container

func (xc xorContainer) Contains(p abstract.Point) bool {
	c, d := xc[0].Contains(p), xc[1].Contains(p)
	return (c && !d) || (d && !c)
}

type subContainer [2]abstract.Container

func (sc subContainer) Contains(p abstract.Point) bool {
	return sc[0].Contains(p) && !sc[1].Contains(p)
}
