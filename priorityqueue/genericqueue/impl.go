package genericqueue

import "container/heap"

// An Item is something we manage in a priority queue.
type Item[T any] struct {
	Value    T
	Priority int
	id       int `json:"-"` // used to fix heap after update in an Item
}

func (i *Item[T]) Pos() int     { return i.id }
func (i *Item[T]) SetId(id int) { i.id = id }

type PriorityQueue[T any] []*Item[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool { return pq[i].Priority > pq[j].Priority }

func (pq PriorityQueue[T]) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.id = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	item.id = -1   // for safety
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue[T]) Update(item *Item[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.id)
}
