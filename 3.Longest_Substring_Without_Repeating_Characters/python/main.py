class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        left = 0
        maxLength = 0
        charIndex = {}

        for right, char in enumerate(s):
            if char in charIndex.keys():
                if charIndex[char]>=left:
                    left = charIndex[char]+1
 
            charIndex[char] = right
            currentLength = right - left + 1

            maxLength = max(currentLength, maxLength)
            

        return maxLength