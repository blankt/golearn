package kth_smallest

import "testing"

func TestKthSmallest(t *testing.T) {
	root := &TreeNode{
		Val: 3,
	}
	root.Left = &TreeNode{
		Val: 1,
	}
	root.Right = &TreeNode{
		Val: 4,
	}
	root.Left.Right = &TreeNode{
		Val: 2,
	}
	if kthSmallest(root, 1) != 1 {
		t.Fail()
	}
}
