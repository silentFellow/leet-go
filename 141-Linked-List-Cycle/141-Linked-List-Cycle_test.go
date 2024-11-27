package leetcode

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "no cycle, single node",
			args: args{head: &ListNode{Val: 1}},
			want: false,
		},
		{
			name: "no cycle, multiple nodes",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}},
			want: false,
		},
		{
			name: "cycle, single node pointing to itself",
			args: args{head: func() *ListNode {
				node := &ListNode{Val: 1}
				node.Next = node
				return node
			}()},
			want: true,
		},
		{
			name: "cycle, multiple nodes",
			args: args{head: func() *ListNode {
				node1 := &ListNode{Val: 1}
				node2 := &ListNode{Val: 2}
				node3 := &ListNode{Val: 3}
				node1.Next = node2
				node2.Next = node3
				node3.Next = node2
				return node1
			}()},
			want: true,
		},
		{
			name: "no cycle, empty list",
			args: args{head: nil},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
