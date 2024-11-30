package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)

	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if i != 0 {
			for i < len(nums)-2 && nums[i] == nums[i-1] {
				i++
			}
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] && nums[r] == nums[r-1] {
					l++
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return ans
}
