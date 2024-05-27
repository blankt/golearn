package main

/*
以下是生成平衡二叉树
*/

type AVLTreeNode struct {
	Val    int
	Left   *AVLTreeNode
	Right  *AVLTreeNode
	Height int
}

type AVLTree struct {
	root *AVLTreeNode
}

func (root *AVLTree) Insert(data int) {
	root.root = root.root.InsertNode(data)
}

func (node *AVLTreeNode) InsertNode(data int) *AVLTreeNode {
	if node == nil {
		return &AVLTreeNode{
			Val:    data,
			Height: 1,
		}
	}

	if node.Val == data {
		return node
	}

	var newNode *AVLTreeNode

	if data > node.Val {
		node.Right = node.Right.InsertNode(data)
		balanceFactor := node.BalanceFactor()
		if balanceFactor == -2 {
			if node.Right != nil && data > node.Right.Val {
				// 右旋
				newNode = LeftRotate(node)
			} else {
				// 左右旋
				newNode = RightLeftRotate(node)
			}
		}

	} else {
		node.Left = node.Left.InsertNode(data)
		balanceFactor := node.BalanceFactor()
		if balanceFactor == 2 {
			if node.Right != nil && data > node.Right.Val {
				// 右旋
				newNode = LeftRightRotate(node)
			} else {
				// 左右旋
				newNode = RightRotate(node)
			}
		}
	}

	if newNode != nil {
		AdjustHeight(newNode)
		return newNode
	} else {
		AdjustHeight(node)
		return node
	}
}

func (node *AVLTreeNode) BalanceFactor() int {
	left, right := 0, 0
	if node.Left != nil {
		left = node.Left.Height
	}
	if node.Right != nil {
		right = node.Right.Height
	}
	return left - right
}

func RightLeftRotate(node *AVLTreeNode) *AVLTreeNode {
	node.Right = RightRotate(node.Right)
	return LeftRotate(node)
}

func LeftRightRotate(node *AVLTreeNode) *AVLTreeNode {
	node.Left = LeftRotate(node.Left)
	return RightRotate(node)
}

func RightRotate(node *AVLTreeNode) *AVLTreeNode {
	index := node.Left
	node.Left = index.Right
	index.Right = node
	AdjustHeight(node)
	return index
}

func LeftRotate(node *AVLTreeNode) *AVLTreeNode {
	index := node.Right
	node.Right = index.Left
	index.Left = node
	AdjustHeight(node)
	return index
}

func AdjustHeight(node *AVLTreeNode) {
	left, right := 0, 0
	if node.Left != nil {
		left = node.Left.Height
	}
	if node.Right != nil {
		right = node.Right.Height
	}

	if left > right {
		node.Height = left + 1
	} else {
		node.Height = right + 1
	}
}
