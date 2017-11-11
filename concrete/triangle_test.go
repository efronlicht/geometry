package concrete

import (
	"local/presentations/geometry/point"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	a             = point.Point(0, 0)
	b             = point.Point(1, 1)
	c             = point.Point(0, 1)
	rightTriangle = Triangle{a, b, c}
)

func TestTriangle_Contains(t *testing.T) {
	assert.True(t, rightTriangle.Contains(point.Point(0.2, 0.7)))
	assert.False(t, rightTriangle.Contains(point.Point(.6, .2)))
}
