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

// one pass algorithm
// func sortColors(nums []int) {
// 	l, m, r := 0, 0, len(nums)-1
//
// 	for m <= r {
// 		if nums[m] == 0 {
// 			nums[l], nums[m] = nums[m], nums[l]
// 			l++
// 			m++
// 		} else if nums[m] == 1 {
// 			m++
// 		} else {
// 			nums[r], nums[m] = nums[m], nums[r]
// 			r--
// 			m++
// 		}
// 	}
// }
