package build_tree

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}
	root := buildTree(preOrder, inOrder)
	if root.Val != 3 {
		t.Fatal("error")
	}
	fmt.Println(root)
}
