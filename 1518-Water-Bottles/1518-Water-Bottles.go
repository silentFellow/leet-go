package leetcode

func numWaterBottles(numBottles int, numExchange int) int {
	ans := 0

	emptyBottles := 0
	for numBottles > 0 {
		ans += numBottles

		emptyBottles += (numBottles % numExchange)
		numBottles /= numExchange

		if emptyBottles >= numExchange {
			numBottles += emptyBottles / numExchange
			emptyBottles %= numExchange
		}
	}

	return ans
}
