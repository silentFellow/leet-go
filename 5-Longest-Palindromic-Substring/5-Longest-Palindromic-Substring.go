package leetcode

// TODO: O(n^2) time complexity try to optimize the solution

func expandOutwards(s string, l, r int, maxLen *int, res *string) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		if (r - l + 1) > *maxLen {
			*maxLen = r - l + 1
			*res = s[l : r+1]
		}
		l--
		r++
	}
}

func longestPalindrome(s string) string {
	res := ""
	maxLen := 0
	var l, r int

  // for odd length
	for i := range s {
		l, r = i, i
		expandOutwards(s, l, r, &maxLen, &res)
	}

  // for even length
	for i := range s {
		l, r = i, i+1
		expandOutwards(s, l, r, &maxLen, &res)
	}

	return res
}
