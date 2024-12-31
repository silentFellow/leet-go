package leetcode

func sumOfSquares(n int) int {
	val := 0

	for n != 0 {
		digit := n % 10
		val += digit * digit
		n /= 10
	}

	return val
}

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	slow, fast := n, n

	for {
		slow = sumOfSquares(slow)
		fast = sumOfSquares(sumOfSquares(fast))

		if slow == 1 || fast == 1 {
			return true
		}
		if slow == fast {
			return false
		}
	}
}

// func isHappy(n int) bool {
// 	hmap := make(map[int]bool)
//
// 	for {
// 		if _, exists := hmap[n]; exists {
// 			break
// 		} else {
// 			hmap[n] = true
// 		}
//
// 		n = sumOfSquares(n)
// 		if n == 1 {
// 			return true
// 		}
// 	}
//
// 	return false
// }

// Node:
// happy number always finish with 1
// if not always finish 2ith 4 in the loop
// just a fact, no proof

// func isHappy(n int) bool {
// 	for n != 1 && n != 4 {
// 		n = sumOfSquares(n)
// 	}
//
// 	return n == 1
// }
