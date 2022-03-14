package priorityqueue_test

import (
	"testing"

	"github.com/jamesrom/order/priorityqueue"
	"github.com/stretchr/testify/assert"
)

type x string

func (s x) Priority() int {
	return len(s)
}

func TestNewPriority(t *testing.T) {
	t.Parallel()
	q := priorityqueue.NewPriority(
		x("this"),
		x("is"),
		x("the"),
		x("longest"),
	)
	assert.EqualValues(t, "longest", q.Pop())
	assert.EqualValues(t, "this", q.Pop())
	assert.EqualValues(t, "the", q.Pop())
	assert.EqualValues(t, "is", q.Pop())
}

func TestNewSimple(t *testing.T) {
	t.Parallel()
	q := priorityqueue.NewSimple(1, 2, 3)
	assert.Equal(t, 1, q.Pop())
}

func TestNewWithComparator(t *testing.T) {
	t.Parallel()
	q := priorityqueue.NewWithComparator(
		func(a, b int) bool { return a > b },
		1, 2, 3,
	)
	assert.Equal(t, 3, q.Pop())
	assert.Equal(t, 2, q.Pop())
	assert.Equal(t, 1, q.Pop())
}

func TestPush(t *testing.T) {
	t.Parallel()
	q := priorityqueue.NewWithComparator(
		func(a, b float64) bool { return a > b },
		1.1, 0.2, 300,
	)
	q.Push(4.1235489e5)
	q.Push(50.314)
	assert.Equal(t, 412354.89, q.Pop())
	assert.Equal(t, 300.0, q.Pop())
	assert.Equal(t, 50.314, q.Pop())
	assert.Equal(t, 1.1, q.Pop())
	assert.Equal(t, 0.2, q.Pop())
}

func TestPeek(t *testing.T) {
	t.Parallel()
	q := priorityqueue.NewPriority(
		x("this"),
		x("is"),
		x("the"),
		x("longest"),
	)
	assert.EqualValues(t, "longest", q.Peek())
	assert.EqualValues(t, "longest", q.Pop())
	assert.EqualValues(t, "this", q.Peek())
	assert.EqualValues(t, "this", q.Pop())
	assert.EqualValues(t, "the", q.Peek())
	assert.EqualValues(t, "the", q.Peek())
	assert.EqualValues(t, "the", q.Peek())
	assert.EqualValues(t, "the", q.Pop())
	assert.EqualValues(t, "is", q.Pop())
}

func TestUnderflow(t *testing.T) {
	t.Parallel()
	q := priorityqueue.NewSimple(1, 2, 3)
	assert.Equal(t, 1, q.Pop())
	assert.Equal(t, 2, q.Pop())
	assert.Equal(t, 3, q.Pop())
	assert.Panics(t, func() {
		q.Pop()
	})
}

func TestLen(t *testing.T) {
	t.Parallel()
	q := priorityqueue.NewSimple(1, 2, 3)
	assert.Equal(t, 3, q.Len())
}
