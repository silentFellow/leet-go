package leetcode

func searchRange(nums []int, target int) []int {
  // edge case
  if len(nums) == 0 {
    return []int{-1,-1}
  }

	lb, ub, m := -1, -1, 0

  // find lowerbound
	l, r := 0, len(nums)-1
	for l <= r {
		m = (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

  // if lowerbound not wxists upperbound won't exists
  if l < 0 || l > len(nums)-1 || nums[l] != target {
    return []int{-1,-1}
  }

  lb = l

  // find upperbound, just start with lowerbound, no need to be from first
	l, r = lb, len(nums)-1
	for l <= r {
		m = (l + r) / 2
		if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
  ub = r // since lowerbound exists no check for valid range

	return []int{lb, ub}
}
