package main

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)

	k := 0
	for k < len(nums) {
		i := 0
		j := i + 1
		for i < len(nums)-k && j < len(nums)-k {
			if nums[i] > nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
				i = j
				j++
			} else {
				i++
				j++
			}
		}
		k++
	}

	k = 0
	targetMap := make(map[int]struct{})
	for k < len(nums) {
		num := nums[k]
		_, ok := targetMap[num]
		if ok {
			k++
			continue
		} else {
			targetMap[num] = struct{}{}
		}
		target := -num
		i := k + 1
		j := len(nums) - 1
		for i < j {
			if nums[i]+nums[j] == target {
				singleResult := make([]int, 3)
				singleResult[0] = num
				singleResult[1] = nums[i]
				singleResult[2] = nums[j]
				result = append(result, singleResult)
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				i++
				j--
			} else if nums[i]+nums[j] > target {
				j--
			} else {
				i++
			}
		}
		k++
	}

	return result
}
