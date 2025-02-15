package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow.Next
	slow.Next = nil

	left := sortList(head)
	right := sortList(mid)

	return merge(left, right)
}

func merge(f, s *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	dummyTail := dummy

	for f != nil && s != nil {
		if f.Val < s.Val {
			dummyTail.Next = f
			dummyTail = dummyTail.Next
			f = f.Next
		} else {
			dummyTail.Next = s
			dummyTail = dummyTail.Next
			s = s.Next
		}
	}

	if f == nil {
		dummyTail.Next = s
	} else {
		dummyTail.Next = f
	}

	return dummy.Next
}
