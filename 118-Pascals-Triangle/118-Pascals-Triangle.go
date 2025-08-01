package leetcode

func generate(numRows int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, []int{1})

	for i := 1; i < numRows; i++ {
		cur := make([]int, 0)
		for j := range i + 1 {
			if j == 0 || j == i {
				cur = append(cur, 1)
			} else {
				cur = append(cur, ans[i-1][j-1]+ans[i-1][j])
			}
		}

		ans = append(ans, cur)
	}

	return ans
}
