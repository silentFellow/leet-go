package leetcode

func countCollisions(directions string) int {
	nonS := 0
	for _, dir := range directions {
		if dir != 'S' {
			nonS++
		}
	}

	leadingLeft := 0
	for _, dir := range directions {
		if dir == 'L' {
			leadingLeft++
		} else {
			break
		}
	}

	trailingRight := 0
	for i := len(directions) - 1; i >= 0; i-- {
		if directions[i] == 'R' {
			trailingRight++
		} else {
			break
		}
	}

	return nonS - leadingLeft - trailingRight
}
