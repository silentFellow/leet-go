package leetcode

func wordSubsets(words1 []string, words2 []string) []string {
	count_2 := make(map[rune]int)
	for _, word := range words2 {
		count_w := wordCounter(word)
		for char, count := range count_w {
			count_2[char] = max(count_2[char], count)
		}
	}

	ans := make([]string, 0)
	for _, word := range words1 {
		count_1 := wordCounter(word)
		flag := true
		for char, count := range count_2 {
			if count_1[char] < count {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, word)
		}
	}

	return ans
}

func wordCounter(word string) map[rune]int {
	counter := make(map[rune]int)
	for _, char := range word {
		counter[char]++
	}
	return counter
}
