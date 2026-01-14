package main

func main() {
	// Example 1: Expected output 1.00000
	squares1 := [][]int{
		{0, 0, 1},
		{2, 2, 1},
	}
	println("Example 1:", separateSquares(squares1)) // Should be 1.0

	// Example 2: Expected output 1.16667
	squares2 := [][]int{
		{0, 0, 2},
		{1, 1, 1},
	}
	println("Example 2:", separateSquares(squares2)) // Should be ~1.16667
}

func separateSquares(squares [][]int) float64 {
	// Step 1: Calculate total area and target
	totalArea := 0
	for _, s := range squares {
		l := s[2]
		totalArea += l * l
	}
	target := float64(totalArea) / 2

	// Step 2: Find search boundaries
	left := float64(squares[0][1])            // min y
	right := float64(squares[0][1] + squares[0][2]) // max y + l

	for _, s := range squares {
		y := float64(s[1])
		l := float64(s[2])
		if y < left {
			left = y
		}
		if y+l > right {
			right = y + l
		}
	}

	// Step 3: Binary search on the answer
	for right-left > 1e-7 {
		mid := (left + right) / 2
		if areaBelow(squares, mid) < target {
			left = mid
		} else {
			right = mid
		}
	}

	return left
}

// areaBelow calculates total area below the horizontal line at height h
func areaBelow(squares [][]int, h float64) float64 {
	total := 0.0
	for _, s := range squares {
		y := float64(s[1])
		l := float64(s[2])
		top := y + l

		if h <= y {
			// Line is below this square, contributes 0
			total += 0
		} else if h >= top {
			// Line is above this square, entire square is below
			total += l * l
		} else {
			// Line cuts through the square
			total += l * (h - y)
		}
	}
	return total
}
