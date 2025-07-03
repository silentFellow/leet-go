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

	nonDistinct, count, idx := len(hmap), 0, 0
	for nonDistinct != 0 {
		for i := idx; i < idx+3 && i < len(nums); i++ {
			hmap[nums[i]]--
			if hmap[nums[i]] == 1 {
				delete(hmap, nums[i])
				nonDistinct--
			}
		}

		idx += 3
		count++
	}

	return count
}
