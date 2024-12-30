package leetcode

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "cycle at node index 1",
			args: args{head: func() *ListNode {
				n1 := &ListNode{Val: 3}
				n2 := &ListNode{Val: 2}
				n3 := &ListNode{Val: 0}
				n4 := &ListNode{Val: -4}
				n1.Next = n2
				n2.Next = n3
				n3.Next = n4
				n4.Next = n2
				return n1
			}()},
			want: func() *ListNode {
				n1 := &ListNode{Val: 3}
				n2 := &ListNode{Val: 2}
				n3 := &ListNode{Val: 0}
				n4 := &ListNode{Val: -4}
				n1.Next = n2
				n2.Next = n3
				n3.Next = n4
				n4.Next = n2
				return n2
			}(),
		},
		{
			name: "cycle at node index 0",
			args: args{head: func() *ListNode {
				n1 := &ListNode{Val: 1}
				n2 := &ListNode{Val: 2}
				n1.Next = n2
				n2.Next = n1
				return n1
			}()},
			want: func() *ListNode {
				n1 := &ListNode{Val: 1}
				n2 := &ListNode{Val: 2}
				n1.Next = n2
				n2.Next = n1
				return n1
			}(),
		},
		{
			name: "no cycle",
			args: args{head: func() *ListNode {
				n1 := &ListNode{Val: 1}
				return n1
			}()},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
