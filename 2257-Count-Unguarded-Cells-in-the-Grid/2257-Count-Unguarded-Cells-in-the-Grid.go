package leetcode

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	guarded := make([][]bool, m)
	for i := range guarded {
		guarded[i] = make([]bool, n)
	}

	guardMap := make(map[[2]int]bool)
	for _, g := range guards {
		guardMap[[2]int{g[0], g[1]}] = true
	}

	wallMap := make(map[[2]int]bool)
	for _, w := range walls {
		wallMap[[2]int{w[0], w[1]}] = true
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for _, g := range guards {
		r, c := g[0], g[1]
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			for nr >= 0 && nr < m && nc >= 0 && nc < n {
				if guardMap[[2]int{nr, nc}] || wallMap[[2]int{nr, nc}] {
					break
				}
				guarded[nr][nc] = true
				nr += d[0]
				nc += d[1]
			}
		}
	}

	ans := 0
	for i := range m {
		for j := range n {
			if !guardMap[[2]int{i, j}] && !wallMap[[2]int{i, j}] && !guarded[i][j] {
				ans++
			}
		}
	}

	return ans
}
