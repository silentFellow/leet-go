package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				head: &ListNode{
					1,
					&ListNode{
						2,
						&ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}},
					},
				},
				val: 6,
			},
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		},
		{
			name: "Example 2",
			args: args{
				head: nil,
				val:  1,
			},
			want: nil,
		},
		{
			name: "Example 3",
			args: args{
				head: &ListNode{7, &ListNode{7, &ListNode{7, &ListNode{7, nil}}}},
				val:  7,
			},
			want: nil,
		},
		{
			name: "No elements to remove",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
				val:  4,
			},
			want: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
		{
			name: "Remove head element",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
				val:  1,
			},
			want: &ListNode{2, &ListNode{3, nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
