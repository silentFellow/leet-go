package leetcode

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)

	join := func(arr []int) int {
		joined := 0
		for _, v := range arr {
			joined = (joined * 10) + v
		}
		return joined
	}

	hmap := make(map[int]struct{})
	var backtrack func(arr []int, left, right int)
	backtrack = func(arr []int, left, right int) {
		if left == right {
			perm := make([]int, len(arr))
			copy(perm, arr)

			joined := join(perm)
			if _, ok := hmap[joined]; !ok {
				res = append(res, perm)
				hmap[joined] = struct{}{}
			}
		}

		for i := left; i <= right; i++ {
			arr[left], arr[i] = arr[i], arr[left]
			backtrack(arr, left+1, right)
			arr[left], arr[i] = arr[i], arr[left]
		}
	}

	backtrack(nums, 0, len(nums)-1)
	return res
}
