package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	dishes := make([]int, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&dishes[i])
	}

	var k int
	_, _ = fmt.Scan(&k)

	result := findKthSmallest(dishes, k)
	fmt.Println(result)
}

func findKthSmallest(arr []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)

	for _, num := range arr {
		heap.Push(h, num)
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}

	return (*h)[0]
}
