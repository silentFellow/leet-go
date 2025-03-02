package leetcode

import (
	"reflect"
	"testing"
)

func Test_modifiedList(t *testing.T) {
	type args struct {
		nums []int
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 2, 3},
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			},
			want: &ListNode{4, &ListNode{5, nil}},
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{1},
				head: &ListNode{
					1,
					&ListNode{2, &ListNode{1, &ListNode{2, &ListNode{1, &ListNode{2, nil}}}}},
				},
			},
			want: &ListNode{2, &ListNode{2, &ListNode{2, nil}}},
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{5},
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
			},
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
		},
		{
			name: "No nodes to remove",
			args: args{
				nums: []int{6, 7, 8},
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			},
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		},
		{
			name: "All nodes to remove",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modifiedList(tt.args.nums, tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("modifiedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
