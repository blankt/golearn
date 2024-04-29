package main

func heap(nums []int) []int {
	// 建立大顶堆
	for i := len(nums) / 2; i >= 0; i-- {
		adjustHeap(nums, i, len(nums))
	}

	//交换顶点到最后并排序
	for j := len(nums) - 1; j >= 0; j-- {
		nums[0], nums[j] = nums[j], nums[0]
		adjustHeap(nums, 0, j)
	}
	return nums
}

func adjustHeap(nums []int, i, length int) {
	temp := nums[i]
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && nums[k] < nums[k+1] {
			k++
		}
		if temp < nums[k] {
			nums[i] = nums[k]
			i = k
		} else {
			break
		}
	}
	nums[i] = temp
}
