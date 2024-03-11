package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//随机生成10w个数字
	rand.Seed(time.Now().UnixNano())
	randomNumbers := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		randomNumbers[i] = rand.Intn(100000)
	}

	result, compareNum := GetTop100(randomNumbers, 100)
	fmt.Println(compareNum)

	otherResult := HeapTake100(randomNumbers)
	sort.Slice(otherResult, func(i, j int) bool {
		return otherResult[i] > otherResult[j]
	})
	failed := false
	i := 0
	for i < len(result) {
		if result[i] != otherResult[i] {
			failed = true
			break
		}
		i++
	}
	if failed {
		fmt.Println("获取前100个数失败")
	} else {
		fmt.Println("获取前100个数成功")
	}

}

type Node struct {
	Num   int
	IsMax bool
	index int
}

func GetTop100(nums []int, top int) ([]int, int) {
	if len(nums) <= top {
		return nums, 0
	}
	var (
		nodes      []*Node
		maxNum     = make([]int, 0, 100)
		compareNum = 0
	)

	nodes = generateBinaryTree(nums, &compareNum)
	for len(maxNum) < top {
		maxNum = append(maxNum, nodes[0].Num)
		nodes[0].IsMax = true
		findMaxIndex(nodes, nodes[0].index, &compareNum)
	}

	return maxNum, compareNum
}

func findMaxIndex(nodes []*Node, index int, compareNum *int) {
	defer func() {
		_err := recover()
		if _err != nil {
			fmt.Println(index)
		}
	}()
	for index > 0 {
		var (
			left, right, parent int
		)
		if index%2 == 0 {
			right = index
			left = index - 1
			parent = left / 2
		} else {
			left = index
			right = index + 1
			parent = index / 2
		}
		// 比较大小
		if nodes[left].IsMax {
			if !nodes[parent].IsMax && nodes[parent].Num >= nodes[right].Num {
				*compareNum++
				index = (index - 1) / 2
				continue
			}
			temp := nodes[right]
			nodes[right] = nodes[parent]
			nodes[parent] = temp
		} else if nodes[right].IsMax {
			if !nodes[parent].IsMax && nodes[parent].Num >= nodes[left].Num {
				*compareNum++
				index = (index - 1) / 2
				continue
			}
			temp := nodes[left]
			nodes[left] = nodes[parent]
			nodes[parent] = temp
		} else if nodes[left].Num > nodes[right].Num {
			if !nodes[parent].IsMax && nodes[parent].Num >= nodes[left].Num {
				*compareNum++
				index = (index - 1) / 2
				continue
			}
			temp := nodes[left]
			nodes[left] = nodes[parent]
			nodes[parent] = temp
		} else {
			if !nodes[parent].IsMax && nodes[parent].Num >= nodes[right].Num {
				*compareNum++
				index = (index - 1) / 2
				continue
			}
			temp := nodes[right]
			nodes[right] = nodes[parent]
			nodes[parent] = temp
		}

		*compareNum++
		index = (index - 1) / 2
	}
}

// 生成第一个全部遍历的二叉树
func generateBinaryTree(num []int, compareNum *int) []*Node {
	var (
		nodes    []*Node
		otherNum []int
	)
	basicNum := make([]int, 0, len(num)*2)

	exponent := int(math.Ceil(math.Log2(float64(len(num)))))
	if int(math.Exp2(float64(exponent))) == len(num) {
		basicNum = num
	} else {
		diff := int(math.Exp2(float64(exponent))) - len(num)
		otherNum = make([]int, 0, diff)
		for i := 0; i < diff; i++ {
			otherNum = append(otherNum, math.MinInt)
		}
		basicNum = append(num, otherNum...)
	}

	nodes = make([]*Node, 2*len(basicNum)-1)
	i := len(nodes) - len(basicNum)
	for _, v := range basicNum {
		var node = &Node{
			Num:   v,
			index: i,
		}
		nodes[i] = node
		i++
	}
	j := len(nodes) - 1
	index := len(nodes) - len(basicNum) - 1
	for j > 0 {
		if nodes[j].Num > nodes[j-1].Num {
			nodes[index] = nodes[j]
		} else {
			nodes[index] = nodes[j-1]
		}
		*compareNum++
		index--

		j -= 2
	}
	return nodes
}

func HeapTake100(nums []int) []int {
	result := make([]int, 100)
	for i := 0; i < 100; i++ {
		result[i] = nums[i]
	}

	buildHeap(result)
	for i := 100; i < len(nums); i++ {
		if result[0] < nums[i] {
			result[0] = nums[i]
			heapify(result, 0, len(result))
		}
	}
	return result
}

func buildHeap(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, i, n)
	}
}

func heapify(arr []int, i, n int) {
	left := 2*i + 1
	right := 2*i + 2
	max := i
	if left < n && arr[left] < arr[max] {
		max = left
	}
	if right < n && arr[right] < arr[max] {
		max = right
	}
	if max != i {
		arr[max], arr[i] = arr[i], arr[max]
		heapify(arr, max, n)
	}
}
