package leetcode

func countLargestGroup(n int) int {
	hmap := make(map[int][]int)

	for i := 1; i <= n; i++ {
		sum := sumOfNo(i)
		hmap[sum] = append(hmap[sum], i)
	}

	largestSize := 0
	for _, val := range hmap {
		n := len(val)
		largestSize = max(largestSize, n)
	}

	ans := 0
	for _, val := range hmap {
		if len(val) == largestSize {
			ans++
		}
	}

	return ans
}

func sumOfNo(n int) int {
	v := 0

	for n != 0 {
		v += n % 10
		n /= 10
	}

	return v
}
