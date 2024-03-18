package main

import "log"

func main() {
	prerequisites := [][]int{{0, 1}}
	result := canFinish(2, prerequisites)
	if result != true {
		log.Fatal("result error")
	}
}

// 转化问题为判断图里面有无环 先建立图 再用深度优先遍历判断是否有环

func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)
	havaCycle := false

	graph := buildGraph(numCourses, prerequisites)

	for i := 0; i < numCourses; i++ {
		traverse(graph, &visited, &onPath, &havaCycle, i)
	}
	return !havaCycle
}

func traverse(graph [][]int, visited, onPath *[]bool, havaCycle *bool, index int) {
	if (*onPath)[index] == true {
		*havaCycle = true
	}

	if (*visited)[index] || *havaCycle {
		return
	}

	(*visited)[index] = true
	(*onPath)[index] = true

	for _, v := range graph[index] {
		traverse(graph, visited, onPath, havaCycle, v)
	}

	(*onPath)[index] = false
}

func buildGraph(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)
	for i := range graph {
		graph[i] = make([]int, 0)
	}

	for i := range prerequisites {
		from, to := prerequisites[i][0], prerequisites[i][1]
		graph[from] = append(graph[from], to)
	}
	return graph
}
