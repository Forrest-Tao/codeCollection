package main

import "container/heap"

type Item []int

//Less Swap Len Push Pop Init

func (h Item) Len() int {
	return len(h)
}

func (h Item) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Item) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Item) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *Item) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func findKthLargest(nums []int, k int) int {
	h := &Item{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}
