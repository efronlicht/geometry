package combine

import (
	"local/presentations/geometry/abstract"
)

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

type invertedContainer struct {
	c abstract.Container
}

func (ic invertedContainer) Contains(p abstract.Point) bool {
	return !ic.c.Contains(p)
}
