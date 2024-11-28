package leetcode

func average(salary []int) float64 {
	max, min, total := salary[0], salary[0], 0

	for _, val := range salary {
		if val > max {
			max = val
		} else if val < min {
			min = val
		}

		total += val
	}

	return float64(total-max-min) / float64(len(salary)-2)
}
