package leetcode

func firstCompleteIndex(arr []int, mat [][]int) int {
	prefixRowSum := make([]int, len(mat))       // to calculate prefix sum for each row
	prefixColumnSum := make([]int, len(mat[0])) // to calculate prefix sum for each column
	valueRowCol := make(map[int][2]int)         // map a value with its row and column

	for i, row := range mat {
		for j, val := range row {
			prefixRowSum[i] += val
			prefixColumnSum[j] += val
			valueRowCol[val] = [2]int{i, j}
		}
	}

	for i, val := range arr {
		rowCol := valueRowCol[val]        // find the value's row and column value from the map
		prefixRowSum[rowCol[0]] -= val    // reduce row prefix sum by the value
		prefixColumnSum[rowCol[1]] -= val // reduce column prefix sum by the value

		// if either prefix becomes zero that it is fully painted
		if prefixRowSum[rowCol[0]] == 0 || prefixColumnSum[rowCol[1]] == 0 {
			return i
		}
	}

	return -1
}
