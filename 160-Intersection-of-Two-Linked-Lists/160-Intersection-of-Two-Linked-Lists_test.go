package leetcode

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "intersect at node with value 2",
			args: args{
				headA: &ListNode{1, &ListNode{9, &ListNode{1, nil}}},
				headB: &ListNode{3, nil},
			},
			want: nil,
		},
		{
			name: "no intersection",
			args: args{
				headA: &ListNode{2, &ListNode{6, &ListNode{4, nil}}},
				headB: &ListNode{1, &ListNode{5, nil}},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive value",
			args: args{val: 5},
			want: 5,
		},
		{
			name: "negative value",
			args: args{val: -5},
			want: 5,
		},
		{
			name: "zero value",
			args: args{val: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.val); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
