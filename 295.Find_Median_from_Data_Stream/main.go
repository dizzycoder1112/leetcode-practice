package main

import "container/heap"

func main() {

}

type MaxHeap struct {
	values []int
}

func (h *MaxHeap) Less(i, j int) bool {
	return h.values[i] > h.values[j]
}

func (h *MaxHeap) Len() int {
	return len(h.values)
}

func (h *MaxHeap) Push(x any) {
	h.values = append(h.values, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := h.values
	n := len(old)
	x := old[n-1]
	h.values = old[:n-1]
	return x
}

func (h *MaxHeap) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

type MinHeap struct {
	values []int
}

func (h *MinHeap) Less(i, j int) bool {
	return h.values[i] < h.values[j]
}

func (h *MinHeap) Len() int {
	return len(h.values)
}

func (h *MinHeap) Push(x any) {
	h.values = append(h.values, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := h.values
	n := len(old)
	x := old[n-1]
	h.values = old[:n-1]
	return x
}

func (h *MinHeap) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

type MedianFinder struct {
	MaxHeap *MaxHeap
	MinHeap *MinHeap
}

func Constructor() MedianFinder {
	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}

	heap.Init(maxHeap)
	heap.Init(minHeap)

	return MedianFinder{
		MaxHeap: maxHeap,
		MinHeap: minHeap,
	}

}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.MaxHeap, num)
	heap.Push(this.MinHeap, heap.Pop(this.MaxHeap))
	if this.MaxHeap.Len() < this.MinHeap.Len() {
		heap.Push(this.MaxHeap, heap.Pop(this.MinHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {

	if this.MaxHeap.Len() > this.MinHeap.Len() {
		// 奇數個元素，中位數是 MaxHeap 堆頂
		return float64(this.MaxHeap.values[0])
	} else {
		// 偶數個元素，中位數是兩個堆頂的平均
		return float64(this.MaxHeap.values[0]+this.MinHeap.values[0]) / 2.0
	}

}
