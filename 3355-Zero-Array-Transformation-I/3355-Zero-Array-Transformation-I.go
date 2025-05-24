package leetcode

func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)

	diff := make([]int, n+1)

	for _, query := range queries {
		start, end := query[0], query[1]

		diff[start]--
		if end + 1 < n {
			diff[end]++
		}
	}

	sum_val := 0
	for i := range n {
		sum_val += diff[i]
		if nums[i] > -sum_val {
			return false
		}
	}

	return true
}
