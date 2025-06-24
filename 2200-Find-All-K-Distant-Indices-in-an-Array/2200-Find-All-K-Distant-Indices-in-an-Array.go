package leetcode

func findKDistantIndices(nums []int, key int, k int) []int {
	loc := make([]int, 0)
	for i, v := range nums {
		if v == key {
			loc = append(loc, i)
		}
	}

	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

	ans := make([]int, 0)
	for i := range len(nums) {
		for _, v := range loc {
			if abs(i-v) <= k {
				ans = append(ans, i)
				break
			}
		}
	}

	return ans
}
