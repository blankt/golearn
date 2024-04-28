package main

func majorityElement(nums []int) int {
	max := nums[0]
	num := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != max {
			num--
			if num == 0 {
				num = 1
				max = nums[i]
			}
		} else {
			num++
		}
	}
	return max
}

func main() {

}
