package geometry

import (
	"github.com/efronlicht/geometry/point"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_invertedContainer_Contains(t *testing.T) {
	invertedUnitCircle := Invert(unitCircle)
	assert.True(t, invertedUnitCircle.Contains(point.Point(math.Inf(1), math.Inf(1))))
	assert.False(t, invertedUnitCircle.Contains(origin))
}
