package genericqueue

type Interface[T any] interface {
	Push(item *Item[T])
	Update(item *Item[T], value T, priority int)
	Pop() *Item[T]
	Len() int
	Arr() PriorityQueue[T]
}
