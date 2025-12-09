package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	leftBoundary := 0
	maxLength := 0
	charIndex := make(map[rune]int)

	for rightBoundary, char := range s {
		lastIndex, ok := charIndex[char]
		if ok && lastIndex >= leftBoundary {
			leftBoundary = lastIndex + 1
		}

		charIndex[char] = rightBoundary

		currentLength := rightBoundary - leftBoundary + 1

		if currentLength > maxLength {
			maxLength = currentLength
		}

	}

	return maxLength

}
