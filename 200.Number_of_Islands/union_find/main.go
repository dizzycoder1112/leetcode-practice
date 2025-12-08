package main

import (
	"fmt"

	"number_of_islands"
)

func main() {
	testCases := testcase.GetTestCases()
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %d\n", i+1, numIslands(testcase.CopyGrid(tc)))
	}
}

func numIslands(grid [][]byte) int {
	return numIslandsByUnionFind(grid)
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
