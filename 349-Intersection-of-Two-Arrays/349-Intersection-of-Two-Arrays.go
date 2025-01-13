package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	hmap := make(map[int]int)

	for _, num := range nums1 {
		hmap[num] = 1
	}

	for _, num := range nums2 {
		if _, ok := hmap[num]; ok {
			hmap[num] = 0
		}
	}

	ans := []int{}
	for key, val := range hmap {
		if val == 0 {
			ans = append(ans, key)
		}
	}

	return ans
}
