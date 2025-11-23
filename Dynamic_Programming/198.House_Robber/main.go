package main

func main() {}

func rob(nums []int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	prev2 := nums[0]
	prev1 := max(nums[0], nums[1])

	for i := 2; i < length; i++ {
		current := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = current
	}

	return prev1

}
