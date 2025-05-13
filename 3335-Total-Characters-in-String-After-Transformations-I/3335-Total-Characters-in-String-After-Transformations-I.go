package leetcode

func lengthAfterTransformations(s string, t int) int {
	const MOD = 1_000_000_007

	// Frequency count of characters
	counts := [26]int{}
	for _, ch := range s {
		counts[ch-'a']++
	}

	for range t {
		next := [26]int{}

		// Apply transformation
		for ch := range 26 {
			c := counts[ch]
			if ch == 25 { // 'z'
				next[0] = (next[0] + c) % MOD // 'a'
				next[1] = (next[1] + c) % MOD // 'b'
			} else {
				next[ch+1] = (next[ch+1] + c) % MOD
			}
		}

		counts = next
	}

	// Sum total length
	total := 0
	for _, c := range counts {
		total = (total + c) % MOD
	}

	return total
}
