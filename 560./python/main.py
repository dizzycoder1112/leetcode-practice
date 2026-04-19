class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        preSum = 0
        seen = {0:1}
        count = 0

        for num in nums:
            preSum += num
            
            complement = preSum - k

            if complement in seen:
                count +=1
            seen[preSum] = seen.get(preSum, 0) +1
        
        return count