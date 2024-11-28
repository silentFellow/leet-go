package leetcode

func checkOnesSegment(s string) bool {
	index := 0

	// find first one
	for i, val := range s {
		if val == '1' {
			index = i
			break
		}
	}

	// loop till 0
	for i := index; i < len(s); i++ {
		if s[i] == '0' {
			index = i
			break
		}

		// custom case when no 0 is found. i.e: all are 1's
		if i == len(s)-1 {
			return true
		}
	}

	// if another 1 exists return false
	for i := index; i < len(s); i++ {
		if s[i] == '1' {
			return false
		}
	}

	return true
}
