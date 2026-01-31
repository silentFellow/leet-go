package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {
	isFound := false
	var ans byte = 'z'

	for _, letter := range letters {
		if letter > target {
			ans = min(ans, letter)
			isFound = true
		}
	}

	if !isFound {
		return letters[0]
	}
	return ans
}
