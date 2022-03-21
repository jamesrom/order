package chansort_test

import (
	"testing"
	"time"

	"github.com/jamesrom/order/chansort"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestSlowConsumer(t *testing.T) {
	t.Parallel()
	const WindowSize = 3 * time.Second
	const BatchSize = 10000

	// fill channel with 10000 messages
	messages := make(chan int, BatchSize)
	defer close(messages)
	for i := 1; i <= BatchSize; i++ {
		messages <- i
	}

	// act
	sortedMessages := chansort.SortSimple(messages, WindowSize)

	// consume slowly
	for i := 1; i <= BatchSize; i++ {
		time.Sleep(time.Millisecond)
		require.Equal(t, i, <-sortedMessages) // assert order
	}
}

func TestEarlyClose(t *testing.T) {
	t.Parallel()
	const WindowSize = 1 * time.Second
	const BatchSize = 50

	// fill channel with 50 messages
	messages := make(chan int, BatchSize)
	for i := 1; i <= BatchSize; i++ {
		messages <- i
	}
	// close channel
	close(messages)

	// act
	sortedMessages := chansort.SortSimple(messages, WindowSize)

	// assert
	for i := 1; i <= BatchSize; i++ {
		time.Sleep(time.Millisecond)
		require.Equal(t, i, <-sortedMessages) // assert order
	}
}
