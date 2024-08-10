package main

import (
	"container/heap"
)

type Post struct {
	ts int
	id string
}

func NewPost(ts int, id string) Post {
	return Post{
		ts: ts,
		id: id,
	}
}

type ItemHeap []Post

func (h ItemHeap) Len() int { return len(h) }

func (h ItemHeap) Less(i, j int) bool {
	return h[i].ts < h[j].ts
}

func (h ItemHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ItemHeap) Push(x interface{}) {
	*h = append(*h, x.(Post))
}

func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &ItemHeap{
		NewPost(10, "1"),
		NewPost(20, "2"),
		NewPost(30, "3"),
	}
	heap.Init(h) //记得初始化
	for h.Len() > 0 {
		post := heap.Pop(h).(Post)
		println(post.ts, post.id)
	}
}
