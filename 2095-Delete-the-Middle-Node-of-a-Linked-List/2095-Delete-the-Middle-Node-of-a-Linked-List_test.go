package leetcode

import (
	"reflect"
	"testing"
)

func Test_deleteMiddle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{head: &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{7, &ListNode{1, &ListNode{2, &ListNode{6, nil}}}}}}}},
			want: &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{1, &ListNode{2, &ListNode{6, nil}}}}}},
		},
		{
			name: "Example 2",
			args: args{head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}},
			want: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		},
		{
			name: "Example 3",
			args: args{head: &ListNode{2, &ListNode{1, nil}}},
			want: &ListNode{2, nil},
		},
		{
			name: "Single node",
			args: args{head: &ListNode{1, nil}},
			want: nil,
		},
		{
			name: "Two nodes",
			args: args{head: &ListNode{1, &ListNode{2, nil}}},
			want: &ListNode{1, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteMiddle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}
