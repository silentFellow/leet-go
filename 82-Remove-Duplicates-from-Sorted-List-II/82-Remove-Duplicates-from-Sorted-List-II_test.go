package leetcode

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example 1",
			args: args{
				head: &ListNode{
					1,
					&ListNode{
						2,
						&ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}},
					},
				},
			},
			want: &ListNode{1, &ListNode{2, &ListNode{5, nil}}},
		},
		{
			name: "example 2",
			args: args{
				head: &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}},
			},
			want: &ListNode{2, &ListNode{3, nil}},
		},
		{
			name: "no duplicates",
			args: args{head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}}},
			want: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
		{
			name: "all duplicates",
			args: args{head: &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, nil}}}}},
			want: nil,
		},
		{
			name: "single element",
			args: args{head: &ListNode{1, nil}},
			want: &ListNode{1, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
