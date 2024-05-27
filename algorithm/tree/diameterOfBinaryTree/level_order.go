package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	front := 0
	last := root
	queue = append(queue, root)
	singleResult := make([]int, 0)
	for front < len(queue) {
		node := queue[front]
		singleResult = append(singleResult, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		if last == node {
			result = append(result, singleResult)
			singleResult = []int{}
			last = queue[len(queue)-1]
		}
		front++
	}
	return result
}
