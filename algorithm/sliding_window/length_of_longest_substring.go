package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	max := 0
	subStrMap := make(map[uint8]int)
	result := make([]uint8, 0, len(s))

	for i := 0; i < len(s); i++ {
		result = append(result, s[i])
		times, _ := subStrMap[s[i]]
		times++
		subStrMap[s[i]] = times

		if times > 1 {
			j := 0
			for ; j < len(result); j++ {
				if result[j] == s[i] {
					subStrMap[s[i]] = 1
					break
				} else {
					delete(subStrMap, result[j])
				}
			}
			result = result[j+1:]
		}

		if len(result) > max {
			max = len(result)
		}
	}
	return max
}

func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	max := 0
	subStrMap := make(map[uint8]int)
	rk := -1

	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(subStrMap, s[i-1])
		}

		for rk+1 < len(s) && subStrMap[s[rk+1]] == 0 {
			subStrMap[s[rk+1]] = 1
			rk++
		}

		if max < rk-i+1 {
			max = rk - i + 1
		}
	}
	return max
}
