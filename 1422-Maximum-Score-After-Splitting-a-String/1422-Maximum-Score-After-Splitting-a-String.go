package leetcode

func maxScore(s string) int {
	// calculate all one count
	// so that we can know how much
	// one will be there in the right side always
	oneCount := 0
	for _, char := range s {
		if string(char) == "1" {
			oneCount++
		}
	}

	zeroCount := 0
	maxScore := 0
	for i, char := range s {
		// avoid last element, since the substring cannot be empty
		if i == len(s)-1 {
			continue
		}

		str := string(char)
		if str == "0" {
			zeroCount++
		} else {
			oneCount--
		}

		maxScore = max(maxScore, oneCount+zeroCount)
	}

	return maxScore
}
