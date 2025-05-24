package leetcode

import "slices"

func setZeroes(matrix [][]int) {
	// Get the dimensions of the matrix
	m, n := len(matrix), len(matrix[0])

	// Flags to track if the first row or first column should be set to zero
	isFirstRowZero, isFirstColZero := false, false

	// Check if the first row contains any zero
	if slices.Contains(matrix[0], 0) {
		isFirstRowZero = true
	}

	// Check if the first column contains any zero
	for i := range m {
		if matrix[i][0] == 0 {
			isFirstColZero = true
			break
		}
	}

	// Mark corresponding first row and first column cells as zero
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// Set cells to zero based on markers in the first row and column
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// Zero out the first row if needed
	if isFirstRowZero {
		for i := range n {
			matrix[0][i] = 0
		}
	}

	// Zero out the first column if needed
	if isFirstColZero {
		for i := range m {
			matrix[i][0] = 0
		}
	}
}
