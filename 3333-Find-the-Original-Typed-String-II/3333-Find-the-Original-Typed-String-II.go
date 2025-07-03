package leetcode

var MOD int = 1_000_000_007

func possibleStringCount(word string, k int) int {
	if len(word) == 0 {
		return 0
	}

	groups, cnt := make([]int, 0), 1
	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			cnt++
		} else {
			groups = append(groups, cnt)
			cnt = 1
		}
	}
	groups = append(groups, cnt)

	total := 1
	for _, v := range groups {
		total = (total * v) % MOD
	}

	minLen := len(groups)
	if minLen >= k {
		return total
	}

	// the minlen of the string will be of minLen, coz each letter in the group should be present atleast once
	// the maxlen is the len of the string
	// to find only possible strings, remove all unwanted things
	// if k = 10, minLen = 5, maxLen = 12
	// req = 10 + 11 + 12
	// this can be writted as (total possible combination) - (5 + 6 + 7 + 8 + 9)

	dp := make([]int, k)
	dp[0] = 1 // only one possible way to build an empty string

	for _, group := range groups {
		newDp := make([]int, k)
		possible := 0

		for i := range k {
			if i > 0 {
				possible = (possible + dp[i-1]) % MOD
			}
			if i > group {
				possible = (possible - dp[i-group-1]) % MOD
			}

			newDp[i] = possible
		}

		dp = newDp
	}

	invalid := 0
	for i := minLen; i < k; i++ {
		invalid = (invalid + dp[i]) % MOD
	}

	return (total - invalid + MOD) % MOD
}
