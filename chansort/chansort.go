package chansort

import (
	"sync"
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
	var m sync.Mutex // mutex ensures FIFO send to out chan
	go func() {
		for el := range in {
			q.Push(el)
			time.AfterFunc(window, func() {
				m.Lock()
				defer m.Unlock()
				out <- q.Pop()
			})
		}
		// if we get here that means the channel has closed, so close the output
		// channel too.
		close(out)
	}()
	return out
}
