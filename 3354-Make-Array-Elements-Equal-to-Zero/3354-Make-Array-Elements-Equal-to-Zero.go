package leetcode

func countValidSelections(nums []int) int {
	n := len(nums)
	ans := 0

	isEmpty := func(a []int) bool {
		for _, v := range a {
			if v != 0 {
				return false
			}
		}
		return true
	}

	for i, v := range nums {
		if v != 0 {
			continue
		}

		arr := append([]int(nil), nums...)
		curr := i
		dir := 1
		for curr >= 0 && curr < n {
			if arr[curr] == 0 {
				curr += dir
			} else {
				arr[curr]--
				dir *= -1
				curr += dir
			}
		}
		if isEmpty(arr) {
			ans++
		}

		arr = append([]int(nil), nums...)
		curr = i
		dir = -1

		for curr >= 0 && curr < n {
			if arr[curr] == 0 {
				curr += dir
			} else {
				arr[curr]--
				dir *= -1
				curr += dir
			}
		}
		if isEmpty(arr) {
			ans++
		}
	}

	return ans
}
