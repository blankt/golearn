package avl_tree

import (
	"fmt"
	"testing"
)

func TestAVLTree_Insert(t *testing.T) {
	nums := []int{3, 2, 1, 4, 5, 6, 7, 10, 9, 8}
	avlTree := &AVLTree{}

	for _, v := range nums {
		avlTree.Insert(v)
	}
	fmt.Println(avlTree)
}

func TestSortTreeToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	//root := sortedArrayToBST(nums)
	root := sortedArrayToBST2(nums)
	fmt.Println(root)
}
