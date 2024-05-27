package main

import "fmt"

func main() {
	//nums := []int{3, 2, 1, 4, 5, 6, 7, 10, 9, 8}
	//avlTree := &AVLTree{}
	//
	//for _, v := range nums {
	//	avlTree.Insert(v)
	//}

	//nums := []int{-10, -3, 0, 5, 9}
	////root := sortedArrayToBST(nums)
	//root := sortedArrayToBST2(nums)
	//fmt.Println(root)

	//root := &TreeNode{
	//	Val: 3,
	//}
	//root.Left = &TreeNode{
	//	Val: 9,
	//}
	//root.Right = &TreeNode{
	//	Val: 20,
	//}
	//root.Right.Left = &TreeNode{
	//	Val: 15,
	//}
	//root.Right.Right = &TreeNode{
	//	Val: 7,
	//}
	result := levelOrder(nil)
	fmt.Println(result)
}
