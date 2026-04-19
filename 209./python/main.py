class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        left = 0
        minLength = float('inf')
        total = 0

        for right, num in enumerate(nums):

            total += num
            
            while target <= total:
                currentLength = right - left+1
                minLength = min(currentLength, minLength)
                total = total - nums[left]
                left +=1

        return minLength if minLength != float('inf') else 0