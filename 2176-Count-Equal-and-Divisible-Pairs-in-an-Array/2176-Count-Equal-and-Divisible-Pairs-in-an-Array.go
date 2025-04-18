package leetcode

func countPairs(nums []int, k int) int {
	ans := 0
	hmap := make(map[int][]int)

	for i, v := range nums {
		if arr, ok := hmap[v]; ok {
			for _, idx := range arr {
				if (i*idx)%k == 0 {
					ans++
				}
			}
		}

		hmap[v] = append(hmap[v], i)
	}

	return ans
}
