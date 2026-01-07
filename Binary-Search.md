func binarySearch(nums []int, target int) int {
      left, right := 0, len(nums)-1

      for left <= right {
          mid := left + (right-left)/2

          if nums[mid] == target {
              return mid           // 找到了
          } else if nums[mid] < target {
              left = mid + 1       // target 在右邊
          } else {
              right = mid - 1      // target 在左邊
          }
      }

      return -1  // 沒找到
  }

  圖解

  nums = [1, 3, 5, 7, 9, 11], target = 7

  Round 1:
  [1, 3, 5, 7, 9, 11]
   L     M         R     mid=2, nums[mid]=5 < 7 → left = 3

  Round 2:
  [1, 3, 5, 7, 9, 11]
            L  M   R     mid=4, nums[mid]=9 > 7 → right = 3

  Round 3:
  [1, 3, 5, 7, 9, 11]
            L
            M
            R            mid=3, nums[mid]=7 == target ✓ return 3

  關鍵記憶點

  | 部分                             | 重點                       |
  |----------------------------------|----------------------------|
  | left <= right                    | 用 <= 確保單一元素也會檢查 |
  | mid = left + (right-left)/2      | 避免 overflow              |
  | left = mid + 1 / right = mid - 1 | 要 +1/-1，否則可能無限迴圈 |
