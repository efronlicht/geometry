package concrete

import (
	"github.com/efronlicht/geometry/point"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	unitSquare = Rectangle{Y0: -1, Y1: 1, X0: -1, X1: 1}
)

func TestRectangle_Contains(t *testing.T) {
	assert.True(t, unitSquare.Contains(point.Point(0.5, 0.5)))
	assert.False(t, unitSquare.Contains(point.Point(1.1, 0)))
}
