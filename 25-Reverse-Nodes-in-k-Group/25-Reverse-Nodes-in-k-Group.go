package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	length := findLength(head)
	maxIteration := length / k

	var prev, cur, next *ListNode = nil, head, nil

	for range maxIteration {
    var lastNodePrev, lastNodeSub *ListNode = prev, cur

		for range k {
			next = cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}

    if lastNodePrev == nil {
      head = prev
    } else {
      lastNodePrev.Next = prev
    }

    lastNodeSub.Next = cur
    prev = lastNodeSub
	}

	return head
}

func findLength(head *ListNode) int {
	temp := head
	length := 0

	for temp != nil {
		temp = temp.Next
		length++
	}

	return length
}
