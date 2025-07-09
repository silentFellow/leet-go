package leetcode

func findLucky(arr []int) int {
	hmap := make(map[int]int)
	for _, v := range arr {
		hmap[v]++
	}

	ans := -1
	for key, val := range hmap {
		if key == val {
			ans = max(ans, key)
		}
	}

	return ans
}
