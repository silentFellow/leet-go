package leetcode

func rotate(matrix [][]int) {
	// transporse
	for i := range matrix {
		for j := (i + 1); j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// reverse
	for i := range len(matrix) {
		for j, k := 0, len(matrix)-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}
