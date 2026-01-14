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

	// Example 2: Expected output 1.16667
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

// Event 代表一個事件：在 y 座標，斜率變化 delta
type Event struct {
	y     int // y 座標
	delta int // 斜率變化量 (+l 進入, -l 離開)
}

func separateSquares(squares [][]int) float64 {
	// Step 1: 建立事件
	events := make([]Event, 0, len(squares)*2)
	totalArea := 0

	for _, s := range squares {
		y := s[1]
		l := s[2]
		events = append(events, Event{y: y, delta: l})      // 進入 square，斜率 +l
		events = append(events, Event{y: y + l, delta: -l}) // 離開 square，斜率 -l
		totalArea += l * l
	}

	target := float64(totalArea) / 2

	// Step 2: 按 y 座標排序
	sort.Slice(events, func(i, j int) bool {
		return events[i].y < events[j].y
	})

	// Step 3: 從下往上掃描
	slope := 0  // 當前斜率（切過的 square 寬度總和）
	area := 0.0 // 累積面積
	prevY := events[0].y

	i := 0
	for i < len(events) {
		currY := events[i].y

		// 計算從 prevY 到 currY 這一段的面積
		if currY > prevY && slope > 0 {
			addedArea := float64(slope) * float64(currY-prevY)

			// 檢查是否超過 target
			if area+addedArea >= target {
				// Step 4: 精確計算答案（線性插值）
				// area + slope * (answer - prevY) = target
				// answer = prevY + (target - area) / slope
				return float64(prevY) + (target-area)/float64(slope)
			}

			area += addedArea
		}

		// 處理所有在 currY 的事件（更新斜率）
		for i < len(events) && events[i].y == currY {
			slope += events[i].delta
			i++
		}

		prevY = currY
	}

	return float64(prevY)
}
