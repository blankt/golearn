package main

import "testing"

func TestIsValidBST(t *testing.T) {
	root := &TreeNode{
		Val: 2,
	}
	root.Left = &TreeNode{
		Val: 1,
	}
	root.Right = &TreeNode{
		Val: 3,
	}

	result := isValidBST(root)
	if result != true {
		t.Fail()
	}
}
