package main

func main() {}

func topKFrequentWithBucketSort(nums []int, k int) []int {
	freq := make(map[int]int, len(nums)/2)
	for i := range nums {
		freq[nums[i]] += 1
	}

	freqAr := make([][]int, len(nums)+1)

	for k, v := range freq {
		freqAr[v] = append(freqAr[v], k)
	}

	result := []int{}

	for i := len(freqAr) - 1; i >= 0 && k > 0; i-- {
		for j := len(freqAr[i]) - 1; j >= 0 && k > 0; j-- {
			result = append(result, freqAr[i][j])
			k -= 1
		}
	}

	return result
}
