package leetcode

// TODO: use a one-pass algorithm using only constant extra space

func sortColors(nums []int) {
	zc, oc, tc := 0, 0, 0
	for _, v := range nums {
		if v == 0 {
			zc++
		} else if v == 1 {
			oc++
		} else {
			tc++
		}
	}

	counts := [3]int{zc, oc, tc}
	idx := 0

	for v, count := range counts {
		for range count {
			nums[idx] = v
			idx++
		}
	}
}
