package compare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLess(t *testing.T) {
	t.Parallel()
	assert.True(t, LessSimple(1, 2))
	assert.True(t, LessFloat(1.01, 2.01))
	assert.False(t, LessFloat(float64(1.0000000000000001), float64(1.00000000000000000000001)))
	assert.True(t, LessSimple("a", "z"))
}
