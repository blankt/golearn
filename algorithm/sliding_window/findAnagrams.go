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
	for ; j < len(s); j++ {
		nums := s[j]
		needTimes, ok := needs[nums]
		if !ok {
			continue
		}

		haveTimes := windows[nums]
		haveTimes++
		windows[nums] = haveTimes
		if haveTimes == needTimes {
			valid++
		}

		for valid == len(needs) {
			numsI := s[i]
			iNeedTime, ok := needs[numsI]
			if !ok {
				i++
				continue
			}

			iHaveTime := windows[numsI]
			iHaveTime--
			windows[numsI] = iHaveTime

			if iHaveTime < iNeedTime {
				valid--
				success := true
				for k, v := range needs {
					if v != windows[k] {
						success = false
						break
					}
				}
				if success {
					result = append(result, i)
				}
			}

			i++
		}

	}
	return result
}
