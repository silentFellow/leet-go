package leetcode

func minimumLength(s string) int {
	hmap := make(map[rune]int)

	for _, char := range s {
		if val, ok := hmap[char]; ok {
			if val == 2 {
				hmap[char] = 0
			}
		}

		hmap[char]++
	}

	ans := 0
	for _, val := range hmap {
		ans += val
	}

	return ans
}
