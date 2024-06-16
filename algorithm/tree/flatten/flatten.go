package flatten

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1.前序遍历后再生成结果
func flatten(root *TreeNode) {
	nodes := make([]*TreeNode, 0)
	Pre(root, &nodes)
	result := root
	for _, v := range nodes {
		if v == root {
			continue
		}
		root.Right = v
		root.Left = nil
		root = root.Right
	}
	root.Right = nil
	root = result
	return
}

func Pre(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}

	*nodes = append(*nodes, root)
	Pre(root.Left, nodes)
	Pre(root.Right, nodes)
}

// 2.
func flatten2(root *TreeNode) {

}
