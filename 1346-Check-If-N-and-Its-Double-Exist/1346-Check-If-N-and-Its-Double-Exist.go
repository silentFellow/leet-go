package leetcode

func checkIfExist(arr []int) bool {
	hmap := make(map[int]bool)

	for _, val := range arr {
		mulTarget := 2 * val
		divTarget := val / 2

		if _, exists := hmap[mulTarget]; exists {
			return true
		} else if _, exists := hmap[divTarget]; exists && val%2 == 0 {
			return true
		}
		hmap[val] = true
	}

	return false
}
