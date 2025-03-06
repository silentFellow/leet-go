package leetcode

func findMissingAndRepeatedValues(grid [][]int) []int {
	arr := make([]int, len(grid)*len(grid))

	for _, row := range grid {
		for _, col := range row {
			arr[col-1]++
		}
	}

	ans := make([]int, 2)
	for i, v := range arr {
		if v == 2 {
			ans[0] = i + 1
		} else if v == 0 {
			ans[1] = i + 1
		}
	}

	return ans
}
