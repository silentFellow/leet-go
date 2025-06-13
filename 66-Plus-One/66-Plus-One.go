package leetcode

func plusOne(digits []int) []int {
	carry := 1

	for i := len(digits) - 1; i >= 0; i-- {
		neoCarry := (digits[i] + carry) / 10
		digits[i] = (digits[i] + carry) % 10
		carry = neoCarry

		if carry == 0 {
			break
		}
	}

	if carry != 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}
