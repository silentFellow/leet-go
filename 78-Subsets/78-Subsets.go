package leetcode

import (
	"strconv"
	"strings"
)

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	exists := make(map[string]struct{})

	var dfs func(cur []int, rem []int, i int)
	dfs = func(cur []int, rem []int, i int) {
		if i == len(rem) {
			var formatted strings.Builder
			for _, v := range cur {
				formatted.WriteString(strconv.Itoa(v))
			}

			if _, ok := exists[formatted.String()]; !ok {
				ans = append(ans, cur)
				exists[formatted.String()] = struct{}{}
			}
			return
		}

		dfs(cur, rem, i+1)
		dfs(append(cur, rem[i]), rem, i+1)
	}

	for i := range len(nums) {
		rem := []int{}
		rem = append(rem, nums[:i]...)
		rem = append(rem, nums[i+1:]...)
		dfs([]int{}, rem, 0)
	}
	return ans
}
