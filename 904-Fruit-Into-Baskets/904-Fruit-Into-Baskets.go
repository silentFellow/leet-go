package leetcode

func totalFruit(fruits []int) int {
	if len(fruits) == 1 {
		return 1
	}

	hmap := make(map[int]int)

	ans := 0
	left := 0
	hmap[fruits[left]]++
	for right := 1; right < len(fruits); right++ {
		hmap[fruits[right]]++

		for len(hmap) == 3 {
			hmap[fruits[left]]--
			if hmap[fruits[left]] == 0 {
				delete(hmap, fruits[left])
			}
			left++
		}
		ans = max(ans, right-left+1)
	}

	return ans
}
