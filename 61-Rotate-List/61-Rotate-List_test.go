package leetcode

import (
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
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
			name: "example 1",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
				k:    2,
			},
			want: &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}},
		},
		{
			name: "example 2",
			args: args{
				head: &ListNode{0, &ListNode{1, &ListNode{2, nil}}},
				k:    4,
			},
			want: &ListNode{2, &ListNode{0, &ListNode{1, nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
