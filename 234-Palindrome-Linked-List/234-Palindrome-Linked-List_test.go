package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Palindrome with even number of elements",
			args: args{head: &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}},
			want: true,
		},
		{
			name: "Palindrome with odd number of elements",
			args: args{head: &ListNode{1, &ListNode{2, &ListNode{1, nil}}}},
			want: true,
		},
		{
			name: "Not a palindrome",
			args: args{head: &ListNode{1, &ListNode{2, nil}}},
			want: false,
		},
		{
			name: "Single element",
			args: args{head: &ListNode{1, nil}},
			want: true,
		},
		{
			name: "Empty list",
			args: args{head: nil},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
