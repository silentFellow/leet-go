package leetcode

func minMaxDifference(num int) int {
	fnz, fnn := -1, -1

	split := make([]int, 0)
	for num != 0 {
		split = append([]int{num % 10}, split...)
		num /= 10
	}

	n := len(split)
	for i := range n {
		if fnn == -1 && split[i] != 9 {
			fnn = split[i]
		}

		if fnz == -1 && split[i] != 0 {
			fnz = split[i]
		}

		if fnz != -1 && fnn != -1 {
			break
		}
	}

	maxv, minv := make([]int, n), make([]int, n)

	for i, v := range split {
		if v == fnn {
			maxv[i] = 9
		} else {
			maxv[i] = v
		}
	}

	for i, v := range split {
		if v == fnz {
			minv[i] = 0
		} else {
			minv[i] = v
		}
	}

	getNumber := func(digits []int) int {
		num := 0
		for _, v := range digits {
			num = (num * 10) + v
		}

		return num
	}

	return getNumber(maxv) - getNumber(minv)
}
