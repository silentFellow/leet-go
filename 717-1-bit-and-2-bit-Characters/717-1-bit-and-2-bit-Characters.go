package leetcode

func isOneBitCharacter(bits []int) bool {
	jumpTwice := false

	for i := 0; i < len(bits); i++ {
		v := bits[i]

		if v == 1 {
			i++
			jumpTwice = true
		} else {
			jumpTwice = false
		}
	}

	return !jumpTwice
}
