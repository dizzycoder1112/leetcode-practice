package main

func main() {}

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)

	// 建表
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Base case: 第一列（s2 是空的，删光 s1）
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}

	// Base case: 第一行（s1 是空的，删光 s2）
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
	}

	// 填表
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				// 相等，看左上角
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 不相等，看上面或左边，取最小
				dp[i][j] = min(
					dp[i-1][j]+int(s1[i-1]), // 删 s1 的字符
					dp[i][j-1]+int(s2[j-1]), // 删 s2 的字符
				)
			}
		}
	}

	return dp[m][n]
}
