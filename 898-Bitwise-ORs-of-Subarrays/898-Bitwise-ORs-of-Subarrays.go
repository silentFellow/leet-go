package leetcode

func subarrayBitwiseORs(arr []int) int {
	prev := make(map[int]struct{})
	res := make(map[int]struct{})

	for _, v := range arr {
		cur := make(map[int]struct{})
		cur[v] = struct{}{}

		for k := range prev {
			cur[v|k] = struct{}{}
		}

		for k := range cur {
			res[v|k] = struct{}{}
		}

		prev = cur
	}

	return len(res)
}
