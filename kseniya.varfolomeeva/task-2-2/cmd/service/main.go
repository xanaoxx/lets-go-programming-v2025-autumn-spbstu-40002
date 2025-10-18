package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var N int
	fmt.Scan(&N)

	dishes := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&dishes[i])
	}

	var k int
	fmt.Scan(&k)

	h := &MinHeap{}
	heap.Init(h)

	for _, num := range dishes {
		heap.Push(h, num)
	}

	var result int
	for i := 0; i < k; i++ {
		result = heap.Pop(h).(int)
	}

	fmt.Println(result)
}
