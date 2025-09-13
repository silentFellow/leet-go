package leetcode

func maxFreqSum(s string) int {
	vowels := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}

	cntMap := make(map[rune]int)
	for _, v := range s {
		cntMap[v]++
	}

	vc, cc := 0, 0
	for k, v := range cntMap {
		if _, ok := vowels[k]; ok {
			vc = max(vc, v)
		} else {
			cc = max(cc, v)
		}
	}

	return vc + cc
}
