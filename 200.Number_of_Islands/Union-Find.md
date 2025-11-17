# Union Find（并查集）详解

## 目录
- [核心思想](#核心思想)
- [数据结构](#数据结构)
- [核心操作](#核心操作)
  - [Find 方法详解](#find-方法详解)
  - [Union 方法详解](#union-方法详解)
- [完整合并过程演示](#完整合并过程演示)
- [应用到岛屿问题](#应用到岛屿问题)
- [代码实现](#代码实现)

---

## 核心思想

**"把所有相邻的陆地合并成一个集合"**

用朋友圈类比：
- 每块陆地一开始都是**独立的集合**
- 如果两块陆地相邻，就把它们**合并到同一个集合**
- 最后统计有**多少个独立的集合**

---

## 数据结构

### parent 数组

```go
parent[i] = j  // 表示 "节点 i 的父节点是 j"
```

**关键规则：**
- `parent[i] == i` → **i 是根节点**（老大）
- `parent[i] != i` → **i 不是根节点**，它还有上级

### 初始化

```go
parent = [0, 1, 2, 3, 4, 5]
         ↑  ↑  ↑  ↑  ↑  ↑
        每个节点都是自己的根
```

**视觉化：**
```
0  1  2  3  4  5
↑  ↑  ↑  ↑  ↑  ↑
每个都是独立的根节点
```

---

## 核心操作

### Find 方法详解

**目的：** 找到 x 所在集合的根节点（老大）

```go
func (uf *UnionFind) Find(x int) int {
    if uf.parent[x] != x {
        // 递归找根节点，同时做路径压缩
        uf.parent[x] = uf.Find(uf.parent[x])
    }
    return uf.parent[x]
}
```

#### 详细执行过程

假设树结构：
```
    3
   ╱
  2
 ╱
1
╱
0

parent = [1, 2, 3, 3]
```

**调用 `Find(0)`：**

**第 1 层递归：Find(0)**
```go
parent[0] = 1, 1 != 0 ✅
→ 调用 Find(1)
```

**第 2 层递归：Find(1)**
```go
parent[1] = 2, 2 != 1 ✅
→ 调用 Find(2)
```

**第 3 层递归：Find(2)**
```go
parent[2] = 3, 3 != 2 ✅
→ 调用 Find(3)
```

**第 4 层递归：Find(3) - 终止条件**
```go
parent[3] = 3, 3 == 3 ❌ 不成立
→ return 3  // 找到根节点！
```

**回溯（路径压缩）：**
```go
Find(2): parent[2] = 3, return 3
Find(1): parent[1] = 3, return 3  // 路径压缩！
Find(0): parent[0] = 3, return 3  // 路径压缩！
```

**结果：**
```go
parent = [3, 3, 3, 3]  // 所有节点都直接连到根！

优化前：          优化后：
    3                 3
   ╱               ╱ | ╲
  2               0  1  2
 ╱
1
╱
0
```

#### 关键要点

1. **递归终止条件：** `parent[x] == x`（找到根节点）
2. **递归方向：** 不断往上找父节点
3. **路径压缩：** 回溯时把路径上所有节点直接连到根
4. **为什么能找到根？** 因为根节点的定义就是 `parent[x] == x`

---

### Union 方法详解

**目的：** 把 x 和 y 所在的两个集合合并成一个

```go
func (uf *UnionFind) Union(x, y int) {
    rootX := uf.Find(x)  // 找 x 的根
    rootY := uf.Find(y)  // 找 y 的根

    if rootX != rootY {
        uf.parent[rootX] = rootY  // 让 rootX 指向 rootY
        uf.count--                // 集合数量减 1
    }
}
```

#### 详细执行过程

**假设初始状态：**
```go
parent = [1, 1, 2, 3, 4, 5]
count = 5

树结构：
1  2  3  4  5
╱  ↑  ↑  ↑  ↑
0
```

**执行 `Union(0, 3)`：**

**第 1 步：找 0 的根**
```go
rootX = Find(0)
  parent[0] = 1, 1 != 0
  → Find(1)
    parent[1] = 1, 1 == 1
    → return 1
  → return 1
rootX = 1  // 注意：不是 0，而是 0 的根 1！
```

**第 2 步：找 3 的根**
```go
rootY = Find(3)
  parent[3] = 3, 3 == 3
  → return 3
rootY = 3
```

**第 3 步：合并**
```go
rootX != rootY  // 1 != 3 ✅
parent[rootX] = rootY  // parent[1] = 3  ← 关键！修改的是根！
count--  // count = 4
```

**结果：**
```go
parent = [1, 3, 2, 3, 4, 5]
         ↑  ↑
         0  1
      保持不变 改变了！

树结构：
  3     2  4  5
 ╱ ╲    ↑  ↑  ↑
1  (3)
╱
0
```

#### 关键要点

1. **修改的是根节点：** `parent[rootX] = rootY`，不是 `parent[x]`！
2. **集合合并：** 两个集合变成一个，`count--`
3. **原来的根节点：** 从"根"降级为"中间节点"
4. **关系链式：** 0 → 1 → 3（层层连接）

---

## 完整合并过程演示

### 示例网格

```
1 1 1
1 1 1

编号（二维转一维）：
0 1 2
3 4 5

编号公式：id = i * 列数 + j
```

### 初始状态

```go
parent = [0, 1, 2, 3, 4, 5]
count = 6  // 6 块陆地，6 个独立集合

视觉化：
0  1  2  3  4  5
↑  ↑  ↑  ↑  ↑  ↑
```

### 依序合并过程

| 步骤 | 操作 | parent 数组 | count | 说明 |
|------|------|-------------|-------|------|
| 初始 | - | `[0,1,2,3,4,5]` | 6 | 6 个独立集合 |
| 1 | Union(0,1) | `[1,1,2,3,4,5]` | 5 | 0→1 |
| 2 | Union(0,3) | `[1,3,2,3,4,5]` | 4 | 1→3 (关键：修改的是 parent[1]，不是 parent[0]！) |
| 3 | Union(1,2) | `[1,3,2,2,4,5]` | 3 | 3→2 |
| 4 | Union(1,4) | `[1,2,4,2,4,5]` | 2 | 2→4，路径压缩 1→2 |
| 5 | Union(2,5) | `[1,2,4,2,5,5]` | 1 | 4→5 ✅ |
| 6 | Union(3,4) | `[1,2,5,5,5,5]` | 1 | 已合并，路径压缩 |
| 7 | Union(4,5) | `[1,2,5,5,5,5]` | 1 | 已合并 |

### 步骤 2 详解（最容易混淆的地方）

**执行 Union(0, 3)：**

**当前状态：**
```go
parent = [1, 1, 2, 3, 4, 5]

集合：{0, 1} 根是 1
      {2} 根是 2
      {3} 根是 3
      ...
```

**过程：**
```go
rootX = Find(0) = 1  // 0 的根是 1，不是 0！
rootY = Find(3) = 3

parent[rootX] = rootY  // parent[1] = 3
```

**结果：**
```go
parent = [1, 3, 2, 3, 4, 5]
         ↑  ↑
    没变！改变了！

// 不是 [3, 1, 2, 1, 4, 5] ❌ 错误！
// 而是 [1, 3, 2, 3, 4, 5] ✅ 正确！
```

**为什么？**
- Union 修改的是**根节点**的关系，不是原始参数！
- `parent[1] = 3` 表示"集合 {0,1} 的根（1）指向集合 {3} 的根（3）"

### 最终树结构

```
       5
    ╱ ╱│╲ ╲
   2 3 4 (5)
  ╱
 1
╱
0

所有节点的根都是 5：
Find(0) → 5
Find(1) → 5
Find(2) → 5
Find(3) → 5
Find(4) → 5
Find(5) → 5
```

---

## 应用到岛屿问题

### 解题思路

1. **初始化：** 每块陆地都是独立的岛屿
2. **合并：** 如果两块陆地相邻，就合并成同一个岛屿
3. **结果：** 最后有多少个独立的集合

### 为什么只检查右边和下边？

```
当前格子：*

检查四个方向：
  上
  ↑
左← * →右
  ↓
  下
```

**如果检查所有四个方向：**
- (0,0) 和 (0,1) 会互相检查，合并两次（重复）

**只检查右和下：**
- 每对相邻的陆地只合并一次
- 因为从左往右、从上往下扫描，左边和上边的已经处理过了

### 示例演示

网格：
```
1 1 0
1 0 1
```

编号：
```
0 1 2
3 4 5
```

**步骤 1：初始化**
```go
parent = [0, 1, 2, 3, 4, 5]
count = 4  // 只有 0, 1, 3, 5 是陆地
```

**步骤 2：检查相邻关系并合并**

检查 (0,0) - id=0：
```go
右边 (0,1) - id=1, 值='1' ✅
Union(0, 1) → parent = [1,1,2,3,4,5], count = 3

下边 (1,0) - id=3, 值='1' ✅
Union(0, 3) → parent = [1,3,2,3,4,5], count = 2
```

检查 (0,1) - id=1：
```go
右边 (0,2) - id=2, 值='0' ❌
下边 (1,1) - id=4, 值='0' ❌
```

检查 (1,2) - id=5：
```go
右边：越界
下边：越界
```

**步骤 3：结果**
```go
count = 2  // 2 个岛屿

集合 1: {0, 1, 3}  → 左边的岛
集合 2: {5}        → 右下的岛

树结构：
  3     5
 ╱ ╲    ↑
1  (3)
╱
0
```

---

## 代码实现

### Go 完整实现

```go
type UnionFind struct {
    parent []int  // parent[i] = i 的父节点
    count  int    // 当前集合数量
}

func NewUnionFind(n int) *UnionFind {
    uf := &UnionFind{
        parent: make([]int, n),
        count:  0,
    }
    // 初始化：每个元素的父节点是自己
    for i := 0; i < n; i++ {
        uf.parent[i] = i
    }
    return uf
}

// Find - 找到 x 的根节点（老大）
func (uf *UnionFind) Find(x int) int {
    if uf.parent[x] != x {
        // 路径压缩：递归找根，顺便把路径上所有节点直接连到根
        uf.parent[x] = uf.Find(uf.parent[x])
    }
    return uf.parent[x]
}

// Union - 合并 x 和 y 所在的集合
func (uf *UnionFind) Union(x, y int) {
    rootX := uf.Find(x)  // 找 x 的老大
    rootY := uf.Find(y)  // 找 y 的老大

    if rootX != rootY {
        // 让一个老大认另一个当老大
        uf.parent[rootX] = rootY
        uf.count--  // 两个集合合并成一个，总数-1
    }
}

func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }

    m, n := len(grid), len(grid[0])
    uf := NewUnionFind(m * n)

    // 第一步：统计陆地数量
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                uf.count++
            }
        }
    }

    // 第二步：合并相邻的陆地
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                id := i*n + j  // 二维坐标转一维编号

                // 只检查右边和下边（避免重复）
                // 检查右边
                if j+1 < n && grid[i][j+1] == '1' {
                    uf.Union(id, id+1)
                }
                // 检查下边
                if i+1 < m && grid[i+1][j] == '1' {
                    uf.Union(id, id+n)
                }
            }
        }
    }

    return uf.count
}
```

### JavaScript 实现

```javascript
class UnionFind {
    constructor(size) {
        this.parent = Array(size).fill(0).map((_, i) => i);
        this.count = 0;
    }

    find(x) {
        if (this.parent[x] !== x) {
            this.parent[x] = this.find(this.parent[x]);
        }
        return this.parent[x];
    }

    union(x, y) {
        const rootX = this.find(x);
        const rootY = this.find(y);
        if (rootX !== rootY) {
            this.parent[rootX] = rootY;
            this.count--;
        }
    }
}

function numIslands(grid) {
    if (!grid.length) return 0;

    const m = grid.length;
    const n = grid[0].length;
    const uf = new UnionFind(m * n);

    // 统计陆地数量
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (grid[i][j] === '1') {
                uf.count++;
            }
        }
    }

    // 合并相邻的陆地
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (grid[i][j] === '1') {
                const id = i * n + j;

                // 检查右边
                if (j + 1 < n && grid[i][j + 1] === '1') {
                    uf.union(id, id + 1);
                }
                // 检查下边
                if (i + 1 < m && grid[i + 1][j] === '1') {
                    uf.union(id, id + n);
                }
            }
        }
    }

    return uf.count;
}
```

---

## 复杂度分析

### 时间复杂度
- **O(m × n × α(m×n))**
  - m × n：遍历所有格子
  - α(m×n)：反阿克曼函数，几乎是常数（≈ 4）

### 空间复杂度
- **O(m × n)**：parent 数组

---

## 优化：按秩合并

```go
type UnionFind struct {
    parent []int
    rank   []int  // 树的高度
    count  int
}

func (uf *UnionFind) Union(x, y int) {
    rootX := uf.Find(x)
    rootY := uf.Find(y)

    if rootX != rootY {
        // 让矮的树指向高的树
        if uf.rank[rootX] < uf.rank[rootY] {
            uf.parent[rootX] = rootY
        } else if uf.rank[rootX] > uf.rank[rootY] {
            uf.parent[rootY] = rootX
        } else {
            uf.parent[rootY] = rootX
            uf.rank[rootX]++
        }
        uf.count--
    }
}
```

**优点：** 避免树过深，查找更快

---

## 总结

### 核心概念
1. **parent[i] == i** → i 是根节点
2. **Find：** 递归找根 + 路径压缩
3. **Union：** 先 Find 再合并根，count--

### 关键要点
- Union 修改的是**根节点**，不是原始参数
- 节点不会消失，只是改变指向
- 最终所有节点都指向同一个根

### 适用场景
- 连通性问题
- 动态集合合并
- 关系列表输入（比 DFS/BFS 更高效）
