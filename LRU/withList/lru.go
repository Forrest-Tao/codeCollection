package withList

import (
	"container/list"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type Entry struct {
	key   int
	value int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (c *LRUCache) Get(key int) (int, bool) {
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		return elem.Value.(*Entry).value, true
	}
	return -1, false
}

func (c *LRUCache) Put(key int, value int) {
	//key存在，则更新
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		elem.Value.(*Entry).value = value
		return
	}
	if c.list.Len() >= c.capacity {
		elm := c.list.Back()
		if elm != nil {
			c.list.Remove(elm)
			delete(c.cache, elm.Value.(*Entry).key)
		}
	}
	//key不存在，则插入
	entry := &Entry{key, value}
	elem := c.list.PushFront(entry)
	c.cache[key] = elem
}

func (c *LRUCache) showList() {
	if c.list.Len() == 0 {
		return
	}
	cur := c.list.Front()
	if cur != nil {
		println("For debug")
	}
	for cur != nil {
		println(cur.Value.(*Entry).key, cur.Value.(*Entry).value)
		cur = cur.Next()
	}
	println("For debug")
}
