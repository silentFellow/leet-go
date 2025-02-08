package leetcode

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	i, j := 0, 0

	up := true
	ans := make([]int, 0)
	for i != m-1 || j != n-1 {
		ans = append(ans, mat[i][j])

		if up {
			if i-1 < 0 && j+1 >= n {
				i++
				up = false
			} else if i-1 < 0 {
				j++
				up = false
			} else if j+1 >= n {
				i++
				up = false
			} else {
				i--
				j++
			}
		} else {
			if i+1 >= m && j-1 < 0 {
				j++
				up = true
			} else if j-1 < 0 {
				i++
				up = true
			} else if i+1 >= m {
				j++
				up = true
			} else {
				i++
				j--
			}
		}
	}

	ans = append(ans, mat[m-1][n-1])

	return ans
}
