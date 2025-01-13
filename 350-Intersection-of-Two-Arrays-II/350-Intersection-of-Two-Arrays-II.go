package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	hmap := make(map[int][]int)

	for _, num := range nums1 {
		if _, ok := hmap[num]; !ok {
			hmap[num] = make([]int, 2)
		}
		hmap[num][0]++
	}

	for _, num := range nums2 {
		if _, ok := hmap[num]; ok {
			hmap[num][1]++
			continue
		}
	}

	ans := []int{}
	for key, val := range hmap {
		if val[0] != 0 && val[1] != 0 {
			time := min(val[0], val[1])
			for range time {
				ans = append(ans, key)
			}
		}
	}

	return ans
}
