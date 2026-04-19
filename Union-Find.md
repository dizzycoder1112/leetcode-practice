# Union Find（並查集）詳解

## 目錄
- [從最簡單的例子開始](#從最簡單的例子開始)
- [資料結構](#資料結構)
- [核心操作 Step by Step](#核心操作-step-by-step)
- [二維轉一維](#二維轉一維)
- [完整例子：島嶼問題](#完整例子島嶼問題)
- [DFS vs Union Find](#dfs-vs-union-find)
- [程式碼模板](#程式碼模板)

---

## 從最簡單的例子開始

想像有 **5 個人**，一開始每個人都是獨立的「小團體」：

```
人員編號: 0  1  2  3  4

每個人都是獨立的團體（共 5 個團體）
```

**Union Find 做兩件事：**
1. **Union（合併）**：把兩個人放到同一個團體
2. **Find（查詢）**：找出某個人屬於哪個團體

---

## 資料結構

### parent 陣列

```python
parent = [0, 1, 2, 3, 4]
```

**意思是：**
```
parent[0] = 0  →  0 號的老大是 0（自己）
parent[1] = 1  →  1 號的老大是 1（自己）
parent[2] = 2  →  2 號的老大是 2（自己）
...

每個人都是自己的老大 = 5 個獨立團體
```

**關鍵規則：**
- `parent[i] == i` → i 是老大（根節點）
- `parent[i] != i` → i 不是老大，要繼續往上找

---

### size 陣列（可選）

```python
size = [1, 1, 1, 1, 1]
```

**意思是：**
```
size[0] = 1  →  0 號團體有 1 個人
size[1] = 1  →  1 號團體有 1 個人
...
```

> **注意：** 只有老大的 size 有意義，其他人的 size 是過時資料

---

### count 變數（可選）

```python
count = 5  # 目前有 5 個獨立團體
```

---

## 核心操作 Step by Step

### Union（合併）

#### 範例：Union(0, 1)

```
之前:
parent = [0, 1, 2, 3, 4]
          ↑  ↑
          各自獨立

步驟:
1. Find(0) = 0  （0 的老大是 0）
2. Find(1) = 1  （1 的老大是 1）
3. 不同老大 → 合併：parent[1] = 0

之後:
parent = [0, 0, 2, 3, 4]
             ↑
             1 的老大變成 0

現在 0 和 1 是同一團
```

---

#### 範例：接著 Union(1, 2)

```
之前:
parent = [0, 0, 2, 3, 4]

步驟:
1. Find(1):
   parent[1] = 0
   0 != 1 → 繼續找
   Find(0):
     parent[0] = 0
     0 == 0 → 找到老大！回傳 0
   回傳 0

2. Find(2) = 2

3. 0 和 2 不同 → 合併：parent[2] = 0

之後:
parent = [0, 0, 0, 3, 4]
                ↑
                2 的老大也變成 0

現在 0, 1, 2 是同一團
```

---

### Find（找老大）

**為什麼不能直接 return parent[x]？**

```
假設:
parent = [0, 0, 1, 3, 4]
              ↑
              2 指向 1，1 指向 0

形成鏈條：2 → 1 → 0

如果直接 return parent[2]：
  得到 1  ← 錯！1 不是老大，0 才是
```

**所以要遞迴往上找：**

```python
def find(self, x):
    if self.parent[x] != x:       # 如果我不是老大
        return self.find(self.parent[x])  # 繼續往上找
    return self.parent[x]         # 找到老大了
```

---

### 路徑壓縮

**問題：** 鏈條太長，每次 Find 都要走很多步

```
2 → 1 → 0
找 2 的老大要走 2 步
```

**優化：** 找到老大後，順便把路徑上所有人直接指向老大

```python
def find(self, x):
    if self.parent[x] != x:
        self.parent[x] = self.find(self.parent[x])  # 順便更新！
    return self.parent[x]
```

```
壓縮前:          壓縮後:
2 → 1 → 0        2 → 0
                 1 → 0

parent = [0, 0, 1]  →  parent = [0, 0, 0]

下次 Find(2) 只要一步！
```

---

## 二維轉一維

### 為什麼要轉？

因為 `parent` 陣列只能用**一個數字**當索引。

```
棋盤 3x3:
(0,0) (0,1) (0,2)
(1,0) (1,1) (1,2)
(2,0) (2,1) (2,2)

問題：怎麼用一個數字代表 (1, 1) 這個位置？
```

### 轉換公式

```
id = y * 列數 + x
```

```
棋盤 3x3（列數 = 3）:

位置:              編號:
(0,0) (0,1) (0,2)     0   1   2
(1,0) (1,1) (1,2)  →  3   4   5
(2,0) (2,1) (2,2)     6   7   8

例如:
(1,1) → id = 1 * 3 + 1 = 4
(2,0) → id = 2 * 3 + 0 = 6
```

### 找鄰居

**一維的「相鄰」不等於棋盤的「相鄰」！**

```
編號 2 和 3 在一維是連續的
但在棋盤上不相鄰：

0  1 [2]
[3] 4  5
6  7  8
```

**所以用二維座標找鄰居，再轉成一維：**

```python
# 當前位置 (x, y)
current_id = y * cols + x

# 找四個鄰居（用二維）
for dx, dy in [(1,0), (-1,0), (0,1), (0,-1)]:
    nx, ny = x + dx, y + dy
    neighbor_id = ny * cols + nx  # 轉成一維
    uf.union(current_id, neighbor_id)
```

---

## 完整例子：島嶼問題

### 題目

```
grid:
1 1 0
1 0 0
0 0 1

找出有幾個島嶼（連通的 1）
答案：2
```

### Step by Step

**編號：**
```
位置:        編號:
1 1 0        0 1 2
1 0 0   →    3 4 5
0 0 1        6 7 8
```

**初始化：**
```
parent = [0, 1, 2, 3, 4, 5, 6, 7, 8]
陸地在: 0, 1, 3, 8
count = 4（4 塊陸地）
```

**遍歷每個陸地，檢查右邊和下邊：**

```
(0,0) id=0：
  → 右邊 (1,0) id=1 是陸地 → Union(0, 1)
    Find(0)=0, Find(1)=1, 不同 → parent[1]=0, count=3
  → 下邊 (0,1) id=3 是陸地 → Union(0, 3)
    Find(0)=0, Find(3)=3, 不同 → parent[3]=0, count=2

parent = [0, 0, 2, 0, 4, 5, 6, 7, 8]
count = 2

(1,0) id=1：
  → 右邊 (2,0) id=2 是水，跳過
  → 下邊 (1,1) id=4 是水，跳過

(0,1) id=3：
  → 右邊 (1,1) 是水，跳過
  → 下邊 (0,2) 是水，跳過

(2,2) id=8：
  → 右邊、下邊都超出範圍
```

**最終：**
```
parent = [0, 0, 2, 0, 4, 5, 6, 7, 8]
count = 2

島嶼 A: {0, 1, 3}  老大是 0
島嶼 B: {8}        老大是 8
```

---

## DFS vs Union Find

### 什麼時候用哪個？

| 場景 | 推薦 | 原因 |
|------|------|------|
| 靜態圖，算一次 | DFS/BFS | 簡單直觀 |
| 動態落子，多次查詢 | Union Find | O(1) 合併和查詢 |
| 需要集合大小 | Union Find | 用 size 直接查 |

### 圍棋例子

```
DFS：每次落子後重新遍歷整個棋盤
     N 次落子 → O(N × m × n)

Union Find：落子時只 Union 相鄰同色
           N 次落子 → O(N × 4) ≈ O(N)
```

---

## 程式碼模板

### Python（完整版）

```python
class UnionFind:
    def __init__(self, n: int):
        self.parent = list(range(n))  # 每個節點的老大
        self.size = [1] * n           # 每個集合的大小
        self.count = n                # 集合數量

    def find(self, x: int) -> int:
        """找老大（帶路徑壓縮）"""
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x: int, y: int) -> bool:
        """合併兩個集合，回傳是否成功合併"""
        root_x = self.find(x)
        root_y = self.find(y)

        if root_x == root_y:
            return False  # 已經同一團

        # 小的併入大的（按秩合併）
        if self.size[root_x] < self.size[root_y]:
            self.parent[root_x] = root_y
            self.size[root_y] += self.size[root_x]
        else:
            self.parent[root_y] = root_x
            self.size[root_x] += self.size[root_y]

        self.count -= 1
        return True

    def get_size(self, x: int) -> int:
        """取得 x 所屬集合的大小"""
        return self.size[self.find(x)]

    def connected(self, x: int, y: int) -> bool:
        """判斷 x 和 y 是否在同一集合"""
        return self.find(x) == self.find(y)
```

### Python（簡潔版，面試用）

```python
class UnionFind:
    def __init__(self, n):
        self.parent = list(range(n))

    def find(self, x):
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x, y):
        px, py = self.find(x), self.find(y)
        if px != py:
            self.parent[px] = py
```

### Go

```go
type UnionFind struct {
    parent []int
    size   []int
    count  int
}

func NewUnionFind(n int) *UnionFind {
    uf := &UnionFind{
        parent: make([]int, n),
        size:   make([]int, n),
        count:  n,
    }
    for i := 0; i < n; i++ {
        uf.parent[i] = i
        uf.size[i] = 1
    }
    return uf
}

func (uf *UnionFind) Find(x int) int {
    if uf.parent[x] != x {
        uf.parent[x] = uf.Find(uf.parent[x])
    }
    return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
    rootX, rootY := uf.Find(x), uf.Find(y)
    if rootX == rootY {
        return false
    }
    if uf.size[rootX] < uf.size[rootY] {
        uf.parent[rootX] = rootY
        uf.size[rootY] += uf.size[rootX]
    } else {
        uf.parent[rootY] = rootX
        uf.size[rootX] += uf.size[rootY]
    }
    uf.count--
    return true
}
```

---

## 複雜度

| 操作 | 時間複雜度 |
|------|-----------|
| Find | O(α(n)) ≈ O(1) |
| Union | O(α(n)) ≈ O(1) |
| 查詢大小 | O(1) |
| 查詢數量 | O(1) |

> α(n) 是反阿克曼函數，對任何實際 n 都 ≤ 4

---

## 總結

### 核心概念

```
parent[i] = i  →  i 是老大
parent[i] ≠ i  →  繼續往上找
```

### 面試記憶點

1. **Find**：遞迴找老大 + 路徑壓縮
2. **Union**：先 Find 兩邊的老大，再合併老大
3. **時間複雜度**：幾乎 O(1)
4. **適用場景**：動態連通性、合併集合、多次查詢

### 推薦練習題

| 題號 | 題目 | 難度 |
|------|------|------|
| 200 | Number of Islands | Medium |
| 695 | Max Area of Island | Medium |
| 547 | Number of Provinces | Medium |
| 684 | Redundant Connection | Medium |
| 721 | Accounts Merge | Medium |
