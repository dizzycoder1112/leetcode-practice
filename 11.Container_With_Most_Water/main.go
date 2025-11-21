package main

func main() {}

func maxArea(height []int) int {
	maxLength := len(height)
	left := 0
	right := maxLength - 1
	result := 0

	for right-left > 0 {
		leftHeight := height[left]
		rightHeight := height[right]

		area := min(leftHeight, rightHeight) * (right - left)

		if area > result {
			result = area
		}

		if leftHeight < rightHeight {
			left++
		} else {
			right--
		}

	}

	return result

}
