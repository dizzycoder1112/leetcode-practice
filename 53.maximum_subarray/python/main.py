from typing import List

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        dp = nums[0]
        max_sum = nums[0]
        
        for i in range(1, len(nums)):
            dp = max(nums[i], dp+nums[i])
            max_sum = max(max_sum, dp)
        
        return max_sum


if __name__ == "__main__":
    sol = Solution()

    # Test 1: [-2,1,-3,4,-1,2,1,-5,4] -> [4,-1,2,1] = 6
    print(f"Test 1: {sol.maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4])}")  # Expected: 6

    # Test 2: [1] -> 单个元素
    print(f"Test 2: {sol.maxSubArray([1])}")  # Expected: 1

    # Test 3: [5,4,-1,7,8] -> 全部加起来
    print(f"Test 3: {sol.maxSubArray([5, 4, -1, 7, 8])}")  # Expected: 23

    # Test 4: [-1] -> 单个负数
    print(f"Test 4: {sol.maxSubArray([-1])}")  # Expected: -1

    # Test 5: [-2, -1] -> 全负数取最大
    print(f"Test 5: {sol.maxSubArray([-2, -1])}")  # Expected: -1
