# BFS（广度优先搜索）详解

## 目录
- [核心思想](#核心思想)
- [算法原理](#算法原理)
- [执行过程演示](#执行过程演示)
- [代码实现](#代码实现)
- [复杂度分析](#复杂度分析)

---

## 核心思想

**"一层一层扩散，像水波纹一样"**

就像在地图上画圈：
- 找到一块陆地，先标记**所有相邻的陆地**
- 再标记**相邻的相邻**
- 一圈一圈往外扩展

---

## 算法原理

### 主要步骤

1. **遍历网格**：从左到右、从上到下扫描每个格子
2. **发现新岛屿**：遇到 `'1'` 且未访问 → 岛屿数量 +1
3. **标记整个岛屿**：用 BFS 迭代标记所有相连的陆地为"已访问"

### BFS 逻辑（使用队列）

```
BFS(i, j):
  创建队列，加入起点 (i, j)
  标记 (i, j) 为已访问

  while 队列不为空:
    取出队首元素 (x, y)

    检查四个方向 (上下左右):
      if 有效的陆地 且 未访问:
        标记为已访问
        加入队列
```

### 关键技巧

**为什么用队列？**
- 队列是 **FIFO**（先进先出）
- 确保**先探索的节点，它的邻居也先被探索**
- 实现"一层一层"的扩散效果

---

## 执行过程演示

### 示例网格

```
1 1 0
1 0 1
```

### 完整执行流程

**初始状态：**
```
count = 0

1 1 0
1 0 1
```

**第 1 步：扫描到 (0,0)**
```
grid[0][0] = '1' ✅ 是陆地
count = 1  // 发现新岛屿
开始 BFS(0, 0)
```

**BFS 迭代过程：**

```
BFS(0, 0):
  初始化队列：queue = [(0, 0)]
  标记 grid[0][0] = '0'

  网格变化：
  0 1 0
  1 0 1

  ===== 第 1 轮循环 =====
  取出 (0, 0)
  检查四个方向：

  1. 下 (1, 0)：grid[1][0] = '1' ✅
     标记 grid[1][0] = '0'
     加入队列：queue = [(1, 0)]

  2. 上 (-1, 0)：越界 ❌

  3. 右 (0, 1)：grid[0][1] = '1' ✅
     标记 grid[0][1] = '0'
     加入队列：queue = [(1, 0), (0, 1)]

  4. 左 (0, -1)：越界 ❌

  网格变化：
  0 0 0
  0 0 1

  ===== 第 2 轮循环 =====
  取出 (1, 0)
  检查四个方向：

  1. 下 (2, 0)：越界 ❌
  2. 上 (0, 0)：grid[0][0] = '0' ❌
  3. 右 (1, 1)：grid[1][1] = '0' ❌
  4. 左 (1, -1)：越界 ❌

  queue = [(0, 1)]

  ===== 第 3 轮循环 =====
  取出 (0, 1)
  检查四个方向：

  1. 下 (1, 1)：grid[1][1] = '0' ❌
  2. 上 (-1, 1)：越界 ❌
  3. 右 (0, 2)：grid[0][2] = '0' ❌
  4. 左 (0, 0)：grid[0][0] = '0' ❌

  queue = []  // 队列为空，BFS 结束
```

**BFS(0, 0) 完成！整个左边的岛屿已标记完毕**

**第 2 步：继续扫描到 (1,2)**
```
grid[1][2] = '1' ✅ 是陆地
count = 2  // 发现新岛屿
开始 BFS(1, 2)
```

**BFS 迭代过程：**
```
BFS(1, 2):
  初始化队列：queue = [(1, 2)]
  标记 grid[1][2] = '0'

  网格变化：
  0 0 0
  0 0 0

  ===== 第 1 轮循环 =====
  取出 (1, 2)
  检查四个方向：
    所有方向都是边界或水 ❌

  queue = []  // 队列为空，BFS 结束
```

**BFS(1, 2) 完成！**

**第 3 步：扫描完所有格子**
```
最终 count = 2  ✅
```

### 视觉化 BFS 层级

```
初始网格：
1 1 0
1 0 1

BFS(0,0) 的扩散层级：

第 0 层（起点）：
(0,0)

第 1 层（邻居）：
(1,0), (0,1)

标记过程（按层）：
层0: 0 1 0  →  层1: 0 0 0
     1 0 1           0 0 1

队列变化：
[(0,0)] → [(1,0),(0,1)] → [(0,1)] → []
```

### BFS vs DFS 路径对比

同样的网格，不同的探索顺序：

**DFS（深度优先）：**
```
(0,0) → (1,0) → 回退 → (0,1) → 完成
路径：先往深处走
```

**BFS（广度优先）：**
```
第0层: (0,0)
第1层: (1,0), (0,1)
路径：先探索同层的所有节点
```

---

## 代码实现

### Go 实现

```go
func numIslands(grid [][]byte) int {
    count := 0

    for i := range grid {
        for j := range grid[0] {
            if grid[i][j] == '1' {
                count++
                bfs(&grid, i, j)
            }
        }
    }
    return count
}

// BFS - 广度优先搜索
func bfs(grid *[][]byte, i, j int) {
    queue := [][]int{{i, j}}
    (*grid)[i][j] = '0'  // 标记起点

    // 四个方向：下、上、右、左
    directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]  // 取出队首

        x, y := current[0], current[1]

        // 检查四个方向
        for _, dir := range directions {
            nx := x + dir[0]
            ny := y + dir[1]

            // 如果是有效的陆地
            if nx >= 0 && ny >= 0 && nx < len(*grid) && ny < len((*grid)[0]) && (*grid)[nx][ny] == '1' {
                (*grid)[nx][ny] = '0'  // 标记
                queue = append(queue, []int{nx, ny})  // 加入队列
            }
        }
    }
}
```

### JavaScript 实现

```javascript
function numIslands(grid) {
    let count = 0;

    for (let i = 0; i < grid.length; i++) {
        for (let j = 0; j < grid[0].length; j++) {
            if (grid[i][j] === '1') {
                count++;
                bfs(grid, i, j);
            }
        }
    }
    return count;
}

function bfs(grid, i, j) {
    const queue = [[i, j]];
    grid[i][j] = '0';  // 标记起点

    const directions = [[1, 0], [-1, 0], [0, 1], [0, -1]];  // 下上右左

    while (queue.length > 0) {
        const [x, y] = queue.shift();  // 取出队首

        // 检查四个方向
        for (const [dx, dy] of directions) {
            const nx = x + dx;
            const ny = y + dy;

            // 如果是有效的陆地
            if (nx >= 0 && ny >= 0 && nx < grid.length && ny < grid[0].length && grid[nx][ny] === '1') {
                grid[nx][ny] = '0';  // 标记
                queue.push([nx, ny]);  // 加入队列
            }
        }
    }
}
```

### Python 实现

```python
from collections import deque

def numIslands(grid: List[List[str]]) -> int:
    count = 0

    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if grid[i][j] == '1':
                count += 1
                bfs(grid, i, j)

    return count

def bfs(grid: List[List[str]], i: int, j: int) -> None:
    queue = deque([(i, j)])
    grid[i][j] = '0'  # 标记起点

    directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]  # 下上右左

    while queue:
        x, y = queue.popleft()  # 取出队首

        # 检查四个方向
        for dx, dy in directions:
            nx, ny = x + dx, y + dy

            # 如果是有效的陆地
            if 0 <= nx < len(grid) and 0 <= ny < len(grid[0]) and grid[nx][ny] == '1':
                grid[nx][ny] = '0'  # 标记
                queue.append((nx, ny))  # 加入队列
```

---

## 复杂度分析

### 时间复杂度
- **O(m × n)**
  - 每个格子最多访问一次
  - 加入队列一次，取出队列一次

### 空间复杂度
- **O(min(m, n))**
  - 队列的最大长度取决于岛屿的"宽度"
  - 最坏情况：整个网格是一个岛屿，队列最大长度为 `min(m, n)`

**与 DFS 对比：**
- DFS 空间复杂度：O(m × n)（递归栈深度）
- BFS 空间复杂度：O(min(m, n))（队列大小）
- **BFS 更节省空间！** ✅

---

## BFS vs DFS

### 示例：蛇形岛屿

```
1 1 1 1 1
0 0 0 0 1
1 1 1 1 1
1 0 0 0 0
1 1 1 1 1
```

**DFS：**
- 递归栈深度：≈ 25（整条蛇的长度）
- 空间：O(m × n) = O(25)

**BFS：**
- 队列最大长度：≈ 2-3（每层最多几个节点）
- 空间：O(min(m, n)) ≈ O(5)

**BFS 更适合这种场景！** ✅

---

## 优化和注意事项

### 使用双端队列

**JavaScript/Python：** 使用 `deque`（双端队列）
```python
from collections import deque
queue = deque()
queue.append((i, j))  # O(1)
queue.popleft()       # O(1)
```

**不推荐：** 使用普通数组的 `shift()`
```javascript
queue.shift();  // O(n) - 性能差！
```

### 如果不能修改原数组

```go
func bfs(grid [][]byte, visited [][]bool, i, j int) {
    queue := [][]int{{i, j}}
    visited[i][j] = true

    directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        x, y := current[0], current[1]

        for _, dir := range directions {
            nx := x + dir[0]
            ny := y + dir[1]

            if nx >= 0 && ny >= 0 && nx < len(grid) && ny < len(grid[0]) &&
                grid[nx][ny] == '1' && !visited[nx][ny] {
                visited[nx][ny] = true
                queue = append(queue, []int{nx, ny})
            }
        }
    }
}
```

---

## 总结

### 优点
- **不会栈溢出**：使用迭代而非递归
- **空间效率更高**：队列通常比递归栈小
- **适合大规模数据**：可以处理 300×300 的网格
- **可扩展性强**：容易改成最短路径算法

### 缺点
- **代码稍复杂**：需要手动管理队列
- **稍难理解**：没有 DFS 的递归那么直观

### 适用场景
- 网格很大，担心栈溢出
- 需要层级信息（如求最短路径）
- 需要空间效率

### 相关题目
- 200\. Number of Islands（本题）
- 542\. 01 Matrix（求最短距离）
- 994\. Rotting Oranges（腐烂的橘子）
- 1162\. As Far from Land as Possible（离陆地最远的点）

---

## BFS 的其他应用

### 最短路径问题

BFS 天然适合求**无权图的最短路径**：

```go
func shortestPath(grid [][]int, start, end []int) int {
    queue := [][]int{{start[0], start[1], 0}}  // (x, y, distance)
    visited := make(map[string]bool)

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        x, y, dist := current[0], current[1], current[2]

        if x == end[0] && y == end[1] {
            return dist  // 找到最短路径！
        }

        // 探索邻居...
        for _, dir := range directions {
            nx, ny := x+dir[0], y+dir[1]
            key := fmt.Sprintf("%d,%d", nx, ny)

            if isValid(nx, ny) && !visited[key] {
                visited[key] = true
                queue = append(queue, []int{nx, ny, dist + 1})
            }
        }
    }

    return -1  // 无法到达
}
```

### 多源 BFS

从多个起点同时扩散：

```go
// 例如：腐烂的橘子问题
func orangesRotting(grid [][]int) int {
    queue := [][]int{}

    // 把所有初始腐烂的橘子加入队列
    for i := range grid {
        for j := range grid[0] {
            if grid[i][j] == 2 {
                queue = append(queue, []int{i, j, 0})
            }
        }
    }

    // 然后统一 BFS
    // ...
}
```

---

## 总结对比

| 特点 | DFS | BFS |
|------|-----|-----|
| **实现方式** | 递归 | 队列 |
| **探索顺序** | 深度优先 | 广度优先 |
| **代码复杂度** | 简单 | 中等 |
| **空间复杂度** | O(m×n) | O(min(m,n)) |
| **栈溢出风险** | 有 | 无 |
| **求最短路径** | 不适合 | 天然适合 |
| **适用场景** | 小规模网格、面试 | 大规模网格、最短路径 |
