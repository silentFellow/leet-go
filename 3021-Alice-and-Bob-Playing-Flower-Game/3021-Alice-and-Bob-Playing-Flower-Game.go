package leetcode

func flowerGame(n int, m int) int64 {
	odd_n, even_n := 0, 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			odd_n++
		} else {
			even_n++
		}
	}

	odd_m, even_m := 0, 0
	for i := 1; i <= m; i++ {
		if i%2 == 1 {
			odd_m++
		} else {
			even_m++
		}
	}

	return int64((odd_n * even_m) + (even_n * odd_m))
}
