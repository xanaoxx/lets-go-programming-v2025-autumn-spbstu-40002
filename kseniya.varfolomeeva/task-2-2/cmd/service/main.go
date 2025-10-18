package main

import (
	"container/heap"
	"fmt"
	"os"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	value, ok := x.(int)
	if !ok {
		return
	}
	*h = append(*h, value)
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func main() {
	var dishCount int
	_, err := fmt.Scan(&dishCount)
	if err != nil {
		os.Exit(1)
	}

	dishes := make([]int, dishCount)
	for i := range dishes {
		_, err = fmt.Scan(&dishes[i])
		if err != nil {
			os.Exit(1)
		}
	}

	var preferenceOrder int
	_, err = fmt.Scan(&preferenceOrder)
	if err != nil {
		os.Exit(1)
	}

	heapInstance := &MaxHeap{}
	heap.Init(heapInstance)

	for _, num := range dishes {
		heap.Push(heapInstance, num)
	}

	var result int
	for i := 0; i < preferenceOrder; i++ {
		value := heap.Pop(heapInstance)
		intValue, ok := value.(int)
		if !ok {
			os.Exit(1)
		}
		result = intValue
	}

	fmt.Println(result)
}
