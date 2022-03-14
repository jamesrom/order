package priorityqueue

import (
	hp "container/heap"
	"sync"

	"github.com/jamesrom/order/compare"
)

type PriorityQueue[T any] struct {
	h   heap[T]
	mtx sync.RWMutex
}

type Prioritizable interface {
	Priority() int
}

// NewSimple - ascending order / minheap
func NewSimple[T compare.SimpleOrdered](items ...T) *PriorityQueue[T] {
	return NewWithComparator(compare.LessSimple[T], items...)
}

// NewPriority - highest priority / descending order / maxheap
func NewPriority[T Prioritizable](items ...T) *PriorityQueue[T] {
	less := func(i, j T) bool {
		return i.Priority() > j.Priority()
	}
	return NewWithComparator(less, items...)
}

// NewWithComparator - order defined by a custom less function
func NewWithComparator[T any](fn compare.LessFunc[T], items ...T) *PriorityQueue[T] {
	h := heap[T]{
		elements: items,
		less:     fn,
	}
	hp.Init(&h)
	return &PriorityQueue[T]{h: h}
}

// Pushes x into the priority queue
func (p *PriorityQueue[T]) Push(x T) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	hp.Push(&p.h, x)
}

// Pop returns the element at the head of the queue and removes it from the queue
func (p *PriorityQueue[T]) Pop() T {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	return hp.Pop(&p.h).(T)
}

// Peek keeps the top element in the queue and returns it's value
func (p *PriorityQueue[T]) Peek() T {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	return p.h.elements[0]
}

// Len returns the size of the queue
func (p *PriorityQueue[T]) Len() int {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	return len(p.h.elements)
}
