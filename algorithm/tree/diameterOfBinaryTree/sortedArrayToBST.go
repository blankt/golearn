package main

func sortedArrayToBST(nums []int) *TreeNode {
	var (
		root      *TreeNode
		heightMap = make(map[*TreeNode]int)
	)

	for _, v := range nums {
		root = Insert(root, v, heightMap)
	}
	return root
}

// Insert 生成平衡二叉树
func Insert(node *TreeNode, data int, heightMap map[*TreeNode]int) *TreeNode {
	if node == nil {
		nodeNode := &TreeNode{
			Val: data,
		}
		heightMap[nodeNode] = 1
		return nodeNode
	}

	var newNode *TreeNode
	node.Right = Insert(node.Right, data, heightMap)
	balanceFactor := BalanceFactor(node, heightMap)
	if balanceFactor == -2 {
		if node.Right != nil && data > node.Right.Val {
			// 右旋
			newNode = LeftRotate1(node, heightMap)
		} else {
			// 左右旋
			newNode = RightLeftRotate1(node, heightMap)
		}
	}
	if newNode != nil {
		AdjustHeight1(newNode, heightMap)
		return newNode
	} else {
		AdjustHeight1(node, heightMap)
		return node
	}
}

func BalanceFactor(node *TreeNode, heightMap map[*TreeNode]int) int {
	left, right := 0, 0
	if node.Left != nil {
		left = heightMap[node.Left]
	}
	if node.Right != nil {
		right = heightMap[node.Right]
	}
	return left - right
}

func RightLeftRotate1(node *TreeNode, heightMap map[*TreeNode]int) *TreeNode {
	node.Right = RightRotate1(node.Right, heightMap)
	return LeftRotate1(node, heightMap)
}

func RightRotate1(node *TreeNode, heightMap map[*TreeNode]int) *TreeNode {
	index := node.Left
	node.Left = index.Right
	index.Right = node
	AdjustHeight1(node, heightMap)
	return index
}

func LeftRotate1(node *TreeNode, heightMap map[*TreeNode]int) *TreeNode {
	index := node.Right
	node.Right = index.Left
	index.Left = node
	AdjustHeight1(node, heightMap)
	return index
}

func AdjustHeight1(node *TreeNode, heightMap map[*TreeNode]int) {
	left, right := 0, 0
	if node.Left != nil {
		left = heightMap[node.Left]
	}
	if node.Right != nil {
		right = heightMap[node.Right]
	}

	if left > right {
		heightMap[node] = left + 1
	} else {
		heightMap[node] = right + 1
	}
}

// 方法2 序列是二叉搜索树的中序遍历的结果 如果能保证两边树高度平衡就是结果 所以取中间节点递归建立二叉搜索树保持平衡
func sortedArrayToBST2(nums []int) *TreeNode {
	return help(0, len(nums)-1, nums)
}

func help(left, right int, nums []int) *TreeNode {
	if left > right {
		return nil
	}
	middle := (left + right) / 2
	node := &TreeNode{
		Val: nums[middle],
	}
	node.Left = help(left, middle-1, nums)
	node.Right = help(middle+1, right, nums)
	return node
}
