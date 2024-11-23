package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	dummy := new(ListNode)
	temp := dummy

	sum := 0
	for head != nil {
		if head.Val != 0 {
			sum += head.Val
		} else {
			if sum != 0 {
				temp.Next = &ListNode{Val: sum}
				temp = temp.Next
			}
			sum = 0
		}

    head = head.Next
	}

	return dummy.Next
}
