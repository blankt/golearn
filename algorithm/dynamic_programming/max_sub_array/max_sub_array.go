package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	dp[0] = nums[0]
	result := nums[0]

	for i := 1; i < length; i++ {
		if dp[i-1] <= 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if result < dp[i] {
			result = dp[i]
		}
	}
	return result
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray1(nums []int) int {
	max := math.MinInt
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if max < sum {
			max = sum
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if max < sum {
				max = sum
			}
		}
	}
	return max
}
