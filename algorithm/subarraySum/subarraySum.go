package main

import "fmt"

/*
560题
*/
func subarraySum(nums []int, k int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				result++
			}
		}
	}
	return result
}

// 前缀和
func subarraySum2(nums []int, k int) int {
	preNum := make([]int, len(nums)+1)
	preNum[0] = 0
	count := 0
	for index, v := range nums {
		preNum[index+1] = v + preNum[index]
	}

	for i := len(preNum) - 1; i > 0; i-- {
		for j := 1; j <= i; j++ {
			if preNum[i]-preNum[j-1] == k {
				count++
			}
		}
	}

	return count
}

// 由于不关心序列是那些只关心结果 可以用map + 前缀和来解决
func subarraySum3(nums []int, k int) int {
	count := 0
	preMap := make(map[int]int)
	preMap[0] = 1
	pre := 0

	for i := 0; i < len(nums); i++ {
		pre += nums[i]

		if preMap[pre-k] != 0 {
			count += preMap[pre-k]
		}
		preMap[pre]++
	}
	return count
}

func main() {
	nums := []int{1, -1, 0}
	fmt.Println(subarraySum(nums, 0))

	fmt.Println(subarraySum2(nums, 0))

	fmt.Println(subarraySum3(nums, 0))
}
