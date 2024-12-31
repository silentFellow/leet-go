package leetcode

// NODE: brute forse loop from both array
// with space o(N) => use set to store the address (or) use stack to store the address
// set prefered
// with space o(1) => use two pointer to find the intersection

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := 0, 0
	first, second := headA, headB
	for first != nil {
		l1++
		first = first.Next
	}
	for second != nil {
		l2++
		second = second.Next
	}

	absLen := abs(l1 - l2)
	first, second = headA, headB
	if l1 > l2 {
		for range absLen {
			first = first.Next
		}
	} else {
		for range absLen {
			second = second.Next
		}
	}

	for first != nil && second != nil {
		if first == second {
			return first
		}
		first = first.Next
		second = second.Next
	}

	return nil
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
