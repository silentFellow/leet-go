package leetcode

func numRabbits(answers []int) int {
	hmap := make(map[int]int)
	for _, v := range answers {
		hmap[v]++
	}

	missing := 0
	for k, v := range hmap {
		groupSize := k + 1
		groups := (v + groupSize - 1) / groupSize
		missing += groups * groupSize
	}

	return missing
}
