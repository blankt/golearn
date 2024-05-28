package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	queue := list.List{}
	queue.Init()
	queue.PushBack(root)
	last := root
	for queue.Len() != 0 {
		node := queue.Front().Value.(*TreeNode)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}

		if last == node {
			result = append(result, node.Val)
			last = queue.Back().Value.(*TreeNode)
		}

		queue.Remove(queue.Front())
	}
	return result
}
