package leetcode

func minimumOperations(nums []int) int {
	hmap := make(map[int]int)
	for _, v := range nums {
		hmap[v]++
	}

	for key, val := range hmap {
		if val == 1 {
			delete(hmap, key)
		}
	}

	count, start := 0, 0
	for len(hmap) != 0 {
		count++
		for i := start; i < min(start+3, len(nums)); i++ {
			if v, ok := hmap[nums[i]]; ok {
				hmap[nums[i]]--
				if v-1 == 1 {
					delete(hmap, nums[i])
				}
			}
		}
		start += 3
	}

	return count
}
