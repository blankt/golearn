package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else {
			return false
		}
	})

	result := make([][]int, len(intervals))
	for k := range result {
		result[k] = make([]int, 2)
	}
	index := -1
	for i := 0; i < len(intervals); i++ {
		if index == -1 || intervals[i][0] > result[index][1] {
			index++
			result[index] = intervals[i]
		} else {
			if result[index][1] < intervals[i][1] {
				result[index][1] = intervals[i][1]
			}
		}
	}

	if index == len(result)-1 {
		return result
	} else {
		return result[:index+1]
	}
}
