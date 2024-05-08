package main

func findAnagrams(s string, p string) []int {
	result := make([]int, 0)
	needs := make(map[byte]int)
	windows := make(map[byte]int)
	for k := range p {
		needs[p[k]]++
	}

	i := 0
	j := 0
	valid := 0
	needShrink := false
	var needShrinkElem byte
	for ; j < len(s); j++ {
		nums := s[j]
		needTimes, ok := needs[nums]
		if !ok {
			needShrink = true
			needShrinkElem = nums
		}

		haveTimes := windows[nums]
		haveTimes++
		windows[nums] = haveTimes
		if haveTimes == needTimes {
			valid++
		}
		if valid == len(needs) {
			needShrink = true
			needShrinkElem = s[i]
			result = append(result, i)
		}
		if haveTimes > needTimes {
			needShrink = true
			needShrinkElem = nums
		}

		for needShrink {
			for i <= j {
				numsI := s[i]
				iNeedNum := needs[numsI]
				iHaveNum := windows[numsI]
				if iNeedNum != 0 && iHaveNum == iNeedNum {
					valid--
				}
				windows[numsI] = iHaveNum - 1
				i++
				if numsI == needShrinkElem {
					needShrink = false
					break
				}
			}
		}
	}
	return result
}

func findAnagrams1(s string, p string) []int {
	result := make([]int, 0)
	needs := make(map[byte]int)
	windows := make(map[byte]int)
	for k := range p {
		needs[p[k]]++
	}

	i, j := 0, 0
	needSize := 0
	for ; j < len(s); j++ {
		num := s[j]
		if needTimes, ok := needs[num]; ok {
			havaTimes := windows[num]
			havaTimes++
			if havaTimes == needTimes {
				needSize++
			}
			windows[num] = havaTimes
		}

		for needSize >= len(needs) {
			letter := s[i]
			if j-i+1 == len(p) {
				result = append(result, i)
			}

			if needTimes, ok := needs[letter]; ok {
				havaTimes := windows[letter]
				havaTimes--
				if havaTimes < needTimes {
					needSize--
				}
				windows[letter] = havaTimes
			}
			i++
		}
	}
	return result
}
