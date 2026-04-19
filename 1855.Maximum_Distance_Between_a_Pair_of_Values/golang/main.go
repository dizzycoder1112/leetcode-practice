package main

func main() {
	nums1, nums2 := []int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}
	maxDistance(nums1, nums2)

	nums1, nums2 = []int{2, 2, 2}, []int{10, 10, 1}
	maxDistance(nums1, nums2)

	nums1, nums2 = []int{30, 29, 19, 5}, []int{25, 25, 25, 25, 25}
	maxDistance(nums1, nums2)
}

func maxDistance(nums1 []int, nums2 []int) int {
	j := 0
	ans := 0
	n2 := len(nums2)
	for i, v := range nums1 {
		if j < i {
			j = i
		}
		for j < n2 && v <= nums2[j] {
			ans = max(ans, j-i)
			j++
		}
	}
	return ans
}
