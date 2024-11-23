package leetcode

import (
	"reflect"
	"testing"
)

func Test_insertGreatestCommonDivisors(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{head: &ListNode{18, &ListNode{6, &ListNode{10, &ListNode{3, nil}}}}},
			want: &ListNode{
				18,
				&ListNode{
					6,
					&ListNode{6, &ListNode{2, &ListNode{10, &ListNode{1, &ListNode{3, nil}}}}},
				},
			},
		},
		{
			name: "test2",
			args: args{head: &ListNode{7, nil}},
			want: &ListNode{7, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertGreatestCommonDivisors(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertGreatestCommonDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
