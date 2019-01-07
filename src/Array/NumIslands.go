package Array

func numIslands(grid [][]byte) int {
	rowLen := len(grid)
	if rowLen == 0 {
		return 0
	}
	colLen := len(grid[0])
	if colLen == 0 {
		return 0
	}

	var res int
	for i:=0;i<rowLen;i++ {
		for j:=0;j<colLen;j++ {
			if grid[i][j] == 1 {
				res++
				dfs(grid, i, j, rowLen, colLen)
			}
		}
	}

	return res
}

func dfs(grid [][]byte, i, j, rowLen, colLen int) {
	grid[i][j] = 2
	if i > 0 && grid[i-1][j] == 1 {
		dfs(grid, i-1, j, rowLen, colLen)
	}
	if i < rowLen-1 && grid[i+1][j] == 1 {
		dfs(grid, i+1, j, rowLen, colLen)
	}
	if j > 0 && grid[i][j-1] == 1 {
		dfs(grid, i, j-1, rowLen, colLen)
	}
	if j < colLen-1 && grid[i][j+1] == 1 {
		dfs(grid, i, j+1, rowLen, colLen)
	}
}