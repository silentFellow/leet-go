package leetcode

func countBadPairs(nums []int) int64 {
	hmap := make(map[int]int)

	ans := 0

	for i, num := range nums {
		diff := num - i
		ans += i - hmap[diff]
		hmap[diff]++
	}

	return int64(ans)
}
