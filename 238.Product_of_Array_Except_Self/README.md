Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

 

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
 

Constraints:

2 <= nums.length <= 105
-30 <= nums[i] <= 30
The input is generated such that answer[i] is guaranteed to fit in a 32-bit integer.
 

Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)




 對於每個位置 i，計算除了 nums[i] 之外，其他所有數字的乘積。

  Example 1 解析

  nums = [1, 2, 3, 4]

  answer[0] = 2 × 3 × 4 = 24  (除了 nums[0]=1)
  answer[1] = 1 × 3 × 4 = 12  (除了 nums[1]=2)
  answer[2] = 1 × 2 × 4 = 8   (除了 nums[2]=3)
  answer[3] = 1 × 2 × 3 = 6   (除了 nums[3]=4)

  Output: [24, 12, 8, 6]



  Example 2 解析（有 0 的情況）

  nums = [-1, 1, 0, -3, 3]

  answer[0] = 1 × 0 × (-3) × 3 = 0      (除了 -1)
  answer[1] = (-1) × 0 × (-3) × 3 = 0   (除了 1)
  answer[2] = (-1) × 1 × (-3) × 3 = 9   (除了 0)  ← 關鍵！
  answer[3] = (-1) × 1 × 0 × 3 = 0      (除了 -3)
  answer[4] = (-1) × 1 × 0 × (-3) = 0   (除了 3)

  Output: [0, 0, 9, 0, 0]


 暴力解法（先理解題目）

  func productExceptSelf(nums []int) []int {
      n := len(nums)
      answer := make([]int, n)

      for i := 0; i < n; i++ {
          product := 1
          for j := 0; j < n; j++ {
              if i != j {  // 跳過自己
                  product *= nums[j]
              }
          }
          answer[i] = product
      }

      return answer
  }

  問題：時間複雜度 O(n²)，不符合題目要求（要 O(n)）