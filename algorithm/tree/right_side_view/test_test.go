package main

import (
	"fmt"
	"testing"
)

func TestRightSideView(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 3,
	}
	root.Right.Right = &TreeNode{
		Val: 4,
	}
	root.Left.Right = &TreeNode{
		Val: 5,
	}

	result := rightSideView(root)
	if len(result) != 3 {
		t.Fail()
	}
	fmt.Println(result)
}
