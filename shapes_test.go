package geometry

import (
	"github.com/efronlicht/geometry/point"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDonut(t *testing.T) {
	_, ok := Donut(origin, -1, 3)
	assert.False(t, ok)
	_, ok = Donut(point.Point(math.Inf(-1), math.NaN()), 0.1, 1)
	assert.False(t, ok)

	donut, _ := Donut(origin, 1, 2)
	assert.True(t, donut.Contains(point.Point(1.2, 1.3)))
	assert.False(t, donut.Contains(origin))
}
