package main

func main() {}

func topKFrequent(nums []int, k int) []int {
	numsMap := make(map[int]*Item)
	heap := NewHeap()
	result := []int{}
	for _, v := range nums {
		num, ok := numsMap[v]
		if !ok {
			numsMap[v] = &Item{
				value: v,
				count: 1,
			}
		} else {
			num.count++
		}
	}

	for _, v := range numsMap {
		heap.push(v)
	}

	for range k {
		result = append(result, heap.pop().value)
	}
	return result
}

type Item struct {
	count int
	value int
}

type MaxHeap struct {
	items []*Item
}

func NewHeap() *MaxHeap {
	heap := &MaxHeap{
		items: []*Item{},
	}
	return heap
}

func (h *MaxHeap) Len() int {
	return len(h.items)
}

func (h *MaxHeap) Less(i, j int) bool {
	return h.items[i].count > h.items[j].count
}

func (h *MaxHeap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]

}

func (h *MaxHeap) down(parentNode int) {
	length := h.Len()
	i := parentNode
	for {
		j := i*2 + 1
		if j >= length || j < 0 {
			break
		}
		leftNode := j
		rightNode := j + 1

		if rightNode < length && h.Less(rightNode, leftNode) {
			j = rightNode
		}

		if h.Less(i, j) {
			break
		}
		h.swap(i, j)
		i = j
	}
}

func (h *MaxHeap) up(currentNode int) {
	for {
		parentNode := (currentNode - 1) / 2
		if parentNode == currentNode || !h.Less(currentNode, parentNode) {
			break
		}

		h.swap(parentNode, currentNode)
		currentNode = parentNode
	}
}

func (h *MaxHeap) push(item *Item) {
	h.items = append(h.items, item)
	h.up(h.Len() - 1)
}

func (h *MaxHeap) pop() *Item {
	n := h.Len()
	h.swap(0, n-1)
	result := h.items[n-1]
	h.items = h.items[:n-1]
	if h.Len() > 0 {
		h.down(0)
	}
	return result

}

func topKFrequentWithBucketSort(nums []int, k int) []int {
	freq := make(map[int]int, len(nums)/2)
	for i := range nums {
		freq[nums[i]] += 1
	}

	freqAr := make([][]int, len(nums)+1)

	for k, v := range freq {
		freqAr[v] = append(freqAr[v], k)
	}

	result := []int{}

	for i := len(freqAr) - 1; i >= 0 && k > 0; i-- {
		for j := len(freqAr[i]) - 1; j >= 0 && k > 0; j-- {
			result = append(result, freqAr[i][j])
			k -= 1
		}
	}

	return result
}

// ========== MinHeap Implementation ==========

type MinHeap struct {
	items []*Item
	k     int // 最大容量
}

func NewMinHeap(k int) *MinHeap {
	return &MinHeap{
		items: []*Item{},
		k:     k,
	}
}

func (h *MinHeap) Len() int {
	return len(h.items)
}

// Less: MinHeap 的特点是父节点 <= 子节点
func (h *MinHeap) Less(i, j int) bool {
	return h.items[i].count < h.items[j].count
}

func (h *MinHeap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

// up: 向上调整（用于 push）
func (h *MinHeap) up(currentNode int) {
	for {
		parentNode := (currentNode - 1) / 2
		if parentNode == currentNode || !h.Less(currentNode, parentNode) {
			break
		}
		h.swap(parentNode, currentNode)
		currentNode = parentNode
	}
}

// down: 向下调整（用于 pop）
func (h *MinHeap) down(parentNode int) {
	length := h.Len()
	i := parentNode
	for {
		j := i*2 + 1
		if j >= length || j < 0 {
			break
		}
		leftNode := j
		rightNode := j + 1

		// 选出两个子节点中较小的那个
		if rightNode < length && h.Less(rightNode, leftNode) {
			j = rightNode
		}

		// 如果父节点已经小于等于子节点，满足堆性质
		if h.Less(i, j) {
			break
		}
		h.swap(i, j)
		i = j
	}
}

// Push: 添加元素
func (h *MinHeap) Push(item *Item) {
	h.items = append(h.items, item)
	h.up(h.Len() - 1)
}

// Pop: 移除并返回堆顶元素（最小的）
func (h *MinHeap) Pop() *Item {
	n := h.Len()
	h.swap(0, n-1)
	result := h.items[n-1]
	h.items = h.items[:n-1]
	if h.Len() > 0 {
		h.down(0)
	}
	return result
}

// Top: 查看堆顶元素但不删除
func (h *MinHeap) Top() *Item {
	if h.Len() == 0 {
		return nil
	}
	return h.items[0]
}

// ========== 使用 MinHeap(k) 的解法 ==========

func topKFrequentWithMinHeap(nums []int, k int) []int {
	// 1. 统计频率
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// 2. 用 MinHeap(k) 维护频率最高的 k 个元素
	minHeap := NewMinHeap(k)

	for num, count := range freq {
		if minHeap.Len() < k {
			// 堆还没满，直接加
			minHeap.Push(&Item{value: num, count: count})
		} else if count > minHeap.Top().count {
			// 新元素频率比堆顶大，替换堆顶
			minHeap.Pop()
			minHeap.Push(&Item{value: num, count: count})
		}
		// else: 新元素频率 <= 堆顶，忽略（肯定不是 top k）
	}

	// 3. 提取结果（预分配 + 倒序填充）
	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = minHeap.Pop().value
	}

	return result
}
