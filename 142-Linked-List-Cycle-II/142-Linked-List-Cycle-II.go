package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return getLoopPoint(head, slow)
		}
	}

	return nil
}

func getLoopPoint(head, slow *ListNode) *ListNode {
	if head == slow {
		return head
	}

	for head != slow {
		head = head.Next
		slow = slow.Next
	}

  return head
}
