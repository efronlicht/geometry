package geometry

import (
	"github.com/efronlicht/geometry/point"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	origin        = point.Point(0, 0)
	unitCircle, _ = Circle(origin, 1)

	r0201, _ = Rectangle(point.Point(2, 0), point.Point(0, 1))

	inCircle    = point.Point(-0.5, -0.5)
	inRectangle = point.Point(1.5, 0.2)
	notInEither = point.Point(-1.2, -5)
	inBoth      = point.Point(0.5, 0.5)
)

func Test_union_Contains(t *testing.T) {
	union := Union(unitCircle, r0201)

	assert.True(t, union.Contains(inCircle))
	assert.True(t, union.Contains(inRectangle))
	assert.True(t, union.Contains(inBoth))
	assert.False(t, union.Contains(notInEither))
}

func Test_intersection_Contains(t *testing.T) {
	intersection := Intersection(unitCircle, r0201)
	assert.False(t, intersection.Contains(inCircle))
	assert.False(t, intersection.Contains(inRectangle))
	assert.True(t, intersection.Contains(inBoth))
	assert.False(t, intersection.Contains(notInEither))
}
