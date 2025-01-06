package leetcode

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	diff := make([]int, n+1)

	for _, shift := range shifts {
		start, end, direction := shift[0], shift[1], shift[2]
		if direction == 1 {
			diff[start]++
			diff[end+1]--
		} else {
			diff[start]--
			diff[end+1]++
		}
	}

	netShift := make([]int, n)
	cumulativeShift := 0
	for i := 0; i < n; i++ {
		cumulativeShift += diff[i]
		netShift[i] = cumulativeShift
	}

	strRune := []rune(s)
	for i := 0; i < n; i++ {
		shift := (netShift[i]%26 + 26) % 26 // Normalize shift to handle negatives
		strRune[i] = 'a' + (rune(s[i]-'a')+rune(shift))%26
	}

	return string(strRune)
}
