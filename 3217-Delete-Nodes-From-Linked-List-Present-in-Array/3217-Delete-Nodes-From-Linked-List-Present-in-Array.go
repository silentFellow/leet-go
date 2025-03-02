package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	hmap := make(map[int]struct{})
	for _, v := range nums {
		hmap[v] = struct{}{}
	}

	for head != nil {
		if _, ok := hmap[head.Val]; ok {
			head = head.Next
		} else {
			break
		}
	}

	if head == nil || head.Next == nil {
		return head
	}

	ans, cur, next := head, head, head.Next
	for next != nil {
		for next != nil {
			if _, ok := hmap[next.Val]; ok {
				next = next.Next
			} else {
				break
			}
		}

		cur.Next = next
		if cur != nil {
			cur = cur.Next
		}
		if next != nil {
			next = next.Next
		}
	}

	return ans
}
