package main

/*
*
冒泡排序
空间复杂度:o(1)
时间复杂度： 最好:o(n) 最坏: o(n^2)
稳定性: 稳定
*/
func insertion(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		preIndex := i - 1
		current := nums[i]
		for preIndex >= 0 && current < nums[preIndex] {
			nums[preIndex+1] = nums[preIndex]
			preIndex -= 1
		}
		nums[preIndex+1] = current
	}
	return nums
}
