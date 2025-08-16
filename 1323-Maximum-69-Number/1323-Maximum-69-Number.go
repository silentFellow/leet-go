package leetcode

func maximum69Number(num int) int {
	narr := make([]int, 0)
	for num != 0 {
		cur := num % 10
		narr = append([]int{cur}, narr...)
		num /= 10
	}

	fs := -1
	for i, v := range narr {
		if v == 6 {
			fs = i
			break
		}
	}

	if fs != -1 {
		narr[fs] = 9
	}

	ans := 0
	for _, v := range narr {
		ans = (ans * 10) + v
	}

	return ans
}
