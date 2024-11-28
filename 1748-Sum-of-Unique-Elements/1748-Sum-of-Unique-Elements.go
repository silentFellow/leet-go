package leetcode

func sumOfUnique(nums []int) int {
	hmap := make(map[int]int)

	for _, val := range nums {
		hmap[val] = hmap[val] + 1
	}

	sum := 0
	for key, val := range hmap {
		if val == 1 {
			sum += key
		}
	}

	return sum
}
