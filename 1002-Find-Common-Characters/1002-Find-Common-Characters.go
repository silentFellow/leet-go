package leetcode

func commonChars(words []string) []string {
	hmap := make(map[rune][]int)

	for i, word := range words {
		for _, char := range word {
			if _, ok := hmap[char]; !ok {
				hmap[char] = make([]int, len(words))
			}
			hmap[char][i]++
		}
	}

	ans := []string{}
	for key, val := range hmap {
		time, exists := findMin(val)
		if !exists {
			continue
		}

		for range time {
			ans = append(ans, string(key))
		}
	}

	return ans
}

func findMin(arr []int) (int, bool) {
	minv := arr[0]
	for _, val := range arr {
		if val == 0 {
			return minv, false
		}
		minv = min(minv, val)
	}

	return minv, true
}
