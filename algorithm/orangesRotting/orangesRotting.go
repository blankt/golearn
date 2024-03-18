package main

import (
	"container/list"
	"fmt"
	"sort"
)

//图广度优先遍历

func main() {
	//[[2,1,1],[1,1,1],[0,1,2]] 1,2
	//grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	grid := [][]int{{1}, {2}, {1}, {2}}
	result := orangesRotting(grid)
	fmt.Println(result)
}

func orangesRotting(grid [][]int) int {
	dr := []int{-1, 0, 1, 0}
	dc := []int{0, -1, 0, 1}
	queue := list.New()
	depthMap := make(map[int]int)

	line := len(grid)
	col := len(grid[0])
	for i := 0; i < line; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 2 {
				index := i*col + j
				queue.PushBack(index)
				depthMap[index] = 0
			}
		}
	}

	var result int
	for queue.Len() != 0 {
		index := queue.Front().Value.(int)
		i := index / col
		j := index % col
		queue.Remove(queue.Front())
		for k := 0; k < 4; k++ {
			ni := i + dr[k]
			nj := j + dc[k]
			if ni >= 0 && ni < line && nj >= 0 && nj < col && grid[ni][nj] == 1 {
				nIndex := ni*col + nj
				queue.PushBack(nIndex)
				grid[ni][nj] = 2
				depthMap[nIndex] = depthMap[index] + 1
				result = depthMap[nIndex]
			}
		}
	}

	for i := 0; i < line; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return result
}

func orangesRotting1(grid [][]int) int {
	line := len(grid)
	col := len(grid[0])
	max := 0
	for i := 0; i < line; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 2 {
				var depth int
				grid[i][j] = 1
				depth = dfs(grid, i, j)
				if depth > max {
					max = depth
				}
			}
		}
	}

	for i := 0; i < line; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] != 0 {
				return -1
			}
		}
	}
	if max > 0 {
		return max - 1
	}
	return max
}

func dfs(grid [][]int, i, j int) int {
	line := len(grid)
	col := len(grid[0])
	if i < 0 || j < 0 || i >= line || j >= col {
		return 0
	}

	if grid[i][j] == 0 || grid[i][j] == 2 {
		return 0
	}

	grid[i][j] = 0
	x1 := dfs(grid, i+1, j)
	x2 := dfs(grid, i-1, j)
	x3 := dfs(grid, i, j+1)
	x4 := dfs(grid, i, j-1)
	nums := []int{x1, x2, x3, x4}
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] < nums[j] {
			return false
		} else {
			return true
		}
	})
	return nums[0] + 1
}
