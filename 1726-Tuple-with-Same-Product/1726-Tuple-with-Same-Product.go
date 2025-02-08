package leetcode

func tupleSameProduct(nums []int) int {
	hmap := make(map[int]int)

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			hmap[nums[i]*nums[j]]++
		}
	}

	ans := 0
	for _, val := range hmap {
		ans += (val - 1) * val * 4
	}

	return ans
}
