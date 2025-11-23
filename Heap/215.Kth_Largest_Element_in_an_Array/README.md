Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?

 

Example 1:

Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
Example 2:

Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4
 

Constraints:

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104


line:44

n/2 - 1 - 找最後一個非葉子節點

  // 給定總共 n 個節點，找最後一個有子節點的節點
  lastParent := n/2 - 1

  // 例如：n=6 個節點
  lastParent = 6/2 - 1 = 2

  為什麼 Init 要用 n/2 - 1？

  看這個例子：
  陣列: [3, 2, 1, 5, 6, 4]  // n = 6
  索引:  0  1  2  3  4  5

  樹狀結構:
        0(3)
       /    \
     1(2)    2(1)
     / \     /
   3(5) 4(6) 5(4)

  葉子節點: 3, 4, 5 (沒有子節點)
  非葉子節點: 0, 1, 2 (有子節點)
