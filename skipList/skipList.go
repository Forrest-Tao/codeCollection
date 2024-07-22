package main

import "math/rand"

const maxLevel = 32

type Node struct {
	val  int
	next []*Node
}

type Skiplist struct {
	head  *Node
	level int
}

func Constructor() Skiplist {
	return Skiplist{
		head: &Node{
			val:  -1,
			next: make([]*Node, maxLevel),
		},
		level: 0,
	}
}

func getRandomLevel() int {
	if rand.Intn(2) == 0 {
		return 0
	}
	return 1
}

func (this *Skiplist) Search(target int) bool {
	cur := this.head
	for i := this.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < target {
			cur = cur.next[i]
		}
	}
	return cur != nil && cur.val == target
}

func (this *Skiplist) Add(num int) {
	updated := make([]*Node, maxLevel)
	for i := range updated {
		updated[i] = this.head
	}
	cur := this.head
	for i := this.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		updated[i] = cur
	}
	this.level = max(this.level, this.level+getRandomLevel())
	newNode := &Node{
		val:  num,
		next: make([]*Node, this.level),
	}
	for i, node := range updated[0:this.level] {
		newNode.next[i] = node.next[i]
		node.next[i] = newNode
	}
}

func (this *Skiplist) Erase(num int) bool {
	updated := make([]*Node, maxLevel)
	cur := this.head
	for i := this.level - 1; i >= 0; i-- {
		for cur.next != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		updated[i] = cur
	}
	cur = cur.next[0]
	if cur == nil || cur.val != num {
		return false
	}
	for i := 0; i < this.level && updated[i] != nil; i++ {
		updated[i].next[i] = cur.next[i]
	}
	for this.level > 1 && this.head.next[this.level-1] == nil {
		this.level--
	}
	return true
}
