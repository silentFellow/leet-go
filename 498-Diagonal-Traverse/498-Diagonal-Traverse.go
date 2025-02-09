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

// hashmap approach
// func findDiagonalOrder(mat [][]int) []int {
// 	hmap := make(map[int][]int)
//
// 	for i, row := range mat {
// 		for j, col := range row {
// 			hmap[i+j] = append(hmap[i+j], col)
// 		}
// 	}
//
// 	m, n := len(mat), len(mat[0])
// 	ans := []int{}
// 	for i := range m + n - 1 {
// 		val := hmap[i]
// 		if i%2 != 0 {
// 			for _, v := range val {
// 				ans = append(ans, v)
// 			}
// 		} else {
// 			for j := len(val) - 1; j >= 0; j-- {
// 				ans = append(ans, val[j])
// 			}
// 		}
// 	}
//
// 	return ans
// }
