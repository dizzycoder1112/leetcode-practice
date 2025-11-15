package lru

import (
	"container/list"
	"fmt"
)

type entry struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

// PrintState 印出 cache 目前的 queue 狀態（從最近到最舊）
func (c *LRUCache) PrintState() {
	fmt.Print("Cache state: [")
	for e := c.list.Front(); e != nil; e = e.Next() {
		kv := e.Value.(*entry)
		fmt.Printf("(%d:%d) ", kv.key, kv.value)
	}
	fmt.Println("]")
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
		v := elem.Value.(*entry).value
		c.PrintState()
		return v, true
	}
	c.PrintState()
	return -1, false
}

func (c *LRUCache) Put(key, value int) {
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		elem.Value.(*entry).value = value
		c.PrintState()
		return
	}
	if c.list.Len() >= c.capacity {
		oldest := c.list.Back()
		if oldest != nil {
			c.list.Remove(oldest)
			delete(c.cache, oldest.Value.(*entry).key)
		}
	}
	newElem := c.list.PushFront(&entry{key: key, value: value})
	c.cache[key] = newElem
	c.PrintState()
}

func (c *LRUCache) CheckCapacity() int {
	return c.capacity
}
