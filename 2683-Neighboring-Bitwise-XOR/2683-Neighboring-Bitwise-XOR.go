package leetcode

func doesValidArrayExist(derived []int) bool {
	first, last := 0, 0

	for _, val := range derived {
		if val == 1 {
			if last == 0 {
				last = 1
			} else {
				last = 0
			}
		}
	}

	return first == last
}
