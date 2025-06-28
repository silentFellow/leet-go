package leetcode

import "sort"

func maxSubsequence(nums []int, k int) []int {
	withIdx := make([][2]int, 0)
	for i, v := range nums {
		withIdx = append(withIdx, [2]int{v, i})
	}

	sort.Slice(withIdx, func(i, j int) bool {
		return withIdx[i][0] > withIdx[j][0]
	})

	req := make(map[int]struct{})
	for i := range k {
		req[withIdx[i][1]] = struct{}{}
	}

	ans := make([]int, 0)
	for i, v := range nums {
		if _, ok := req[i]; ok {
			ans = append(ans, v)
		}
	}

	return ans
}
