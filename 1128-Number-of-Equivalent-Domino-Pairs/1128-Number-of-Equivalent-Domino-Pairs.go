package leetcode

func numEquivDominoPairs(dominoes [][]int) int {
	hmap := make(map[[2]int]int)

	for _, v := range dominoes {
		a, b := v[0], v[1]
		if a > b {
			a, b = b, a
		}
		hmap[[2]int{a, b}]++
	}

	ans := 0
	for _, v := range hmap {
		ans += v * (v - 1) / 2
	}

	return ans
}
