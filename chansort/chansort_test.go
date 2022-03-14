package chansort_test

import (
	"testing"
	"time"

	"github.com/jamesrom/order/chansort"
	"github.com/stretchr/testify/assert"
)

func TestAscending(t *testing.T) {
	t.Parallel()
	const WindowSize = 3 * time.Second
	const MinValue = -10
	const MaxValue = 10

	// arrange
	messages := make(chan int, 10)
	defer close(messages)
	messages <- 1
	messages <- 2
	messages <- 3
	messages <- 4
	messages <- 5

	// act
	// push after two seconds (inside window)
	time.AfterFunc(time.Second*2, func() {
		messages <- MinValue
	})
	sortedMessages := chansort.SortSimple(messages, WindowSize)

	// assert
	assert.Equal(t, MinValue, <-sortedMessages)
	assert.Equal(t, 1, <-sortedMessages)
	assert.Equal(t, 2, <-sortedMessages)
	assert.Equal(t, 3, <-sortedMessages)
	assert.Equal(t, 4, <-sortedMessages)
	assert.Equal(t, 5, <-sortedMessages)
}

func TestDescending(t *testing.T) {
	t.Parallel()
	const WindowSize = 3 * time.Second
	const MinValue = -10
	const MaxValue = 10

	// arrange
	messages := make(chan int, 10)
	defer close(messages)
	messages <- 1
	messages <- 2
	messages <- 3
	messages <- 4
	messages <- 5

	// act
	// push after two seconds (inside window)
	time.AfterFunc(time.Second*2, func() {
		messages <- MaxValue
	})
	sortedMessages := chansort.SortSimple(messages, WindowSize)

	// assert
	assert.Equal(t, 1, <-sortedMessages)
	assert.Equal(t, 2, <-sortedMessages)
	assert.Equal(t, 3, <-sortedMessages)
	assert.Equal(t, 4, <-sortedMessages)
	assert.Equal(t, 5, <-sortedMessages)
	assert.Equal(t, MaxValue, <-sortedMessages)
}

func TestWindowSlide(t *testing.T) {
	t.Parallel()
	const WindowSize = 3 * time.Second
	const MinValue = -10
	const MaxValue = 10

	// arrange
	messages := make(chan int, 10)
	defer close(messages)
	messages <- 1
	messages <- 2
	messages <- 3
	messages <- 4
	messages <- 5

	// act
	// push after 5 seconds, outside window
	time.AfterFunc(time.Second*5, func() {
		messages <- MinValue
	})
	sortedMessages := chansort.SortSimple(messages, WindowSize)

	// assert
	assert.Equal(t, 1, <-sortedMessages)
	assert.Equal(t, 2, <-sortedMessages)
	assert.Equal(t, 3, <-sortedMessages)
	assert.Equal(t, 4, <-sortedMessages)
	assert.Equal(t, 5, <-sortedMessages)
	assert.Equal(t, MinValue, <-sortedMessages)
}
