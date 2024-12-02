package leetcode

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
				k:    2,
			},
			want: &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{5, nil}}}}},
		},
		{
			name: "Example 2",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
				k:    3,
			},
			want: &ListNode{3, &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{5, nil}}}}},
		},
		{
			name: "Single node",
			args: args{
				head: &ListNode{1, nil},
				k:    1,
			},
			want: &ListNode{1, nil},
		},
		{
			name: "k equals list length",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
				k:    3,
			},
			want: &ListNode{3, &ListNode{2, &ListNode{1, nil}}},
		},
		{
			name: "k greater than list length",
			args: args{
				head: &ListNode{1, &ListNode{2, nil}},
				k:    3,
			},
			want: &ListNode{1, &ListNode{2, nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
