package leetcode

import "slices"

func triangleNumber(nums []int) int {
	slices.Sort(nums)

	ans := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			l, r := j+1, len(nums)

			for l < r {
				mid := (l + r) / 2

				if nums[i]+nums[j] > nums[mid] {
					l = mid + 1
				} else {
					r = mid
				}
			}

			ans += l - (j + 1)
		}
	}

	return ans
}
