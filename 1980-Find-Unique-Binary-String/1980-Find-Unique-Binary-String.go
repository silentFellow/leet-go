package leetcode

// TODO: do it with backtracking

func findDifferentBinaryString(nums []string) string {
	queue, n := []string{"0", "1"}, len(nums)

	for range n - 1 {
		cur := queue
		for _, v := range cur {
			queue = append(queue, v+"0")
			queue = append(queue, v+"1")
		}

		queue = queue[len(cur):]
	}

	set := make(map[string]struct{})
	for _, v := range nums {
		set[v] = struct{}{}
	}

	for _, v := range queue {
		if _, ok := set[v]; !ok {
			return v
		}
	}

	return ""
}
