package leetcode

func maxDiff(num int) int {
	split := make([]int, 0)
	for num != 0 {
		split = append([]int{num % 10}, split...)
		num /= 10
	}

	n := len(split)
	fnn, fnz := -1, -1
	for _, v := range split {
		if fnn == -1 && v != 9 {
			fnn = v
		}

		if split[0] != 1 {
			fnz = split[0]
		} else {
			for i := 1; i < n; i++ {
				if split[i] != 0 && split[i] != 1 {
					fnz = split[i]
					break
				}
			}
		}

		if fnn != -1 && fnz != -1 {
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

		if fnz != -1 && v == fnz {
			if fnz == split[0] {
				minv[i] = 1
			} else {
				minv[i] = 0
			}
		} else {
			minv[i] = v
		}
	}

	arrayToDigit := func(arr []int) int {
		num := 0
		for _, v := range arr {
			num = (num * 10) + v
		}

		return num
	}

	return arrayToDigit(maxv) - arrayToDigit(minv)
}
