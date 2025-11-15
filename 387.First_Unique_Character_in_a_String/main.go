package main

func main() {
	example_1 := "leetcode"

	firstUniqChar(example_1)

}

func firstUniqChar(s string) int {
	counts := make(map[rune]int)

	for _, e := range s {
		counts[e]++
	}

	for i, e := range s {
		val, ok := counts[e]
		if ok && val == 1 {
			return i
		}
	}
	return -1
}
