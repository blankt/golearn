package level_order

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
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
