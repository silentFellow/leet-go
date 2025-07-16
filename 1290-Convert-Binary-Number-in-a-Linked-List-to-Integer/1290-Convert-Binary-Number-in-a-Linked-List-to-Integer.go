package leetcode

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	ans := 0
	var pos float64 = 0
	for prev != nil {
		ans += prev.Val * int(math.Pow(2, pos))
		prev = prev.Next
		pos++
	}

	return ans
}
