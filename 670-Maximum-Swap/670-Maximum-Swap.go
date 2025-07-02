package leetcode

// you can actually pre-compute the last occurrence of each digit
// last := make([]int, 10)
// for i, v := range split {
// 	last[v] = i
// }

// while finding the max, wee can actually use
// for i, v := range split {
// 	for d:=9; d>v; d-- {
// 		if last[d] > i {
// 			split[i], split[last[d]] = split[last[d]], split[i]
// 			break
// 		}
// 	}
// }

func maximumSwap(num int) int {
	split := make([]int, 0)
	for num != 0 {
		split = append([]int{num % 10}, split...)
		num /= 10
	}

	for i, v := range split {
		maxDigit := v
		maxIdx := i

		for j := len(split) - 1; j > i; j-- {
			if split[j] > maxDigit {
				maxDigit = split[j]
				maxIdx = j
			}
		}

		if maxIdx != i {
			split[maxIdx], split[i] = split[i], split[maxIdx]
			break
		}
	}

	ans := 0
	for _, v := range split {
		ans = (ans * 10) + v
	}

	return ans
}
