package leetcode

var conv map[byte]int = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	ans := 0

  for i:=0; i<len(s); i++ {
		cur := conv[s[i]]
		if i < len(s)-1 {
			next := conv[s[i+1]]

			if next > cur {
				ans += (next - cur)
				i++
			} else {
				ans += cur
			}
		} else {
			ans += cur
		}
	}

	return ans
}
