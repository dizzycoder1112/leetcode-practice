package testcase

func GetTestCases() [][][]byte {
	return [][][]byte{
		{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}, // Expected: 1
		{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}, // Expected: 3
	}
}

func CopyGrid(grid [][]byte) [][]byte {
	cp := make([][]byte, len(grid))
	for i := range grid {
		cp[i] = make([]byte, len(grid[i]))
		copy(cp[i], grid[i])
	}
	return cp
}
