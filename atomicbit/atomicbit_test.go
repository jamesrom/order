package atomicbit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlip(t *testing.T) {
	bit := New(true)
	fmt.Println("hi")
	assert.True(t, bit.Get())
	bit.Flip()
	assert.False(t, bit.Get())
	bit.Flip()
	assert.True(t, bit.Get())
}

func TestGetSet(t *testing.T) {
	bit := New(false)
	assert.False(t, bit.Get())
	bit.Set(true)
	assert.True(t, bit.Get())
}
