package leetcode

func canMakeSubsequence(str1 string, str2 string) bool {
	l, r := 0, 0

	for r < len(str2) && l < len(str1) {
		var inc byte
		if str1[l] == 122 { // if 'z'
			inc = 97 // make 'a'
		} else { // just increate byte value
			inc = str1[l] + 1
		}

		if str1[l] == str2[r] || inc == str2[r] {
			l++
			r++
			continue
		}

		l++
	}

	if r == len(str2) {
		return true
	}
	return false
}
