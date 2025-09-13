package leetcode

func doesAliceWin(s string) bool {
	vowels := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}

	cnt := 0
	for _, char := range s {
		if _, ok := vowels[char]; ok {
			cnt++
		}
	}

	return cnt > 0
}
