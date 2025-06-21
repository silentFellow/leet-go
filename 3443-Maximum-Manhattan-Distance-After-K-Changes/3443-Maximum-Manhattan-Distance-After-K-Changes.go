package leetcode

func maxDistance(s string, k int) int {
	netN, netE, result := 0, 0, 0

	abs := func(v int) int {
		if v < 0 {
			return -v
		}

		return v
	}

	for i, v := range s {
		switch v {
		case 'N':
			netN++
		case 'S':
			netN--
		case 'E':
			netE++
		case 'W':
			netE--
		}

		result = max(result, min(i+1, abs(netN)+abs(netE)+2*k))
	}

	return result
}
