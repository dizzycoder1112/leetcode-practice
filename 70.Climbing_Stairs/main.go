package main

import "fmt"

func main() {
	fmt.Println("n=1:", climbStairs(1)) // 应该是 1
	fmt.Println("n=2:", climbStairs(2)) // 应该是 2
	fmt.Println("n=3:", climbStairs(3)) // 应该是 3
	fmt.Println("n=4:", climbStairs(4)) // 应该是 5
	fmt.Println("n=5:", climbStairs(5)) // 应该是 8
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev2 := 1 // f(1)
	prev1 := 2 // f(2)

	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

func climbStairs2(n int) int {
	memo := make(map[int]int)
	return climb(n, memo)
}

func climb(n int, memo map[int]int) int {
	if n <= 2 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val // 已经计算过，直接返回
	}
	memo[n] = climb(n-1, memo) + climb(n-2, memo)
	return memo[n]
}

func fibonacci(n int) int {
	if n <= 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
