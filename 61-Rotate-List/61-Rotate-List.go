package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	length := 0
	temp := head
	for temp != nil {
		length++
		temp = temp.Next
	}

	if head == nil || k%length == 0 {
		return head
	}
	k = k % length

	slow, fast := head, head
	for range k {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	temp = slow.Next
	slow.Next = nil
	fast.Next = head

	return temp
}
