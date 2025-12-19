package main

func main() {}

func characterReplacement(s string, k int) int {
	leftBoundary := 0
	maxLength := 0
	maxFreq := 0
	freq := make(map[rune]int)

	for rightBoundary, char := range s {
		freq[char]++
		maxFreq = max(maxFreq, freq[char])

		for (rightBoundary-leftBoundary+1)-maxFreq > k {
			freq[rune(s[leftBoundary])]--
			leftBoundary++
		}
		maxLength = max(maxLength, rightBoundary-leftBoundary+1)

	}
	return maxLength

}
