package main

import "log"

func main() {
	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	result := numIslands(grid)
	if result != 3 {
		log.Fatal(result)
	}
}

// 图深度优先遍历
func numIslands(grid [][]byte) int {
	line := len(grid)
	col := len(grid[0])
	result := 0

	for i := 0; i < line; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				result++
				dfs(grid, i, j)
			}
		}
	}
	return result
}

func dfs(grid [][]byte, i, j int) {
	line := len(grid)
	col := len(grid[0])
	if i < 0 || i >= line || j < 0 || j >= col {
		return
	}

	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'

	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}
