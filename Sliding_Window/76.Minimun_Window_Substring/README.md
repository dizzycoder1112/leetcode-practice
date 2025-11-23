Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.

 

Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.
 

Constraints:

m == s.length
n == t.length
1 <= m, n <= 105
s and t consist of uppercase and lowercase English letters.
 

Follow up: Could you find an algorithm that runs in O(m + n) time?

---

## 解題思路

### 算法類型
**滑動窗口 (Sliding Window)**

### 核心概念

使用兩個指針 `left` 和 `right` 維護一個窗口 `s[left:right]`：
1. **擴展階段**：`right` 指針不斷右移，直到窗口包含所有 t 中的字符
2. **收縮階段**：`left` 指針右移縮小窗口，在保持包含所有字符的前提下找最小窗口
3. 記錄過程中的最小窗口

### 詳細步驟

#### 1. 初始化
```
need[字符] = t 中該字符出現的次數
window[字符] = 當前窗口中該字符的次數
valid = 當前窗口中滿足條件的字符種類數
```

#### 2. 擴展窗口（right 右移）
```
- 將 s[right] 加入窗口
- 如果 s[right] 是需要的字符，更新 window
- 如果該字符數量滿足需求，valid++
- right++
```

#### 3. 收縮窗口（left 右移）
```
當 valid == len(need) 時（所有字符都滿足）：
  - 記錄當前窗口大小，更新最小窗口
  - 移除 s[left]
  - 如果移除的字符導致不滿足條件，valid--
  - left++
```

### 關鍵技巧

#### 1. 字符頻次追蹤
- 使用 `map[byte]int` 或 `[128]int` 數組
- 數組更快：字符的 ASCII 碼直接作為索引（`'A'` = 65）

#### 2. valid 計數器
- 不需要每次都遍歷檢查所有字符
- 只在字符數量剛好滿足 `window[c] == need[c]` 時更新
- O(1) 時間判斷窗口是否有效

#### 3. 只記錄需要的字符
```go
if _, ok := need[c]; ok {
    window[c]++  // 只追蹤 t 中出現的字符
}
```

### 複雜度分析

- **時間複雜度**：O(m + n)
  - m = len(s), n = len(t)
  - 每個字符最多被訪問兩次（right 擴展一次，left 收縮一次）

- **空間複雜度**：O(k)
  - k = 字符集大小（使用 map）
  - 或 O(128) / O(256)（使用固定大小數組）

### 示例演示

```
s = "ADOBECODEBANC", t = "ABC"
need = {A:1, B:1, C:1}

擴展到 right=5:
窗口 "ADOBEC" → valid=3 ✓
收縮 left=0 → 移除 'A' → valid=2 ✗

繼續擴展到 right=12:
窗口 "ADOBECODEBANCA" → valid=3 ✓
收縮 left 從 0→9:
  移除 A,D,O,B,E,C,O,D,E
  窗口變成 "BANC" (長度4) ← 最小窗口 ✓
```

### 優化版本

使用數組代替 map，用 count 代替 valid：

```go
freqMap := [128]int{}  // ASCII 數組
count := len(t)        // 還需匹配的字符總數

// freqMap[c] 的含義：
// > 0: 還需要這個字符
// = 0: 字符數量剛好
// < 0: 字符多餘（不在 t 中，或窗口裡多了）
```

**性能提升**：
- 數組訪問比 map 更快（直接內存訪問 vs 哈希計算）
- 單一數據結構，減少內存分配