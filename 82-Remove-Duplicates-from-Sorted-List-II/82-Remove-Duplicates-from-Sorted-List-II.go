package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummyHead := dummy

	l, r := head, head
	for l != nil {
		for r != nil && l.Val == r.Val {
			r = r.Next
		}

		if l.Next == r {
			dummyHead.Next = l
			dummyHead = dummyHead.Next
			dummyHead.Next = nil // since, we use l directly here, break the connection of l.Next with dummyHead.Next
		}
		l = r
	}

	return dummy.Next
}
