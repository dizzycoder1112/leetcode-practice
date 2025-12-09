package main

import (
	"fmt"

	"github.com/zktoner/leetcode/lru_cache/golang/internal/lru"
)

func main() {
	cache := lru.NewLRUCache(3)
	fmt.Println("Capacity:", cache.CheckCapacity())

	cache.Put(1, 100)
	cache.Put(2, 200)
	cache.Put(3, 300)

	if v, ok := cache.Get(1); ok {
		fmt.Println("Get 1:", v)
	}
	cache.Put(4, 400) // 2 應該被移除

	if v, ok := cache.Get(2); ok {
		fmt.Println("Get 2:", v)
	} else {
		fmt.Println("Get 2: Not found (evicted)")
	}

	cache.Put(5, 500) // 3 應該被移除

	if v, ok := cache.Get(3); ok {
		fmt.Println("Get 3:", v)
	} else {
		fmt.Println("Get 3: Not found (evicted)")
	}

	if v, ok := cache.Get(4); ok {
		fmt.Println("Get 4:", v)
	}
	if v, ok := cache.Get(5); ok {
		fmt.Println("Get 5:", v)
	}
}
