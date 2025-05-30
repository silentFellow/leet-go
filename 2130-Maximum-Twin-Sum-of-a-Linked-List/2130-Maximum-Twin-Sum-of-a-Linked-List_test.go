package leetcode

import "testing"

func Test_pairSum(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{head: &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}},
			want: 6,
		},
		{
			name: "test2",
			args: args{head: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}},
			want: 7,
		},
		{
			name: "test3",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 100000}}},
			want: 100001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairSum(tt.args.head); got != tt.want {
				t.Errorf("pairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
