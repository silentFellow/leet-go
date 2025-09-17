package leetcode

func replaceNonCoprimes(nums []int) []int {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	lcm := func(a, b int) int {
		return a / gcd(a, b) * b
	}

	stack := []int{}
	for _, num := range nums {
		for len(stack) > 0 && gcd(stack[len(stack)-1], num) > 1 {
			num = lcm(stack[len(stack)-1], num)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	return stack
}
