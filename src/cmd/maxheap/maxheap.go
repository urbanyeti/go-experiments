package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{3, 2, 20, 5, 3, 1, 2, 5, 6, 9, 10, 4}

	// initialize the heap data structure
	h := &IntHeap{}

	// add all the values to heap, O(n log n)
	for _, val := range nums { // O(n)
		heap.Push(h, val) // O(log n)
	}

	// print all the values from the heap
	// which should be in ascending order
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d,", heap.Pop(h).(int))
	}
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}
