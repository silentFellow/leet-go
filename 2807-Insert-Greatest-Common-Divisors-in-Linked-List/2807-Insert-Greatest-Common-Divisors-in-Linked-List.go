package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func gcd(m, n int) int {
	if n == 0 {
		return m
	}

	return gcd(n, m%n)
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	first := head
	second := head.Next

	for first != nil && second != nil {
		val := gcd(first.Val, second.Val)
		middle := ListNode{
			Val:  val,
			Next: second,
		}
		first.Next = &middle

    first = second
    second = second.Next
	}

	return head
}
