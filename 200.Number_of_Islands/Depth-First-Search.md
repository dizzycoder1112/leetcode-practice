# DFS（深度优先搜索）详解

## 目录
- [核心思想](#核心思想)
- [算法原理](#算法原理)
- [执行过程演示](#执行过程演示)
- [代码实现](#代码实现)
- [复杂度分析](#复杂度分析)

---

## 核心思想

**"一条路走到底，走不通了再回头"**

就像在迷宫里探险：
- 找到一块陆地，就沿着这块陆地**一直往深处走**
- 遇到边界或水就**回退**，换个方向继续
- 直到把整个岛屿都探索完

---

## 算法原理

### 主要步骤

1. **遍历网格**：从左到右、从上到下扫描每个格子
2. **发现新岛屿**：遇到 `'1'` 且未访问 → 岛屿数量 +1
3. **标记整个岛屿**：用 DFS 递归标记所有相连的陆地为"已访问"

### 递归逻辑

```
DFS(i, j):
  如果 (i, j) 越界 或 是水 或 已访问：
    return

  标记 (i, j) 为已访问

  递归探索四个方向：
    DFS(i+1, j)  // 下
    DFS(i-1, j)  // 上
    DFS(i, j+1)  // 右
    DFS(i, j-1)  // 左
```

### 关键技巧

**如何标记"已访问"？**
- 方案 1：使用额外的 `visited` 数组
- 方案 2：直接把 `'1'` 改成 `'0'`（原地修改，节省空间）✅

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
开始 DFS(0, 0)
```

**DFS 递归过程：**

```
DFS(0, 0):
  标记 grid[0][0] = '0'

  网格变化：
  0 1 0
  1 0 1

  递归探索四个方向：

  1. DFS(1, 0) - 下
     grid[1][0] = '1' ✅
     标记 grid[1][0] = '0'

     网格变化：
     0 1 0
     0 0 1

     递归探索：
       DFS(2, 0) - 下：越界，return
       DFS(0, 0) - 上：grid[0][0] = '0'，return
       DFS(1, 1) - 右：grid[1][1] = '0'，return
       DFS(1,-1) - 左：越界，return

  2. DFS(-1, 0) - 上：越界，return

  3. DFS(0, 1) - 右
     grid[0][1] = '1' ✅
     标记 grid[0][1] = '0'

     网格变化：
     0 0 0
     0 0 1

     递归探索：
       DFS(1, 1) - 下：grid[1][1] = '0'，return
       DFS(-1, 1) - 上：越界，return
       DFS(0, 2) - 右：grid[0][2] = '0'，return
       DFS(0, 0) - 左：grid[0][0] = '0'，return

  4. DFS(0, -1) - 左：越界，return
```

**DFS(0, 0) 完成！整个左边的岛屿已标记完毕**

**第 2 步：继续扫描到 (1,2)**
```
grid[1][2] = '1' ✅ 是陆地
count = 2  // 发现新岛屿
开始 DFS(1, 2)
```

**DFS 递归过程：**
```
DFS(1, 2):
  标记 grid[1][2] = '0'

  网格变化：
  0 0 0
  0 0 0

  递归探索四个方向：
    DFS(2, 2) - 下：越界，return
    DFS(0, 2) - 上：grid[0][2] = '0'，return
    DFS(1, 3) - 右：越界，return
    DFS(1, 1) - 左：grid[1][1] = '0'，return
```

**DFS(1, 2) 完成！**

**第 3 步：扫描完所有格子**
```
最终 count = 2  ✅
所有格子都被访问过
```

### 视觉化 DFS 路径

```
初始网格：
1 1 0
1 0 1

DFS(0,0) 的探索路径：
(0,0) → (1,0) → 回退 → (0,1) → 完成

标记过程：
1 1 0  →  0 1 0  →  0 0 0
1 0 1     0 0 1     0 0 1

DFS(1,2) 的探索路径：
(1,2) → 四周都是边界或水 → 完成

最终：
0 0 0
0 0 0
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
                dfs(&grid, i, j)  // 标记整个岛屿
            }
        }
    }
    return count
}

// DFS - 深度优先搜索
func dfs(grid *[][]byte, i, j int) {
    // 终止条件：越界、是水、已访问
    if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) || (*grid)[i][j] == '0' {
        return
    }

    // 标记为已访问（把陆地变成水）
    (*grid)[i][j] = '0'

    // 递归探索四个方向（深度优先）
    dfs(grid, i+1, j)  // 下
    dfs(grid, i-1, j)  // 上
    dfs(grid, i, j+1)  // 右
    dfs(grid, i, j-1)  // 左
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
                dfs(grid, i, j);
            }
        }
    }
    return count;
}

function dfs(grid, i, j) {
    // 终止条件：越界、是水、已访问
    if (i < 0 || j < 0 || i >= grid.length || j >= grid[0].length || grid[i][j] === '0') {
        return;
    }

    // 标记为已访问（把陆地变成水）
    grid[i][j] = '0';

    // 递归探索四个方向（深度优先）
    dfs(grid, i + 1, j);  // 下
    dfs(grid, i - 1, j);  // 上
    dfs(grid, i, j + 1);  // 右
    dfs(grid, i, j - 1);  // 左
}
```

### Python 实现

```python
def numIslands(grid: List[List[str]]) -> int:
    count = 0

    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if grid[i][j] == '1':
                count += 1
                dfs(grid, i, j)

    return count

def dfs(grid: List[List[str]], i: int, j: int) -> None:
    # 终止条件：越界、是水、已访问
    if i < 0 or j < 0 or i >= len(grid) or j >= len(grid[0]) or grid[i][j] == '0':
        return

    # 标记为已访问
    grid[i][j] = '0'

    # 递归探索四个方向
    dfs(grid, i + 1, j)  # 下
    dfs(grid, i - 1, j)  # 上
    dfs(grid, i, j + 1)  # 右
    dfs(grid, i, j - 1)  # 左
```

---

## 复杂度分析

### 时间复杂度
- **O(m × n)**
  - m × n：最坏情况下，所有格子都是陆地，每个格子访问一次

### 空间复杂度
- **O(m × n)**
  - 递归调用栈的最大深度
  - 最坏情况：所有格子形成一条蛇形的岛屿，递归深度 = m × n

---

## 优化和注意事项

### 避免栈溢出

**问题：** 对于非常大的网格（如 300×300），递归深度可能导致栈溢出

**解决方案：**
1. 使用 BFS（迭代方式，用队列）
2. 增加栈大小（语言相关）
3. 手动用栈模拟递归

### 如果不能修改原数组

```go
func numIslands(grid [][]byte) int {
    m, n := len(grid), len(grid[0])
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }

    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' && !visited[i][j] {
                count++
                dfs(grid, visited, i, j)
            }
        }
    }
    return count
}

func dfs(grid [][]byte, visited [][]bool, i, j int) {
    if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) ||
        grid[i][j] == '0' || visited[i][j] {
        return
    }

    visited[i][j] = true

    dfs(grid, visited, i+1, j)
    dfs(grid, visited, i-1, j)
    dfs(grid, visited, i, j+1)
    dfs(grid, visited, i, j-1)
}
```

---

## 总结

### 优点
- **代码简洁**：只需要几行递归代码
- **容易理解**：递归逻辑直观
- **无需额外数据结构**：可以原地修改

### 缺点
- **可能栈溢出**：递归深度取决于岛屿大小
- **不适合大规模数据**：300×300 的网格可能有问题

### 适用场景
- 网格不是特别大
- 面试时的首选方案（代码最简洁）
- 需要快速实现的情况

### 相关题目
- 695\. Max Area of Island（求最大岛屿面积）
- 733\. Flood Fill（图像填充）
- 130\. Surrounded Regions（被围绕的区域）
- 417\. Pacific Atlantic Water Flow（太平洋大西洋水流）
