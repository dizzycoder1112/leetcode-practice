package main

import "fmt"

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

func findKthLargest(nums []int, k int) int {
	heap := NewHeap(nums)
	var result int
	for range k {
		result = heap.Pop()
	}
	return result
}

// parent(i) = (i - 1) / 2    // parent node
// left(i)   = 2*i + 1        // left child node
// right(i)  = 2*i + 2        // right child node

type Heap struct {
	items []int
}

func NewHeap(nums []int) *Heap {
	heap := &Heap{
		items: nums,
	}
	heap.Init()
	return heap
}

func (h *Heap) Init() {
	n := h.Len()

	// 給定總共 n 個節點，找最後一個有子節點的節點
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *Heap) Len() int {
	return len(h.items)
}

// sort big one
func (h *Heap) Less(i, j int) bool {
	return h.items[i] > h.items[j]
}

func (h *Heap) Pop() int {
	n := h.Len()
	h.swap(0, n-1)
	result := h.items[n-1]
	h.items = h.items[:n-1]
	if h.Len() > 0 {
		h.down(0)
	}
	return result

}

func (h *Heap) down(parentNode int) {
	length := h.Len()
	i := parentNode
	for {
		j := 2*i + 1 // find left node index
		if j >= length || j < 0 {
			break
		}
		leftNode := j
		rightNode := j + 1

		//compare right and left child node which one is better
		if rightNode < length && h.Less(rightNode, leftNode) {
			j = rightNode
		}

		//compare the parent node
		if !h.Less(j, i) {
			break
		}
		h.swap(i, j)
		i = j // continue down

	}
}

func (h *Heap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]

}

func (h *Heap) up(currentNode int) {
	for {
		// find parent node
		parentNode := (currentNode - 1) / 2

		// stop if reach the root node or it satisfy the condition of heap
		if parentNode == currentNode || !h.Less(currentNode, parentNode) {
			break
		}

		// swap the parent and child
		h.swap(parentNode, currentNode)
		currentNode = parentNode // continue up
	}
}
