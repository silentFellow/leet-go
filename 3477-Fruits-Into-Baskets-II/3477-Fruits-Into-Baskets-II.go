package leetcode

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	consumed := 0

	for l := range len(fruits) {
		for r := range len(baskets) {
			if fruits[l] <= baskets[r] {
				consumed++
				baskets[r] *= -1
				break
			}
		}
	}

	return len(fruits) - consumed
}
