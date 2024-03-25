package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var max int
	depth(root, &max)
	return max
}

func depth(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := depth(root.Left, max)
	right := depth(root.Right, max)
	if left+right > *max {
		*max = left + right
	}

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
