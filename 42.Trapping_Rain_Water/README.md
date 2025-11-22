Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

 

Example 1:


Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
Example 2:

Input: height = [4,2,0,3,2,5]
Output: 9
 

Constraints:

n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105



用 [3,0,0,2] 详细执行

  举例 [3,0,0,2]：
  索引: 0 1 2 3
  高度: 3 0 0 2
  
  █
  █ ~ ~ █
  █ ~ ~ █


  初始状态：
  - left=0, right=3
  - left_max=0, right_max=0
  - result=0

  Step 1: left=0, right=3

  比较 height[0]=3 vs height[3]=2
  3 > 2，所以走 else 分支（移动 right）
  - height[3]=2 >= right_max=0
  - 更新 right_max = 2
  - 无积水
  - right-- → right=2

  Step 2: left=0, right=2

  比较 height[0]=3 vs height[2]=0
  3 > 0，所以走 else 分支（移动 right）
  - height[2]=0 < right_max=2
  - result += 2 - 0 = 2 ✅（用的是 right_max，不是3！）
  - right-- → right=1

  Step 3: left=0, right=1

  比较 height[0]=3 vs height[1]=0
  3 > 0，所以走 else 分支（移动 right）
  - height[1]=0 < right_max=2
  - result += 2 - 0 = 2 ✅（还是用 right_max）
  - right-- → right=0

  Step 4: left=0, right=0

  - left < right 不成立，循环结束
  - 总水量 = 4 ✅