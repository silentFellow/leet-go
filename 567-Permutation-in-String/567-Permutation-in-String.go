package leetcode

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

  f, s := counter(s1), counter(s2[:len(s1)])

	if isEqual(f, s) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		start := i - len(s1)
		s[rune(s2[i])]++
		s[rune(s2[start])]--

		if v := s[rune(s2[start])]; v == 0 {
			delete(s, rune(s2[start]))
		}

		if isEqual(f, s) {
			return true
		}
	}

	return false
}

func counter(s string) map[rune]int {
	hmap := make(map[rune]int)

	for _, v := range s {
		hmap[v]++
	}

	return hmap
}

func isEqual(m1, m2 map[rune]int) bool {
	for key, val := range m1 {
		if v2, ok := m2[key]; !ok || val != v2 {
			return false
		}
	}

	return true
}
