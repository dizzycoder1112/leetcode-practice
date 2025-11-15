package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(betterProductExceptSelf([]int{1, 2, 3, 4}))
}

// extra space complexity O(n) version
func productExceptSelf(nums []int) []int {
	n := len(nums)
	leftProduct := make([]int, n)
	rightProduct := make([]int, n)
	result := make([]int, n)

	leftProduct[0] = 1
	for i := 1; i < n; i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
	}

	rightProduct[n-1] = 1

	for i := n - 2; i >= 0; i-- {
		rightProduct[i] = rightProduct[i+1] * nums[i+1]
	}

	for i := range n {
		result[i] = leftProduct[i] * rightProduct[i]
		fmt.Println("index:", i)
	}
	fmt.Println("rightProduct: ", rightProduct)
	fmt.Println("leftProduct: ", leftProduct)

	return result

}

//extra space complexity O(1) version

func betterProductExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// 步驟 1：計算左邊的乘積，存進 answer
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	// 步驟 2：從右往左，累積右邊的乘積並更新 answer
	right := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] = answer[i] * right
		right = right * nums[i]
	}

	return answer

}
