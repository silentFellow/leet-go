package leetcode

func findLexSmallestString(s string, a int, b int) string {
	addVal := func(s string, v int) string {
		conv := []rune(s)
		for i := 1; i < len(conv); i += 2 {
			conv[i] = rune((int(conv[i]-'0')+v)%10 + '0')
		}
		return string(conv)
	}

	rotate := func(s string, times int) string {
		mod := times % len(s)
		return s[len(s)-mod:] + s[:len(s)-mod]
	}

	visited := make(map[string]struct{})
	queue := []string{s}
	ans := s

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if _, ok := visited[cur]; ok {
			continue
		}
		visited[cur] = struct{}{}

		if cur < ans {
			ans = cur
		}

		added := addVal(cur, a)
		rotated := rotate(cur, b)
		if _, ok := visited[added]; !ok {
			queue = append(queue, added)
		}
		if _, ok := visited[rotated]; !ok {
			queue = append(queue, rotated)
		}
	}

	return ans
}
