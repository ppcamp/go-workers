// This example demonstrates a priority queue built using the heap interface.
package genericqueue

import (
	"container/heap"
)

type pqHeap[T any] struct {
	h PriorityQueue[T]
}

func New[T any](items ...*Item[T]) Interface[T] {
	h := make(PriorityQueue[T], 0)

	for it, item := range items {
		item.SetId(it)
		h = append(h, item)
	}

	l := &pqHeap[T]{h}
	heap.Init(&l.h)
	return l
}

func (s *pqHeap[T]) Push(item *Item[T]) {
	heap.Push(&s.h, item)
}

func (s *pqHeap[T]) Update(item *Item[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(&s.h, item.id)
}

func (s *pqHeap[T]) Pop() *Item[T] {
	item := heap.Pop(&s.h).(*Item[T])
	return item
}

func (s *pqHeap[T]) Len() int { return len(s.h) }

func (s *pqHeap[T]) Arr() PriorityQueue[T] { return s.h }
