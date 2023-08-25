package main

import "fmt"

type intHeap []int

func (h intHeap) Len() int      { return len(h) }
func (h intHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x int) {
	*h = append(*h, x)
	old := *h
	for i := h.Len() - 1; i > 0; i-- {
		if old[i-1] < old[i] {
			break
		}
		old.Swap(i, i-1)
	}
	*h = old
}

func longestConsecutive(nums []int) int {
	var (
		h = intHeap{}
	)

	//fixme: problem with unsorted array
	for _, v := range nums {
		h.Push(v)
	}

	var max, localMax = 1, 1
	for i := 1; i < h.Len(); i++ {
		if h[i] != h[i-1]+1 {
			if localMax > max {
				max = localMax
			}
			localMax = 1
			continue
		}
		localMax++
	}

	return max
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
