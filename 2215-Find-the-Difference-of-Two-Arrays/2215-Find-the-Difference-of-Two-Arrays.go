package leetcode

func findDifference(nums1 []int, nums2 []int) [][]int {
	hmap := make(map[int]bool)
	for _, val := range nums1 {
		hmap[val] = true
	}

	ans := make([][]int, 2)
	appended := make(map[int]struct{})
	for _, val := range nums2 {
		if _, ok := hmap[val]; ok {
			hmap[val] = false
			continue
		}

		if _, ok := appended[val]; !ok {
			ans[1] = append(ans[1], val)
			appended[val] = struct{}{}
		}
	}

	appended = make(map[int]struct{})
	for key, val := range hmap {
		if val == true {
			if _, ok := appended[key]; !ok {
				ans[0] = append(ans[0], key)
				appended[key] = struct{}{}
			}
		}
	}

	return ans
}
