package concrete

import (
	"local/presentations/geometry/point"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	origin     = point.Point(0, 0)
	p20        = point.Point(2, 0)
	unitCircle = Circle{origin, 1}
)

func TestCircle_Contains(t *testing.T) {
	badCircle := Circle{origin, -1}
	assert.True(t, unitCircle.Contains(origin))
	assert.False(t, unitCircle.Contains(p20))
	assert.True(t, unitCircle.Contains(point.Point(-.5, -.5)))
	assert.False(t, badCircle.Contains(origin))

}

func TestSemiCircle_Contains(t *testing.T) {
	topHalf := SemiCircle{unitCircle, 0, pi}
	assert.False(t, topHalf.Contains(point.Point(-.5, -.5)))
	assert.False(t, topHalf.Contains(point.Point(1.1, 1.1)))
	assert.True(t, topHalf.Contains(point.Point(.1, .1)))
}
