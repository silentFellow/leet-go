package leetcode

func minimizeXor(num1 int, num2 int) int {
	countOnes := func(n int) int {
		count := 0
		for n > 0 {
			count += n & 1
			n >>= 1
		}
		return count
	}

	num1Ones := countOnes(num1)
	num2Ones := countOnes(num2)

	if num1Ones == num2Ones {
		return num1
	}

	result := 0
	if num1Ones > num2Ones {
		for i := 31; i >= 0 && num2Ones > 0; i-- {
			if num1&(1<<i) != 0 {
				result |= 1 << i
				num2Ones--
			}
		}
	} else {
		result = num1
		for i := 0; i < 32 && num2Ones > num1Ones; i++ {
			if result&(1<<i) == 0 {
				result |= 1 << i
				num1Ones++
			}
		}
	}

	return result
}
