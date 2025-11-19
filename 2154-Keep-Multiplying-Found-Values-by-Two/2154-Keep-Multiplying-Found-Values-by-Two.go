package leetcode

func findFinalValue(nums []int, original int) int {
	hmap := make(map[int]struct{})
	for _, v := range nums {
		hmap[v] = struct{}{}
	}

	for {
		if _, ok := hmap[original]; !ok {
			return original
		}
		original *= 2
	}
}
