package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	length := len(nums)
	result := make([][]int, 0)

	for i, num := range nums {
		leftPointer := i + 1
		rightPointer := length - 1
		if i > 0 && num == nums[i-1] {
			continue // 跳過這次迭代
		}

		for leftPointer < rightPointer {
			sum := nums[leftPointer] + nums[rightPointer]
			if sum+num < 0 {
				leftPointer++
			} else if sum+num > 0 {
				rightPointer--
			} else {
				result = append(result, []int{nums[i], nums[leftPointer], nums[rightPointer]})
				// 跳過重複的 left
				for leftPointer < rightPointer && nums[leftPointer] == nums[leftPointer+1] {
					leftPointer++
				}
				// 跳過重複的 right
				for leftPointer < rightPointer && nums[rightPointer] == nums[rightPointer-1] {
					rightPointer--
				}
				leftPointer++  // ✅ 移動指針
				rightPointer-- // ✅ 移動指針
			}
		}

	}

	return result
}
