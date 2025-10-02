package leetcode

func maxBottlesDrunk(numBottles int, numExchange int) int {
	ans := numBottles
	empty := numBottles

	for empty >= numExchange {
		empty -= numExchange
		numExchange++
		ans++
		empty++
	}

	return ans
}
