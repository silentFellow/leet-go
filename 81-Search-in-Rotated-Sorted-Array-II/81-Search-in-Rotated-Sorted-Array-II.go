package leetcode

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2

		if nums[m] == target {
			return true
		}

		// reduce the search space if cannot find which side is sorted
		if nums[m] == nums[l] && nums[m] == nums[r] {
			l++
			r--
			continue
		}

		// left sorted
		if nums[l] <= nums[m] {
			if nums[l] <= target && target <= nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else { // right sorted
			if nums[m] <= target && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}

	}

	return false
}
