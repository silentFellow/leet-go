package leetcode

func replaceElements(arr []int) []int {
	maxv := -1

	for r := len(arr) - 1; r >= 0; r-- {
		v := arr[r]
		arr[r] = maxv
		maxv = max(maxv, v)
	}

	return arr
}
