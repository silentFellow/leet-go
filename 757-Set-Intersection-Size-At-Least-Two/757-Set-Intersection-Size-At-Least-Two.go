package leetcode

import "sort"

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] > intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	n := len(intervals)
	ans := 0
	picked := make([]int, 0)

	for i := range n {
		cnt := 0
		for _, v := range picked {
			if v >= intervals[i][0] && v <= intervals[i][1] {
				cnt++
			}
		}
		for j := intervals[i][1]; cnt < 2; j-- {
			picked = append(picked, j)
			ans++
			cnt++
		}
	}

	return ans
}
