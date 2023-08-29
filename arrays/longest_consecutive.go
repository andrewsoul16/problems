package arrays

import "problems/pkg/heap"

func LongestConsecutive(nums []int) int {
	var (
		h = heap.IntHeap{}
	)

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
