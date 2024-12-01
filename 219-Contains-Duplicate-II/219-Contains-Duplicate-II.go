package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	hmap := make(map[int]int)

	for i, val := range nums {
		if existIndex, exists := hmap[val]; exists {
			if i-existIndex == k {
				return true
			}
		}

		hmap[val] = i
	}

	return false
}
