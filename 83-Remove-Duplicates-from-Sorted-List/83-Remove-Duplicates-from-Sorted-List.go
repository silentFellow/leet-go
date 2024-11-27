package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	thead, temp := head, head

	for thead != nil {
		for temp != nil && thead.Val == temp.Val {
			temp = temp.Next
		}
		thead.Next = temp
		thead = thead.Next
	}

	return head
}
