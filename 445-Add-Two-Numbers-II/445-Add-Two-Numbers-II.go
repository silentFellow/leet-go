package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)

	carry := 0
	dummy := new(ListNode)
	dummyTemp := dummy
	for l1 != nil || l2 != nil {
		var l1v, l2v int

		if l1 != nil {
			l1v = l1.Val
		}

		if l2 != nil {
			l2v = l2.Val
		}

		sum := l1v + l2v + carry

		carry = sum / 10
		dummyTemp.Next = &ListNode{Val: sum % 10}
		dummyTemp = dummyTemp.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

  if carry != 0 {
    dummyTemp.Next = &ListNode{Val: carry}
  }

	return reverse(dummy.Next)
}

func reverse(head *ListNode) *ListNode {
	var prev, next *ListNode = nil, nil
	cur := head

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}
