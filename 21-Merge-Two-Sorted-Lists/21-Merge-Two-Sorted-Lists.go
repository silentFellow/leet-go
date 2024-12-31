package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	dummyTail := dummy

	first, second := list1, list2

	for first != nil && second != nil {
		if first.Val < second.Val {
			neo := &ListNode{Val: first.Val}
			dummyTail.Next = neo
			dummyTail = neo
			first = first.Next
		} else {
			neo := &ListNode{Val: second.Val}
			dummyTail.Next = neo
			dummyTail = neo
			second = second.Next
		}
	}

	if first == nil {
		dummyTail.Next = second
	} else {
		dummyTail.Next = first
	}

	return dummy.Next
}
