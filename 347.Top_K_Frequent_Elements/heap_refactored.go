package main

// ========== 通用堆实现（MaxHeap 和 MinHeap 共享） ==========

type Heap struct {
	items []*Item
	less  func(i, j int) bool // 比较函数（唯一的差异）
	k     int                 // 最大容量（可选，0 表示无限制）
}

// NewMaxHeap 创建最大堆
func NewMaxHeap() *Heap {
	h := &Heap{items: []*Item{}}
	h.less = func(i, j int) bool {
		return h.items[i].count > h.items[j].count // 父节点 > 子节点
	}
	return h
}

// NewMinHeap 创建最小堆
func NewMinHeap2(k int) *Heap {
	h := &Heap{items: []*Item{}, k: k}
	h.less = func(i, j int) bool {
		return h.items[i].count < h.items[j].count // 父节点 < 子节点
	}
	return h
}

// ========== 以下方法对 MaxHeap 和 MinHeap 完全相同 ==========

func (h *Heap) Len() int {
	return len(h.items)
}

func (h *Heap) Less(i, j int) bool {
	return h.less(i, j) // 调用注入的比较函数
}

func (h *Heap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *Heap) up(currentNode int) {
	for {
		parentNode := (currentNode - 1) / 2
		if parentNode == currentNode || !h.Less(currentNode, parentNode) {
			break
		}
		h.swap(parentNode, currentNode)
		currentNode = parentNode
	}
}

func (h *Heap) down(parentNode int) {
	length := h.Len()
	i := parentNode
	for {
		j := i*2 + 1
		if j >= length || j < 0 {
			break
		}
		leftNode := j
		rightNode := j + 1

		// 选出两个子节点中"更符合堆性质"的那个
		if rightNode < length && h.Less(rightNode, leftNode) {
			j = rightNode
		}

		// 如果父节点已经满足堆性质
		if h.Less(i, j) {
			break
		}
		h.swap(i, j)
		i = j
	}
}

func (h *Heap) Push(item *Item) {
	h.items = append(h.items, item)
	h.up(h.Len() - 1)
}

func (h *Heap) Pop() *Item {
	n := h.Len()
	h.swap(0, n-1)
	result := h.items[n-1]
	h.items = h.items[:n-1]
	if h.Len() > 0 {
		h.down(0)
	}
	return result
}

func (h *Heap) Top() *Item {
	if h.Len() == 0 {
		return nil
	}
	return h.items[0]
}

// ========== 使用示例 ==========

// MaxHeap 解法
func topKFrequentRefactored(nums []int, k int) []int {
	freq := make(map[int]*Item)
	for _, v := range nums {
		if item, ok := freq[v]; ok {
			item.count++
		} else {
			freq[v] = &Item{value: v, count: 1}
		}
	}

	// 使用 MaxHeap
	heap := NewMaxHeap()
	for _, item := range freq {
		heap.Push(item)
	}

	// pop k 次，直接就是按频率从高到低
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop().value
	}
	return result
}

// MinHeap(k) 解法
func topKFrequentRefactoredMinHeap(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// 使用 MinHeap(k)
	heap := NewMinHeap2(k)

	for num, count := range freq {
		if heap.Len() < k {
			heap.Push(&Item{value: num, count: count})
		} else if count > heap.Top().count {
			heap.Pop()
			heap.Push(&Item{value: num, count: count})
		}
	}

	// 预分配 + 倒序填充
	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop().value
	}
	return result
}
