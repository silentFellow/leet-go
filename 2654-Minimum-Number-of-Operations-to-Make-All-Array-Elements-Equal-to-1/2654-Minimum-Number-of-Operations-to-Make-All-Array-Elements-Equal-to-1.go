package leetcode

func gcd(m, n int) int {
	if n == 0 {
		return m
	}

	return gcd(n, m%n)
}

func minOperations(nums []int) int {
	n := len(nums)

	hmap := make(map[int]int)
	for _, v := range nums {
		hmap[v]++
	}

	if _, ok := hmap[1]; ok {
		count := 0
		for k, v := range hmap {
			if k != 1 {
				count += v
			}
		}
		return count
	}

	minLen := n + 1
	for i, v := range nums {
		gcdVal := v
		for j := i + 1; j < n; j++ {
			gcdVal = gcd(gcdVal, nums[j])

			if gcdVal == 1 {
				minLen = min(minLen, j-i+1)
				break
			}
		}
	}

	if minLen == n+1 {
		return -1
	}

	return (minLen - 1) + (n - 1)
}
