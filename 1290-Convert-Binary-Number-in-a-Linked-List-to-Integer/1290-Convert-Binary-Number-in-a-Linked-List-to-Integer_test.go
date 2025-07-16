package leetcode

import "testing"

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{head: &ListNode{1, &ListNode{0, &ListNode{1, nil}}}},
			want: 5,
		},
		{
			name: "single zero",
			args: args{head: &ListNode{0, nil}},
			want: 0,
		},
		{
			name: "single one",
			args: args{head: &ListNode{1, nil}},
			want: 1,
		},
		{
			name: "long binary",
			args: args{
				head: &ListNode{
					1,
					&ListNode{
						0,
						&ListNode{1, &ListNode{1, &ListNode{1, &ListNode{0, &ListNode{1, nil}}}}},
					},
				},
			},
			want: 93,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
