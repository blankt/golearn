package kth_smallest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	index, result := 0, 0
	middleSort(root, k, &index, &result)
	return result
}

func middleSort(root *TreeNode, k int, index, result *int) {
	if root == nil || *index > k {
		return
	}

	middleSort(root.Left, k, index, result)
	*index++
	if *index == k {
		*result = root.Val
	}
	middleSort(root.Right, k, index, result)
}
