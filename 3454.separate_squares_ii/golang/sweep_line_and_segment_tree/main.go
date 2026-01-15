package main

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1: Expected output 1.00000
	squares1 := [][]int{
		{0, 0, 1},
		{2, 2, 1},
	}
	fmt.Printf("Example 1: %.5f\n", separateSquares(squares1))

	// Example 2: Expected output 1.00000 (注意：3454 的答案和 3453 不同！)
	squares2 := [][]int{
		{0, 0, 2},
		{1, 1, 1},
	}
	fmt.Printf("Example 2: %.5f\n", separateSquares(squares2))

	squares3 := [][]int{
		{522261215, 954313664, 225462},
		{628661372, 718610752, 10667},
		{619734768, 941310679, 44788},
		{352367502, 656774918, 289036},
		{860247066, 905800565, 100123},
		{817623994, 962847576, 71460},
		{691552058, 782740602, 36271},
		{911356, 152015365, 513881},
		{462847044, 859151855, 233567},
		{672324240, 954509294, 685569},
	}
	fmt.Printf("Example 3: %.5f\n", separateSquares(squares3))
}

// ============================================================
// 線段樹 (Segment Tree)
// ============================================================

type SegmentTree struct {
	xs    []int   // 壓縮後的 x 座標
	n     int     // 區間數量
	count []int   // count[i]: 節點 i 被完整覆蓋的次數
	len   []int64 // len[i]: 節點 i 的覆蓋長度
}

func NewSegmentTree(xs []int) *SegmentTree {
	n := len(xs) - 1
	size := max(4*n, 4)
	return &SegmentTree{
		xs:    xs,
		n:     n,
		count: make([]int, size),
		len:   make([]int64, size),
	}
}

// Update 更新區間 [l, r]，delta = +1 (加入) 或 -1 (移除)
// node: 當前節點編號 (從 1 開始)
// [start, end]: 當前節點代表的索引區間
// [l, r]: 要更新的索引區間
func (st *SegmentTree) Update(node, start, end, l, r, delta int) {
	if r <= start || end <= l {
		// 完全不重疊
		return
	}

	if l <= start && end <= r {
		// 完全包含，更新 count
		st.count[node] += delta
	} else {
		// 部分重疊，往下遞迴
		mid := (start + end) / 2
		st.Update(node*2, start, mid, l, r, delta)
		st.Update(node*2+1, mid, end, l, r, delta)
	}

	// 重新計算 len
	if st.count[node] > 0 {
		// 整個區間被覆蓋
		st.len[node] = int64(st.xs[end] - st.xs[start])
	} else if start+1 == end {
		// 葉節點，沒被覆蓋
		st.len[node] = 0
	} else {
		// 從子節點計算
		st.len[node] = st.len[node*2] + st.len[node*2+1]
	}
}

// Query 查詢總覆蓋長度
func (st *SegmentTree) Query() int64 {
	return st.len[1]
}

// Reset 重置線段樹
func (st *SegmentTree) Reset() {
	for i := range st.count {
		st.count[i] = 0
		st.len[i] = 0
	}
}

// ============================================================
// 主要解法
// ============================================================

// Event 代表一個事件
type Event struct {
	y     int // y 座標
	delta int // +1 進入, -1 離開
	x1    int // x 區間左端點 (壓縮後的索引)
	x2    int // x 區間右端點 (壓縮後的索引)
}

func separateSquares(squares [][]int) float64 {
	// ========== Step 1: 座標壓縮 ==========
	xSet := make(map[int]bool)
	for _, s := range squares {
		x, l := s[0], s[2]
		xSet[x] = true
		xSet[x+l] = true
	}

	// 排序 x 座標
	xs := make([]int, 0, len(xSet))
	for x := range xSet {
		xs = append(xs, x)
	}
	sort.Ints(xs)

	// 建立映射
	xToIdx := make(map[int]int)
	for i, x := range xs {
		xToIdx[x] = i
	}

	// ========== Step 2: 建立事件 ==========
	events := make([]Event, 0, len(squares)*2)
	for _, s := range squares {
		x, y, l := s[0], s[1], s[2]
		x1Idx := xToIdx[x]
		x2Idx := xToIdx[x+l]
		events = append(events, Event{y: y, delta: 1, x1: x1Idx, x2: x2Idx})      // 進入
		events = append(events, Event{y: y + l, delta: -1, x1: x1Idx, x2: x2Idx}) // 離開
	}

	// 按 y 座標排序
	sort.Slice(events, func(i, j int) bool {
		return events[i].y < events[j].y
	})

	// ========== Step 3: 建立線段樹 ==========
	st := NewSegmentTree(xs)

	// ========== Step 4: 第一次掃描，計算總面積 ==========
	var totalArea int64 = 0
	prevY := events[0].y

	i := 0
	for i < len(events) {
		currY := events[i].y

		// 計算這一段的面積
		if currY > prevY {
			width := st.Query()
			totalArea += width * int64(currY-prevY)
		}

		// 處理所有在 currY 的事件
		for i < len(events) && events[i].y == currY {
			e := events[i]
			st.Update(1, 0, st.n, e.x1, e.x2, e.delta)
			i++
		}

		prevY = currY
	}

	target := float64(totalArea) / 2

	// ========== Step 5: 第二次掃描，找答案 ==========
	st.Reset()
	var area float64 = 0
	prevY = events[0].y

	i = 0
	for i < len(events) {
		currY := events[i].y

		// 計算這一段的面積
		if currY > prevY {
			width := float64(st.Query())
			addedArea := width * float64(currY-prevY)

			// 檢查是否超過 target
			if area+addedArea >= target {
				// 答案在這一段，線性插值
				return float64(prevY) + (target-area)/width
			}

			area += addedArea
		}

		// 處理所有在 currY 的事件
		for i < len(events) && events[i].y == currY {
			e := events[i]
			st.Update(1, 0, st.n, e.x1, e.x2, e.delta)
			i++
		}

		prevY = currY
	}

	return float64(prevY)
}
