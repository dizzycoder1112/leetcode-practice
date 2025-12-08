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

func dfs(grid *[][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) || (*grid)[i][j] == '0' {
		return
	}

	(*grid)[i][j] = '0'

	dfs(grid, i+1, j) // down
	dfs(grid, i-1, j) // up
	dfs(grid, i, j+1) // right
	dfs(grid, i, j-1) // left
}
