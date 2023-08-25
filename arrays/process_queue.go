package arrays

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type proc struct {
	a    int
	time int
}

type procHeap []proc

func (p procHeap) Len() int           { return len(p) }
func (p procHeap) Less(i, j int) bool { return p[i].time < p[j].time }
func (p procHeap) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *procHeap) Push(x any) {
	*p = append(*p, x.(proc))
}

func (p *procHeap) Pop() any {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

func (h *procHeap) getLessThan(t int) (res []int) {
	for {
		if h.Len() == 0 {
			return
		}
		el := heap.Pop(h).(proc)
		if el.time <= t {
			res = append(res, el.a)
			continue
		}
		heap.Push(h, el)
		break
	}
	return res
}

// An IntHeap is a min-heap of ints.
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Работает с stdin и stdout
func task() {
	inFile, _ := os.Open("input")
	in := bufio.NewReader(inFile)

	var n, m int
	fmt.Fscanln(in, &n, &m)
	free := make(intHeap, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Fscan(in, &free[i])
		if err != nil {
			os.Exit(2)
		}
		fmt.Fscan(in)
	}
	heap.Init(&free)

	var (
		inUse  = &procHeap{}
		result int
	)

	for i := 0; i < m; i++ {
		var t, l int
		fmt.Fscan(in, &t, &l)
		released := inUse.getLessThan(t)
		for _, v := range released {
			heap.Push(&free, v)
		}
		if free.Len() == 0 {
			continue
		}
		a := heap.Pop(&free).(int)
		heap.Push(inUse, proc{
			a:    a,
			time: t + l,
		})
		result += a * l
	}
	fmt.Println(result)
}