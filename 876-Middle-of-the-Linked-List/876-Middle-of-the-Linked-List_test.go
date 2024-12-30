package leetcode

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test case 1",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			},
			want: &ListNode{3, &ListNode{4, &ListNode{5, nil}}},
		},
		{
			name: "Test case 2",
			args: args{
				head: &ListNode{
					1,
					&ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}},
				},
			},
			want: &ListNode{4, &ListNode{5, &ListNode{6, nil}}},
		},
		{
			name: "Test case 3",
			args: args{head: &ListNode{1, nil}},
			want: &ListNode{1, nil},
		},
		{
			name: "Test case 4",
			args: args{head: &ListNode{1, &ListNode{2, nil}}},
			want: &ListNode{2, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
