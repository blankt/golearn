package main

/*
*
快速排序
空间复杂度:  o(logn)
时间复杂度： 最好:o(nlogn) 最坏: o(n^2) 平均: o(nlogn)
稳定性: 不稳定
*/
func quick(nums []int, low, high int) []int {
	if low < high {
		pivot := partition(nums, low, high)
		quick(nums, low, pivot-1)
		quick(nums, pivot+1, high)
	}
	return nums
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]
	index := low
	for i := low; i < high; i++ {
		if nums[i] <= pivot {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[index], nums[high] = nums[high], nums[index]
	return index
}
