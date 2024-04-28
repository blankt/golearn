package main

/*
*
选择排序
空间复杂度:  o(1)
时间复杂度： 最好:o(nlogn) 最坏: o(n^2) 平均: o(nlogn)
稳定性: 不稳定
*/
func shell(nums []int) []int {
	gap := len(nums) / 2
	for gap > 0 {
		for i := gap; i < len(nums); i++ {
			preIndex := i - gap
			current := nums[i]
			for preIndex >= 0 && current < nums[preIndex] {
				nums[preIndex+gap] = nums[preIndex]
				preIndex -= gap
			}
			nums[preIndex+gap] = current
		}
		gap = gap / 2
	}
	return nums
}
