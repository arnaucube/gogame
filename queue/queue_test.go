package queue

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	// Some items and their priorities.
	items := map[string]int64{
		"id0": 2,
		"id1": 5,
		"id2": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	assert.Equal(t, pq.Look().value, "id0")

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "id3",
		priority: 3,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 1)
	assert.Equal(t, pq.Look().value, "id3")

	// Take the items out; they arrive in increasing priority order.
	for pq.Len() > 0 {
		_ = heap.Pop(&pq).(*Item)
		// fmt.Println("priority", item.priority, ":", item.value)
	}
	assert.Nil(t, pq.Look())
}
