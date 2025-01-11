package leetcode

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}

	hmap := make(map[rune]int)

	for _, val := range s {
		hmap[val]++
	}

	oddCount := 0
	for _, val := range hmap {
		if val%2 == 1 {
			oddCount++
		}
	}

	return oddCount <= k
}
