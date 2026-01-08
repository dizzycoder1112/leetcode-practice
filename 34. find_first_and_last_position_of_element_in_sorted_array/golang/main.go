package main

func main() {

}

func searchRange(nums []int, target int) []int {
	return []int{findLeft(nums, target), findRight(nums, target)}
}

func findLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := right - (right-left)/2

		if nums[mid] == target {
			result = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}
	return result
}

func findRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := right - (right-left)/2

		if nums[mid] == target {
			result = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}
	return result
}
