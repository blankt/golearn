package main

func rotate(nums []int, k int) {
	result := make([]int, len(nums))
	for index, v := range nums {
		step := (index + k) % len(nums)
		result[step] = v
	}
	for index := range nums {
		nums[index] = result[index]
	}
}

/*
*
相当于往前走n步，那么后面n个数会跑到前面去。
1.先把数组反转
2.分割出要跑到前面去的n个数
3.再把前n个数和剩下的数据分别反转
4.得到最终结果
*/
func rotate2(nums []int, k int) {
	i := 0
	j := len(nums) - 1
	k = k % len(nums)
	reverse(nums, i, j)
	reverse(nums, 0, k-1)
	reverse(nums, k, j)
}

func reverse(nums []int, left, right int) {
	for left < right {
		temp := nums[left]
		nums[left] = nums[right]
		nums[right] = temp
		left++
		right--
	}
}
