package leetcode

import "testing"

func Test_deleteNode(t *testing.T) {
	type args struct {
		node *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "delete node with value 5",
			args: args{node: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9}}}},
			want: []int{4, 1, 9},
		},
		{
			name: "delete node with value 1",
			args: args{node: &ListNode{Val: 1, Next: &ListNode{Val: 9}}},
			want: []int{4, 5, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9}}}}
			// Find the node to delete
			node := head
			for node != nil && node.Val != tt.args.node.Val {
				node = node.Next
			}
			deleteNode(node)
			got := listToSlice(head)
			if !equal(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
