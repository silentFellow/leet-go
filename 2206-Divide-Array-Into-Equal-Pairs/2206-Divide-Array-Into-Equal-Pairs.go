package leetcode

func divideArray(nums []int) bool {
	counter := make(map[int]int)
	for _, v := range nums {
		counter[v]++
	}

	possible := 0
	for _, v := range counter {
		possible += v / 2
	}

	return possible == len(nums)/2
}
