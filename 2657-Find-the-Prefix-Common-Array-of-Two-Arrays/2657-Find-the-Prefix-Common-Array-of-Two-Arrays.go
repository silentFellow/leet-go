package leetcode

func findThePrefixCommonArray(A []int, B []int) []int {
	ans := make([]int, len(A))
	hmap := make(map[int]int)

	twoCount := 0
	for i := 0; i < len(A); i++ {
		hmap[A[i]]++
		hmap[B[i]]++

		if hmap[A[i]] == 2 {
			twoCount++
		}

		if A[i] != B[i] && hmap[B[i]] == 2 {
			twoCount++
		}

		ans[i] = twoCount
	}

	return ans
}
