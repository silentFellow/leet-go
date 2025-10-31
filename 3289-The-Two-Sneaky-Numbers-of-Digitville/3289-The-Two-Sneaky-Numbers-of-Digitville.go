package leetcode

func getSneakyNumbers(nums []int) []int {
	ans := make([]int, 0)

	hmap := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := hmap[num]; ok {
			ans = append(ans, num)
		} else {
			hmap[num] = struct{}{}
		}
	}

	return ans
}
