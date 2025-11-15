package main

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for index, item := range nums {
		hashMap[item] = index
	}

	for index, item := range nums {
		complement := target - item
		val, ok := hashMap[complement]
		if ok && val != index {
			return []int{index, val}
		}
	}

	return nil

}
func main() {
	nums1 := []int{7, 11, 15}
	target1 := 9
	twoSum(nums1, target1)
}
