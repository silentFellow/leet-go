package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	head := list1
	for range a - 1 {
		head = head.Next
	}

	tail := head
	for range b - a + 2 {
		tail = tail.Next
	}

	head.Next = list2
	for head.Next != nil {
		head = head.Next
	}

	head.Next = tail

	return list1
}
