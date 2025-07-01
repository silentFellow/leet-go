package leetcode

func findLHS(nums []int) int {
	hmap := make(map[int]int)
	for _, v := range nums {
		hmap[v]++
	}

	ans := 0
	for key, val := range hmap {
		if less, ok := hmap[key-1]; ok {
			ans = max(ans, val+less)
		}

		if greater, ok := hmap[key+1]; ok {
			ans = max(ans, val+greater)
		}
	}

	return ans
}
