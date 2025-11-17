package main

import "fmt"

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
}

func numIslands(grid [][]byte) int {
	count := 0

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				count++
				dfs(&grid, i, j)
			}
		}
	}
	return count

}

// Depth First Search
func dfs(grid *[][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) || (*grid)[i][j] == '0' {
		return
	}

	(*grid)[i][j] = '0'

	dfs(grid, i+1, j) //down
	dfs(grid, i-1, j) //up
	dfs(grid, i, j+1) //right
	dfs(grid, i, j-1) //left
}

// Breadth First Search
func bfs(grid *[][]byte, i, j int) {
	queue := [][]int{
		{i, j},
	}
	(*grid)[i][j] = '0'
	direction := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, v := range direction {
			x := current[0]
			y := current[1]
			dx := v[0]
			dy := v[1]
			xPrime := x + dx
			yPrime := y + dy

			if xPrime >= 0 && yPrime >= 0 && xPrime < len(*grid) && yPrime < len((*grid)[0]) && (*grid)[xPrime][yPrime] == '1' {
				(*grid)[xPrime][yPrime] = '0'
				queue = append(queue, []int{xPrime, yPrime})
			}
		}
	}
}

type UnionFind struct {
	parent []int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		count:  0,
	}

	for i := range n {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		uf.parent[rootX] = rootY
		uf.count--
	}
}

func numIslandsByUnionFind(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	uf := NewUnionFind(m * n)

	// 第一步：统计陆地数量
	for i := range m {
		for j := range n {
			if grid[i][j] == '1' {
				uf.count++
			}
		}
	}

	// 第二步：合并相邻的陆地
	for i := range m {
		for j := range n {
			if grid[i][j] == '1' {
				id := i*n + j // 二维坐标转一维编号

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
