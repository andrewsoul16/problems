package arrays

import (
	"container/heap"
	"sort"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	var (
		m = make(map[int]*Item)
		h = &PriorityQueue{}
	)

	for _, v := range nums {
		item, ok := m[v]
		if !ok {
			item := &Item{
				value:    v,
				priority: 1,
			}
			m[v] = item
			*h = append(*h, item)
			continue
		}
		item.priority++
	}

	sort.Sort(h)

	var res = make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(h).(*Item).value
	}

	return res
}
