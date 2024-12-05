package leetcode

func canChange(start string, target string) bool {
	l, r := 0, 0

	for l < len(start) && r < len(target) {
		for l < len(start)-1 && start[l] == '_' {
			l++
		}

		for r < len(target)-1 && target[r] == '_' {
			r++
		}

		// start[l] and start[r] can be target[l] && target[r] from 2nd condition
    // since both are anyway wqual by first condition
		if (start[l] != target[r]) || (start[l] == 'L' && l < r) || (start[l] == 'R' && l > r) {
			return false
		}

		l++
		r++
	}

	for l < len(start) && start[l] == '_' {
		l++
	}

	for r < len(target) && target[r] == '_' {
		r++
	}

	return ((l == len(start)) && (r == len(target)))
}
