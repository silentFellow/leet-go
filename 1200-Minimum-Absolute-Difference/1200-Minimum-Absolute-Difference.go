package leetcode

import "slices"

func minimumAbsDifference(arr []int) [][]int {
	slices.Sort(arr)

	ans := make([][]int, 0)
	ans = append(ans, []int{arr[0], arr[1]})

	mindiff := abs(arr[0] - arr[1])
	for i := 1; i < len(arr); i++ {
		diff := abs(arr[i] - arr[i+1])
		if diff < mindiff {
			ans = [][]int{{arr[i], arr[i+1]}}
			mindiff = diff
		} else if diff == mindiff {
			ans = append(ans, []int{arr[i], arr[i+1]})
		}
	}

	return ans
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}
