package leetcode

func repeatedNTimes(nums []int) int {
	hmap := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := hmap[num]; ok {
			return num
		}
		hmap[num] = struct{}{}
	}

	return -1
}
