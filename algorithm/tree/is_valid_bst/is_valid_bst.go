package is_valid_bst

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	result := true
	min := math.MinInt
	middleSort(root, &min, &result)
	return result
}

func middleSort(root *TreeNode, max *int, valid *bool) {
	if root == nil {
		return
	}
	if !*valid {
		return
	}
	middleSort(root.Left, max, valid)
	if root.Val <= *max {
		*valid = false
	}
	*max = root.Val
	middleSort(root.Right, max, valid)
}
