package main

import "math"

func main() {}

func minWindow(s string, t string) string {
	// 1. 初始化 need 和 window
	need := make(map[byte]int)
	window := make(map[byte]int)

	// 2. 统计 t 中每个字符需要的次数
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	// 3. 初始化变量
	left, right := 0, 0
	valid := 0                        // 窗口中满足条件的字符种类数
	start, length := 0, math.MaxInt32 // 记录最小窗口的起始位置和长度

	// 4. 右指针不断扩展
	for right < len(s) {
		// 4.1 取出右边界字符
		c := s[right]
		// 4.2 右指针右移，扩大窗口
		right++

		// 4.3 更新窗口数据
		if _, ok := need[c]; ok { // 如果这个字符是我们需要的
			window[c]++
			if window[c] == need[c] { // 如果窗口中这个字符的数量满足需求了
				valid++
			}
		}

		// 5. 判断是否需要收缩窗口
		for valid == len(need) { // 当所有字符都满足条件时
			// 5.1 更新最小窗口
			if right-left < length {
				start = left
				length = right - left
			}

			// 5.2 取出左边界字符
			d := s[left]
			// 5.3 左指针右移，缩小窗口
			left++

			// 5.4 更新窗口数据
			if _, ok := need[d]; ok { // 如果移除的字符是我们需要的
				if window[d] == need[d] { // 如果移除后会导致不满足条件
					valid--
				}
				window[d]--
			}
		}
	}

	// 6. 返回结果
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}
