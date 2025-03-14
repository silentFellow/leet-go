package leetcode

func maximumCount(nums []int) int {
	pos, neg := 0, 0

	for _, v := range nums {
		if v < 0 {
			neg++
		} else if v > 0 {
			pos++
		}
	}

	return max(pos, neg)
}
