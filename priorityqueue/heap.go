package priorityqueue

// implements heap.Interface
type heap[T any] struct {
	elements []T
	less     func(T, T) bool
}

func (h *heap[T]) Len() int {
	return len(h.elements)
}

func (h *heap[T]) Push(x any) {
	h.elements = append(h.elements, x.(T))
}

func (h *heap[T]) Less(i, j int) bool {
	return h.less(h.elements[i], h.elements[j])
}

// Pop is called after the first element is swapped with the last
// so return the last element and resize the slice
func (h *heap[T]) Pop() any {
	last := len(h.elements) - 1
	element := h.elements[last]
	h.elements = h.elements[:last]
	return element
}

func (h *heap[T]) Swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}
