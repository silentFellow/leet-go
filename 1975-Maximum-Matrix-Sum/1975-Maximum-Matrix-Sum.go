package leetcode

func maxMatrixSum(matrix [][]int) int64 {
	abs := func(v int) int {
		if v < 0 {
			return -v
		}

		return v
	}

	sum := 0
	containsZero := false
	negativeCount := 0
	minAbs := abs(matrix[0][0])
	for _, row := range matrix {
		for _, val := range row {
			if val == 0 {
				containsZero = true
			} else if val < 0 {
				negativeCount++
			}

			sum += abs(val)
			minAbs = min(minAbs, abs(val))
		}
	}

	if containsZero || negativeCount%2 == 0 {
		return int64(sum)
	} else {
		return int64(sum - 2*minAbs)
	}
}
