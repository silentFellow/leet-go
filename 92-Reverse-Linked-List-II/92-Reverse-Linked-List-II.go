package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	var leftModifier, rightModifier *ListNode

	dummy := &ListNode{Next: head}
	leftModifier = dummy

	for range left-1 {
		leftModifier = leftModifier.Next
	}

	rightModifier = leftModifier.Next

	var prev, next *ListNode
	cur := rightModifier
	for range (right-left+1) {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	leftModifier.Next = prev
	rightModifier.Next = cur

	return dummy.Next
}
