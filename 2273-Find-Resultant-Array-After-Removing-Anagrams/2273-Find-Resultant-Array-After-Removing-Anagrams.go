package leetcode

func removeAnagrams(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	isAnagram := func(a, b string) bool {
		first := make([]int, 26)
		for _, char := range a {
			first[char-'a']++
		}

		for _, char := range b {
			first[char-'a']--
		}

		for _, v := range first {
			if v != 0 {
				return false
			}
		}

		return true
	}

	ans := make([]string, 0)
	ans = append(ans, words[0])

	for i := 1; i < len(words); i++ {
		if !isAnagram(words[i], ans[len(ans)-1]) {
			ans = append(ans, words[i])
		}
	}

	return ans
}
