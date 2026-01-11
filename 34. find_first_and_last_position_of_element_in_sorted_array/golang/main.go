package main

func main() {

}

func searchRange(nums []int, target int) []int {
	return []int{find(nums, target, "left"), find(nums, target, "right")}
}

func find(nums []int, target int, direction string) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := right - (right-left)/2

		if nums[mid] == target {
			result = mid
			switch direction {
			case "left":
				right = mid - 1
			case "right":
				left = mid + 1
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}
	return result
}
