package leetcode

func validPalindrome(s string) bool {
	isPalindrome := func(s string) bool {
		l, r := 0, len(s)-1

		for l < r {
			if s[l] != s[r] {
				return false
			}

			l++
			r--
		}

		return true
	}

	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return isPalindrome(s[l+1:r+1]) || isPalindrome(s[l:r])
		}

		l++
		r--
	}

	return true
}
