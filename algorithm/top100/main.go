package main

import (
	"fmt"
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
