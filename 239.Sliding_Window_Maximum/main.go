package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	fmt.Println("Input:", nums)
	fmt.Println("Window size:", k)
	fmt.Println()

	result1 := maxSlidingWindowHeap(nums, k)
	fmt.Println("Heap Solution (O(n log n)):", result1)

	result2 := maxSlidingWindowDeque(nums, k)
	fmt.Println("Deque Solution (O(n)):", result2)
}

// ============================================
// 解法1: Max Heap - O(n log n)
// ============================================
// 思路：
// 1. 用最大堆維護窗口內的元素（存儲 value 和 index）
// 2. 每次窗口滑動，加入新元素到堆
// 3. 檢查堆頂是否在窗口外，如果是就移除
// 4. 堆頂就是當前窗口最大值
//
// 時間複雜度：O(n log n)
// - 每個元素入堆一次：O(log n)
// - 最壞情況下每個元素出堆一次：O(log n)
// - 總共 n 個元素 → O(n log n)
//
// 空間複雜度：O(n) - 堆最多存 n 個元素

type HeapItem struct {
	value int
	index int
}

type MaxHeap []HeapItem

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].value > h[j].value } // 最大堆
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(HeapItem))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxSlidingWindowHeap(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	result := make([]int, 0, len(nums)-k+1)
	h := &MaxHeap{}
	heap.Init(h)

	for i := range nums {
		// 加入新元素到堆 - O(log n)
		heap.Push(h, HeapItem{value: nums[i], index: i})

		// 移除所有不在窗口內的元素（堆頂的過期元素）
		// 注意：這裡可能需要多次 pop - O(log n) per pop
		for h.Len() > 0 && (*h)[0].index < i-k+1 {
			heap.Pop(h)
		}

		// 窗口形成後，記錄最大值
		if i >= k-1 {
			result = append(result, (*h)[0].value)
		}
	}

	return result
}

// ============================================
// 解法2: Monotonic Deque (單調雙端隊列) - O(n)
// ============================================
// 思路：
// 1. 維護一個遞減的雙端隊列（存儲索引）
// 2. 隊首永遠是當前窗口的最大值索引
// 3. 新元素進來時，從隊尾移除所有比它小的元素（它們永遠不可能是最大值）
// 4. 如果隊首索引超出窗口，從隊首移除
//
// 時間複雜度：O(n)
// - 每個元素最多進隊一次：O(1)
// - 每個元素最多出隊一次：O(1)
// - 總共 n 個元素 → O(n)
//
// 空間複雜度：O(k) - 雙端隊列最多存 k 個元素

func maxSlidingWindowDeque(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	result := make([]int, 0, len(nums)-k+1)
	deque := make([]int, 0, k) // 存儲索引

	for i := range nums {
		// 1. 移除窗口外的元素（隊首） - O(1)
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:] // 移除隊首
		}

		// 2. 維護遞減性質：從隊尾移除所有比當前元素小的元素
		//    因為它們永遠不可能成為最大值
		//    關鍵：每個元素最多被移除一次！ - 均攤 O(1)
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1] // 移除隊尾
		}

		// 3. 加入當前元素索引 - O(1)
		deque = append(deque, i)

		// 4. 窗口形成後，隊首就是最大值 - O(1)
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

// ============================================
// 性能對比總結
// ============================================
//
// Heap 解法：
// ✓ 較容易理解（標準堆操作）
// ✗ O(n log n) 時間複雜度
// ✗ 需要額外處理過期元素
// ✗ 空間複雜度 O(n)
//
// Monotonic Deque 解法：
// ✓ O(n) 最優時間複雜度
// ✓ 空間複雜度更優 O(k)
// ✓ 每個元素最多進出隊列一次
// ✗ 思路相對不直觀（需要理解單調性維護）
//
// 當 n = 10^5, k = 10^4 時：
// Heap: ~10^5 * log(10^5) ≈ 1,660,000 次操作
// Deque: ~10^5 次操作（快 16 倍！）
