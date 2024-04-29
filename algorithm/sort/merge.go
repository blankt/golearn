package main

/*
*
选择排序
空间复杂度:  o(n)
时间复杂度： 最好:o(nlogn) 最坏: o(nlogn) 平均: o(nlogn)
稳定性: 稳定
*/
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	middle := len(nums) / 2
	left := nums[0:middle]
	right := nums[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	result := make([]int, 0)
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	if i != len(left) {
		result = append(result, left[i:]...)
	}
	if j != len(right) {
		result = append(result, right[j:]...)
	}
	return result
}
