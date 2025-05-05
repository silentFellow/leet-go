package leetcode

func pushDominoes(dominoes string) string {
	n := len(dominoes)
	curr := []rune(dominoes)

	for {
		changed := false
		next := make([]rune, n)
		copy(next, curr)

		for i := range n {
			if curr[i] == 'R' && i+1 < n && curr[i+1] == '.' {
				if i+2 >= n || curr[i+2] != 'L' {
					next[i+1] = 'R'
					changed = true
				}
			} else if curr[i] == 'L' && i-1 >= 0 && curr[i-1] == '.' {
				if i-2 < 0 || curr[i-2] != 'R' {
					next[i-1] = 'L'
					changed = true
				}
			}
		}

		if !changed {
			break
		}
		curr = next
	}

	return string(curr)
}
