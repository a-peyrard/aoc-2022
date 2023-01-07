package collection

/*
	Based on the example provided here: https://pkg.go.dev/container/heap#example__priorityQueue
*/

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item[T any] struct {
	value    T   // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func CreateItem[T any](value T, priority int) *Item[T] {
	return &Item[T]{
		value:    value,
		priority: priority,
	}
}

func (i *Item[T]) GetValue() T {
	return i.value
}

func (i *Item[T]) GetPriority() int {
	return i.priority
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[T any] []*Item[T]

func (pq *PriorityQueue[T]) Len() int { return len(*pq) }

func (pq *PriorityQueue[T]) IsEmpty() bool { return pq.Len() == 0 }

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority, so we use greater than here.
	return (*pq)[i].priority < (*pq)[j].priority
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue[T]) Update(item *Item[T], value T, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
