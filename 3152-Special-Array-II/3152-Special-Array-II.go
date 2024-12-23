package leetcode

func isArraySpecial(nums []int, queries [][]int) []bool {
	// the following snippet of code seperates the nums into individual subarray
	// till no parity breaks arrive into a sequence
	subArray := make([][]int, 0)

	array := []int{0}
	for i := 1; i < len(nums); i++ {
		if nums[i-1]%2 == nums[i]%2 {
			subArray = append(subArray, array)
			array = []int{i}
			continue
		}

		array = append(array, i)
	}
	if len(array) != 0 {
		subArray = append(subArray, array)
	}

	// the following create a binary search
	// identifies which subArray contains from and to in each queries
	// if both in same then no parity breaks, meaning true else false
	ans := make([]bool, len(queries))
	for i, val := range queries {
		si, ei := 0, 0

		// find start index => si
		l, r := 0, len(subArray)-1
		for l <= r {
			m := (l + r) / 2

			if val[0] >= subArray[m][0] && val[0] <= subArray[m][len(subArray[m])-1] {
				si = m
				break
			} else if val[0] < subArray[m][0] {
				r = m - 1
			} else {
				l = m + 1
			}
		}

		// find end index => ei
		l, r = l, len(subArray)-1
		for l <= r {
			m := (l + r) / 2

			if val[1] >= subArray[m][0] && val[1] <= subArray[m][len(subArray[m])-1] {
				ei = m
				break
			} else if val[1] < subArray[m][0] {
				r = m - 1
			} else {
				l = m + 1
			}
		}

		ans[i] = (si == ei)
	}

	return ans
}

// NOTE: generate a helper function for binary search, if needed for clean code
