package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeNodes(t *testing.T) {
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
				head: &ListNode{5, &ListNode{2, &ListNode{13, &ListNode{3, &ListNode{8, nil}}}}},
			},
			want: &ListNode{13, &ListNode{8, nil}},
		},
		{
			name: "example 2",
			args: args{head: &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, nil}}}}},
			want: &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, nil}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
