package main

func main() {}

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left := 0
	right := len(height) - 1
	leftMax := 0
	rightMax := 0
	result := 0

	for left < right {
		if height[left] < height[right] {
			if leftMax > height[left] {
				result = leftMax - height[left] + result
			} else {
				leftMax = height[left]
			}
			left++
		} else {
			if rightMax > height[right] {
				result = rightMax - height[right] + result
			} else {
				rightMax = height[right]
			}
			right--
		}
	}
	return result

}
