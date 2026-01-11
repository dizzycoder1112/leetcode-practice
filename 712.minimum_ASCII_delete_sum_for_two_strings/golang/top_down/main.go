package topdown

func minimumDeleteSum(s1 string, s2 string) int {
	memo := make(map[[2]int]int)

	var dp func(i, j int) int
	dp = func(i, j int) int {
		// Base case
		if i == 0 && j == 0 {
			return 0
		}
		if i == 0 {
			return dp(0, j-1) + int(s2[j-1])
		}
		if j == 0 {
			return dp(i-1, 0) + int(s1[i-1])
		}

		// 查 memo
		if val, ok := memo[[2]int{i, j}]; ok {
			return val
		}

		var result int
		if s1[i-1] == s2[j-1] {
			result = dp(i-1, j-1) // 相等，问左上
		} else {
			result = min(
				dp(i-1, j)+int(s1[i-1]), // 问上面
				dp(i, j-1)+int(s2[j-1]), // 问左边
			)
		}

		memo[[2]int{i, j}] = result
		return result
	}

	return dp(len(s1), len(s2)) // 从大问题开始问
}
