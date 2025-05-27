package leetcode

func longestPalindrome(s string) int {
	ans := 0

	hmap := make(map[rune]int)
	for _, char := range s {
		hmap[char]++
	}

	for k, v := range hmap {
		red := v / 2
		ans += red * 2
		hmap[k] -= red * 2
		if hmap[k] == 0 {
			delete(hmap, k)
		}
	}

	if ans%2 == 0 && len(hmap) > 0 {
		ans++
	}

	return ans
}
