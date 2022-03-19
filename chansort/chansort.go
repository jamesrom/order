package chansort

import (
	"time"

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

	// This goroutine listens for signals that the next element is ready to pop
	// and sends to the out chan.
	popsig := make(chan any)
	go func() {
		for range popsig {
			out <- q.Pop()
		}
	}()

	go func() {
		for el := range in {
			q.Push(el)

			// We shouldn't pop from the AfterFunc directly as a slow consumer
			// will cause a backlog of goroutines that are blocked. And since
			// there's no guarantee which goroutine will be scheduled to send to
			// the receiving channel first, this would lead to unordered values
			// being sent to out chan.
			//
			// Instead, send a signal after the window duration that the next
			// element is ready to be popped.
			time.AfterFunc(window, func() { popsig <- nil })
		}
		// if we get here that means the input channel has closed, so close the
		// output channel too.
		close(out)
		close(popsig)
	}()
	return out
}
