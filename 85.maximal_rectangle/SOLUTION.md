# LeetCode 85: Maximal Rectangle

## Problem

Given a `rows x cols` binary matrix filled with `0`'s and `1`'s, find the largest rectangle containing only `1`'s and return its area.

**Example:**
```
Input: matrix = [
    ["1","0","1","0","0"],
    ["1","0","1","1","1"],
    ["1","1","1","1","1"],
    ["1","0","0","1","0"]
]
Output: 6
```

---

## Pattern Analysis (題型分析)

| 直覺 | 實際 |
|------|------|
| DP（看到 maximal） | ✓ 部分正確 |
| DFS/BFS（找連通區域） | ✗ 錯誤，這題要的是矩形，不是任意形狀 |

**正確解法**：將 2D 問題轉換成多個 1D 的「Histogram 最大矩形問題」(LeetCode 84)

---

## Core Idea (核心概念)

### Step 1: Build Height Array (建立高度陣列)

把每一行當作 histogram 的底部，計算每個位置往上連續 1 的數量：

```
Original Matrix:        Heights per Row:
["1","0","1","0","0"]   [1, 0, 1, 0, 0]  ← Row 0
["1","0","1","1","1"]   [2, 0, 2, 1, 1]  ← Row 1
["1","1","1","1","1"]   [3, 1, 3, 2, 2]  ← Row 2
["1","0","0","1","0"]   [4, 0, 0, 3, 0]  ← Row 3
```

**計算規則：**
```python
if matrix[row][col] == "1":
    heights[col] = heights[col] + 1  # 累加
else:
    heights[col] = 0                  # 歸零
```

### Step 2: Largest Rectangle in Histogram (直方圖最大矩形)

對每一行的 heights 陣列，用 monotonic stack 找最大矩形。

---

## Monotonic Stack Explained (單調堆疊詳解)

### Stack 存什麼？

```python
stack: list[tuple[int, int]] = []  # (起始位置, 高度)
```

**重點**：存的不是「bar 在哪裡」，而是「**這個高度可以往左延伸到哪裡**」

### 為什麼 index 會變小？

當矮的 bar 出現時，它可以「接管」高的 bar 的起始位置：

```
heights = [4, 1]

█
█  ← 高度 4 被截斷
█  █
█  █
[4][1]
 0  1

當 h=1 出現時：
1. Pop (0, 4) — 計算高度 4 的面積 = 4 × 1 = 4
2. Push (0, 1) — 高度 1 可以往左延伸到 index 0！
```

### Step-by-Step Trace

**heights = [4, 1, 3, 3, 2]**

```
█
█        █  █
█        █  █  █
█  █     █  █  █
[4][1]  [3][3][2]
 0  1    2  3  4
```

| i | h | Action | Stack | max_area |
|---|---|--------|-------|----------|
| 0 | 4 | push (0,4) | [(0,4)] | 0 |
| 1 | 1 | pop (0,4) → area=4, push (0,1) | [(0,1)] | 4 |
| 2 | 3 | push (2,3) | [(0,1),(2,3)] | 4 |
| 3 | 3 | push (3,3) | [(0,1),(2,3),(3,3)] | 4 |
| 4 | 2 | pop (3,3), pop (2,3) → area=6, push (2,2) | [(0,1),(2,2)] | **6** |
| end | - | process remaining | - | 6 |

**i=4 詳細過程：**
```python
# 第一次 while：(3,3) 的 3 > 2
pop (3, 3) → width = 4-3 = 1, area = 3×1 = 3, start = 3

# 第二次 while：(2,3) 的 3 > 2
pop (2, 3) → width = 4-2 = 2, area = 3×2 = 6 ⭐, start = 2

# 第三次 while：(0,1) 的 1 > 2? NO，停止

push (2, 2)  # 用 start=2，不是 i=4
```

---

## Complete Python Solution

```python
from typing import List

class Solution:
    def maximalRectangle(self, matrix: List[List[str]]) -> int:
        if not matrix or not matrix[0]:
            return 0

        rows = len(matrix)
        cols = len(matrix[0])
        heights: list[int] = [0] * cols  # 每個 column 的高度
        max_area: int = 0

        for row in range(rows):
            # Step 1: 更新這一行的高度
            for col in range(cols):
                if matrix[row][col] == "1":
                    heights[col] += 1
                else:
                    heights[col] = 0

            # Step 2: 對這行的 histogram 找最大矩形
            max_area = max(max_area, self.largestRectangleInHistogram(heights))

        return max_area

    def largestRectangleInHistogram(self, heights: List[int]) -> int:
        stack: list[tuple[int, int]] = []  # (起始位置, 高度)
        max_area: int = 0

        for i, h in enumerate(heights):
            start = i
            # Pop 比當前高的 bar，計算它們的面積
            while stack and stack[-1][1] > h:
                idx, height = stack.pop()
                width = i - idx
                max_area = max(max_area, height * width)
                start = idx  # 繼承起始位置！
            stack.append((start, h))

        # 處理 stack 中剩餘的 bar
        for idx, height in stack:
            width = len(heights) - idx
            max_area = max(max_area, height * width)

        return max_area
```

---

## Complexity Analysis (複雜度分析)

| | Time | Space |
|---|------|-------|
| Build heights | O(rows × cols) | O(cols) |
| Histogram per row | O(cols) | O(cols) |
| **Total** | **O(rows × cols)** | **O(cols)** |

---

## Python Syntax Notes (Python 語法筆記)

### 1. Boolean Operators

| Python | JavaScript/C/Java |
|--------|-------------------|
| `not` | `!` |
| `and` | `&&` |
| `or` | `\|\|` |

### 2. List Repetition

```python
heights = [0] * 5  # [0, 0, 0, 0, 0]
```

### 3. Negative Indexing

```python
arr = [10, 20, 30]
arr[-1]  # 30 (last element)
arr[-2]  # 20 (second to last)
```

### 4. Tuple vs List

```python
(0, 3)  # Tuple - immutable, fixed size
[0, 3]  # List - mutable, variable size
```

### 5. enumerate()

```python
for i, h in enumerate(heights):
    # i = index, h = value
```

### 6. Type Hints

```python
stack: list[tuple[int, int]] = []
```

---

## Key Takeaways (重點整理)

1. **2D → 1D 轉換**：把每一行當作 histogram 的底部
2. **Monotonic Stack**：維護遞增的高度序列
3. **start = idx**：矮的 bar 可以往左延伸到高的 bar 曾經的位置
4. **Time O(rows × cols)**：每個 cell 最多被 push 和 pop 各一次
