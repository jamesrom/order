package chansort

import (
	"time"

	"github.com/jamesrom/order/atomicbit"
	"github.com/jamesrom/order/compare"
	"github.com/jamesrom/order/priorityqueue"
)

// SortSimple sorts channel messages in ascending order. Messages received
// inside the sliding-window buffer defined by _window_ are sent to the
// output channel in ascending order. That is to say: a message received
// at time _Z_  from the output channel is guaranteed to be the smallest
// message since _Z − window_.
func SortSimple[T compare.SimpleOrdered](in <-chan T, window time.Duration) <-chan T {
	return SortWithComparator(in, window, compare.LessSimple[T])
}

// SortWithComparator sorts channel messages in the order defined by the given
// comparator function. Messages received inside the sliding-window buffer
// defined by _window_ are sent to the output channel in order.
// That is to say: a message received at time _Z_ from the output channel is
// guaranteed to be the smallest message since _Z − window_.
func SortWithComparator[T any](in <-chan T, window time.Duration, fn compare.LessFunc[T]) <-chan T {
	q := priorityqueue.NewWithComparator(fn)
	out := make(chan T)
	closing := atomicbit.New(false)

	// This goroutine waits for a signal that the next element is ready to pop
	// and sends to the out channel.
	popsig := make(chan any)
	go func() {
		// If this defer statement runs, it means the signal channel has closed
		// and so we can safely close the output channel.
		defer close(out)

		for range popsig {
			out <- q.Pop()
			if closing.Get() && q.Len() == 0 {
				// The last instance of a thing takes the class with it.
				// Turns out the light and is gone. -CM
				close(popsig)
			}
		}
	}()

	go func() {
		for el := range in {
			q.Push(el)
			// We shouldn't pop from the AfterFunc directly as a slow consumer
			// will cause a backlog of goroutines that are blocked. And since
			// there's no guarantee which goroutine will be scheduled to send to
			// the receiving channel first, this would lead to unordered values
			// being sent to the out channel.
			//
			// Instead, send a signal after the window duration that the next
			// element is ready to be popped.
			time.AfterFunc(window, func() { popsig <- nil })
		}
		closing.Set(true)
	}()
	return out
}
