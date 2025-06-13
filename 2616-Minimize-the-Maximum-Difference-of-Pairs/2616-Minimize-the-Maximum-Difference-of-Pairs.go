package leetcode

import "slices"

func minimizeMax(nums []int, p int) int {
	slices.Sort(nums)

	n := len(nums)

	isValid := func(diff int) bool {
		cnt, idx := 0, 1
		for idx < n {
			if nums[idx]-nums[idx-1] <= diff {
				cnt++
				idx++
			}
			idx++

			if cnt >= p {
				break
			}
		}

		return cnt >= p
	}

	ans := -1
	l, r := 0, nums[n-1]-nums[0]
	for l <= r {
		m := (l + r) / 2
		if isValid(m) {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return ans
}
