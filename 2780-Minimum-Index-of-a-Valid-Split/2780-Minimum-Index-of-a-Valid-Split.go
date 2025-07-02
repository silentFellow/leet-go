package leetcode

func minimumIndex(nums []int) int {
	n := len(nums)
	isMajority := func(n int, count int) bool {
		return count*2 > n
	}

	initialMajority := nums[0]
	hmap := make(map[int]int)
	for _, v := range nums {
		hmap[v]++

		if isMajority(n, hmap[v]) {
			initialMajority = v
		}
	}
	majorityCount := hmap[initialMajority]

	count := 0
	for i, v := range nums {
		if v == initialMajority {
			count++
		}

		if isMajority(i+1, count) && isMajority(n-i-1, majorityCount-count) {
			return i
		}
	}

	return -1
}
