package leetcode

func largestGoodInteger(num string) string {
	ans := ""

	for i := 2; i < len(num); i++ {
		if num[i] == num[i-1] && num[i-1] == num[i-2] {
			if ans == "" {
				ans = string(num[i])
			} else {
				if ans < string(num[i]) {
					ans = string(num[i])
				}
			}
		}
	}

	return ans + ans + ans
}
