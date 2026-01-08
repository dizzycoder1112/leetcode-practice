package main

func main() {}

func maxDotProduct(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := range m {
		for j := range n {
			dp[i][j] = nums1[i] * nums2[j]
			if i > 0 && j > 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+nums1[i]*nums2[j])
			}
			if i > 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j])
			}
			if j > 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-1])
			}
		}
	}

	return dp[m-1][n-1]
}
