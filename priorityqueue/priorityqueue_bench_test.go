package priorityqueue_test

import (
	"testing"

	"github.com/jamesrom/order/priorityqueue"
)

// Push N elements onto queue of size M. The pushed element has the highest priority
// reflecting the worst-case 'swim up' scenario.
func pushMax(b *testing.B, M int) {
	q := setupPriorityQueue(b, M)
	for n := 0; n < b.N; n++ {
		q.Push(M + n)
	}
}
func BenchmarkPushMaxOntoSize10(b *testing.B)       { pushMax(b, 10) }
func BenchmarkPushMaxOntoSize100(b *testing.B)      { pushMax(b, 100) }
func BenchmarkPushMaxOntoSize1000(b *testing.B)     { pushMax(b, 1000) }
func BenchmarkPushMaxOntoSize10000(b *testing.B)    { pushMax(b, 10000) }
func BenchmarkPushMaxOntoSize100000(b *testing.B)   { pushMax(b, 100000) }
func BenchmarkPushMaxOntoSize1000000(b *testing.B)  { pushMax(b, 1000000) }
func BenchmarkPushMaxOntoSize10000000(b *testing.B) { pushMax(b, 10000000) }

// Push N elements onto queue of size M. The pushed element has the lowest priority
// reflecting the best-case 'swim up' scenario.
func pushMin(b *testing.B, M int) {
	q := setupPriorityQueue(b, M)
	for n := 0; n < b.N; n++ {
		q.Push(0 - n)
	}
}
func BenchmarkPushMinOntoSize10(b *testing.B)       { pushMin(b, 10) }
func BenchmarkPushMinOntoSize100(b *testing.B)      { pushMin(b, 100) }
func BenchmarkPushMinOntoSize1000(b *testing.B)     { pushMin(b, 1000) }
func BenchmarkPushMinOntoSize10000(b *testing.B)    { pushMin(b, 10000) }
func BenchmarkPushMinOntoSize100000(b *testing.B)   { pushMin(b, 100000) }
func BenchmarkPushMinOntoSize1000000(b *testing.B)  { pushMin(b, 1000000) }
func BenchmarkPushMinOntoSize10000000(b *testing.B) { pushMin(b, 10000000) }

// Pop N elements from queue of size M+N.
func pop(b *testing.B, M int) {
	q := setupPriorityQueue(b, M+b.N)
	for n := 0; n < b.N; n++ {
		q.Pop()
	}
}

func BenchmarkPopFromSize10(b *testing.B)       { pop(b, 10) }
func BenchmarkPopFromSize100(b *testing.B)      { pop(b, 100) }
func BenchmarkPopFromSize1000(b *testing.B)     { pop(b, 1000) }
func BenchmarkPopFromSize10000(b *testing.B)    { pop(b, 10000) }
func BenchmarkPopFromSize100000(b *testing.B)   { pop(b, 100000) }
func BenchmarkPopFromSize1000000(b *testing.B)  { pop(b, 1000000) }
func BenchmarkPopFromSize10000000(b *testing.B) { pop(b, 10000000) }

// Peek N times on queue of size M.
func peek(b *testing.B, M int) {
	q := setupPriorityQueue(b, M)
	for n := 0; n < b.N; n++ {
		q.Peek()
	}
}

func BenchmarkPeekFromSize10(b *testing.B)       { peek(b, 10) }
func BenchmarkPeekFromSize100(b *testing.B)      { peek(b, 100) }
func BenchmarkPeekFromSize1000(b *testing.B)     { peek(b, 1000) }
func BenchmarkPeekFromSize10000(b *testing.B)    { peek(b, 10000) }
func BenchmarkPeekFromSize100000(b *testing.B)   { peek(b, 100000) }
func BenchmarkPeekFromSize1000000(b *testing.B)  { peek(b, 1000000) }
func BenchmarkPeekFromSize10000000(b *testing.B) { peek(b, 10000000) }

// setup used for some tests, should be called at the beginning of a benchmark test
// as ResetTimer is called when this function exits
func setupPriorityQueue(b *testing.B, size int) *priorityqueue.PriorityQueue[int] {
	b.Helper()
	defer b.ResetTimer()

	q := priorityqueue.NewWithComparator(func(a, b int) bool { return a > b })
	for i := 0; i < size; i++ {
		q.Push(i)
	}
	return q
}
