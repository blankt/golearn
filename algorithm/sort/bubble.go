package main

/*
*
冒泡排序
空间复杂度:o(1)
时间复杂度： 最好:o(n) 最坏: o(n^2)
稳定性: 稳定
*/
func bubbleSort(nums []int) []int {
	var flag = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = false
			}
		}
		if flag {
			break
		}
		flag = true
	}
	return nums
}
