package color

import (
	"local/presentations/geometry"
	"local/presentations/geometry/point"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	origin        = point.Point(0, 0)
	unitCircle, _ = geometry.Circle(origin, 0)
	farAway       = point.Point(22, 34)
)

func TestColoredContainer_ContainsMatching(t *testing.T) {
	blueOrigin := coloredPoint{origin, BLUE}
	redOrigin := coloredPoint{origin, RED}
	redFarAway := coloredPoint{farAway, RED}
	redCircle := coloredContainer{unitCircle, RED}
	assert.False(t, redCircle.ContainsMatching(blueOrigin))
	assert.True(t, redCircle.Contains(blueOrigin))
	assert.True(t, redCircle.ContainsMatching(redOrigin))
	assert.False(t, redCircle.Contains(redFarAway))
}
