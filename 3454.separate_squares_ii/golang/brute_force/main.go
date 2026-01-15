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

	// Example 2: Expected output 1.00000 (3454: 重疊只算一次)
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
// 暴力法：Sweep Line + Merge Intervals
// 時間複雜度：O(n² log n)
// ============================================================

// Event 代表一個事件
type Event struct {
	y     int    // y 座標
	delta int    // +1 進入, -1 離開
	x1    int    // x 區間左端點
	x2    int    // x 區間右端點
}

// Interval 代表一個 x 區間
type Interval struct {
	start int
	end   int
}

// unionLength 計算區間聯集的總長度（LeetCode 56 的做法）
func unionLength(intervals []Interval) int64 {
	if len(intervals) == 0 {
		return 0
	}

	// Step 1: 按起點排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	// Step 2: 合併重疊區間
	merged := []Interval{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		last := &merged[len(merged)-1]

		if curr.start <= last.end {
			// 重疊，合併
			if curr.end > last.end {
				last.end = curr.end
			}
		} else {
			// 不重疊，加入新區間
			merged = append(merged, curr)
		}
	}

	// Step 3: 計算總長度
	var total int64 = 0
	for _, interval := range merged {
		total += int64(interval.end - interval.start)
	}

	return total
}

func separateSquares(squares [][]int) float64 {
	// ========== Step 1: 建立事件 ==========
	events := make([]Event, 0, len(squares)*2)

	for _, s := range squares {
		x, y, l := s[0], s[1], s[2]
		events = append(events, Event{y: y, delta: 1, x1: x, x2: x + l})         // 進入
		events = append(events, Event{y: y + l, delta: -1, x1: x, x2: x + l})    // 離開
	}

	// 按 y 座標排序
	sort.Slice(events, func(i, j int) bool {
		return events[i].y < events[j].y
	})

	// ========== Step 2: 第一次掃描，計算總面積 ==========
	active := make([]Interval, 0)  // 目前 active 的 x 區間
	var totalArea int64 = 0
	prevY := events[0].y

	i := 0
	for i < len(events) {
		currY := events[i].y

		// 計算這一段的面積
		if currY > prevY && len(active) > 0 {
			// 複製一份來計算聯集（不修改原本的 active）
			intervalsCopy := make([]Interval, len(active))
			copy(intervalsCopy, active)
			width := unionLength(intervalsCopy)
			totalArea += width * int64(currY-prevY)
		}

		// 處理所有在 currY 的事件
		for i < len(events) && events[i].y == currY {
			e := events[i]
			if e.delta == 1 {
				// 進入：加入區間
				active = append(active, Interval{start: e.x1, end: e.x2})
			} else {
				// 離開：移除區間
				for j := 0; j < len(active); j++ {
					if active[j].start == e.x1 && active[j].end == e.x2 {
						active = append(active[:j], active[j+1:]...)
						break
					}
				}
			}
			i++
		}

		prevY = currY
	}

	target := float64(totalArea) / 2

	// ========== Step 3: 第二次掃描，找答案 ==========
	active = make([]Interval, 0)
	var area float64 = 0
	prevY = events[0].y

	i = 0
	for i < len(events) {
		currY := events[i].y

		// 計算這一段的面積
		if currY > prevY && len(active) > 0 {
			intervalsCopy := make([]Interval, len(active))
			copy(intervalsCopy, active)
			width := float64(unionLength(intervalsCopy))
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
			if e.delta == 1 {
				active = append(active, Interval{start: e.x1, end: e.x2})
			} else {
				for j := 0; j < len(active); j++ {
					if active[j].start == e.x1 && active[j].end == e.x2 {
						active = append(active[:j], active[j+1:]...)
						break
					}
				}
			}
			i++
		}

		prevY = currY
	}

	return float64(prevY)
}
