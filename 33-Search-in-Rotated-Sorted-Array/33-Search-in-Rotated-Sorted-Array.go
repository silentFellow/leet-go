package leetcode

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l+r)/2

		if nums[m] == target { return m }

		// left sorted
		if nums[l] <= nums[m] {
			if nums[l] <= target && target <= nums[m] {
				r = m-1
			} else {
				l = m+1
			}
		} else {
			if nums[m] <= target && target <= nums[r] {
				l = m+1
			} else {
				r = m-1
			}
		}
	}

	return -1
}

// O(2 * log n) => O(log n)
// func search(nums []int, target int) int {
// 	l, r := 0, len(nums)-1
//
// 	// loop to find the minimum element
// 	for l < r {
// 		m := (l + r) / 2
//
// 		if nums[m] > nums[r] {
// 			l = m + 1
// 		} else {
//       r = m
//     }
// 	}
//
//   pivot := l
// 	l, r = 0, len(nums)-1
// 	for l <= r {
// 		m := (l + r) / 2
//     readMid := (m+pivot) % len(nums)
//
// 		if nums[readMid] == target {
// 			return readMid
// 		} else if nums[readMid] < target {
//       l = m+1
// 		} else {
//       r = m-1
// 		}
// 	}
//
// 	return -1
// }
