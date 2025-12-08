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
				bfs(&grid, i, j)
			}
		}
	}
	return count

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
