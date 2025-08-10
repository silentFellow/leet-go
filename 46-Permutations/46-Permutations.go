package leetcode

func permute(nums []int) [][]int {
	res := make([][]int, 0)

	var backtrack func(arr []int, left, right int)
	backtrack = func(arr []int, left, right int) {
		if left == right {
			perm := make([]int, len(arr))
			copy(perm, arr)
			res = append(res, perm)
			return
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
