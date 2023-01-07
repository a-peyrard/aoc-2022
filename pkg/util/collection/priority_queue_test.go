package collection

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	t.Run("it should work as expected", func(t *testing.T) {
		// Some items and their priorities.
		items := map[string]int{
			"banana": 3, "apple": 2, "pear": 4,
		}

		// Create a priority queue, put the items in it, and
		// establish the priority queue (heap) invariants.
		pq := make(PriorityQueue[string], len(items))
		i := 0
		for value, priority := range items {
			pq[i] = &Item[string]{
				value:    value,
				priority: priority,
				index:    i,
			}
			i++
		}
		heap.Init(&pq)

		// Insert a new item and then modify its priority.
		item := &Item[string]{
			value:    "orange",
			priority: 1,
		}
		heap.Push(&pq, item)
		pq.Update(item, item.value, 5)

		// Take the items out; they arrive in decreasing priority order.
		item = heap.Pop(&pq).(*Item[string])
		assert.Equal(t, 2, item.priority)
		assert.Equal(t, "apple", item.value)

		item = heap.Pop(&pq).(*Item[string])
		assert.Equal(t, 3, item.priority)
		assert.Equal(t, "banana", item.value)

		item = heap.Pop(&pq).(*Item[string])
		assert.Equal(t, 4, item.priority)
		assert.Equal(t, "pear", item.value)

		item = heap.Pop(&pq).(*Item[string])
		assert.Equal(t, 5, item.priority)
		assert.Equal(t, "orange", item.value)

		assert.True(t, pq.IsEmpty())
	})
}
