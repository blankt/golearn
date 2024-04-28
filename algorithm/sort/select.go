package main

/*
*
选择排序
空间复杂度:  o(1)
时间复杂度： 最好:o(n^2) 最坏: o(n^2)
稳定性: 不稳定
*/
func selectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
	return nums
}
