package main

import (
	"container/heap"
	"fmt"
	"os"
)

type MaxHeap []int

func (h *MaxHeap) Len() int { return len(*h) }

func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }

func (h *MaxHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

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
	var count int
	_, err := fmt.Scan(&count)
	if err != nil {
		os.Exit(1)
	}

	numbers := make([]int, count)
	for index := 0; index < count; index++ {
		_, err = fmt.Scan(&numbers[index])
		if err != nil {
			os.Exit(1)
		}
	}

	var order int
	_, err = fmt.Scan(&order)
	if err != nil {
		os.Exit(1)
	}

	heapInstance := &MaxHeap{}
	heap.Init(heapInstance)

	for _, number := range numbers {
		heap.Push(heapInstance, number)
	}

	var result int
	for iteration := 0; iteration < order; iteration++ {
		value := heap.Pop(heapInstance)
		intValue, ok := value.(int)
		if !ok {
			os.Exit(1)
		}
		result = intValue
	}

	fmt.Println(result)
}
