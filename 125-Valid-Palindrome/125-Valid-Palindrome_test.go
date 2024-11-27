package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{s: "A man, a plan, a canal: Panama"},
			want: true,
		},
		{
			name: "Example 2",
			args: args{s: "race a car"},
			want: false,
		},
		{
			name: "Empty string",
			args: args{s: ""},
			want: true,
		},
		{
			name: "Single character",
			args: args{s: "a"},
			want: true,
		},
		{
			name: "Palindrome with numbers",
			args: args{s: "12321"},
			want: true,
		},
		{
			name: "Non-palindrome with numbers",
			args: args{s: "12345"},
			want: false,
		},
		{
			name: "Mixed alphanumeric palindrome",
			args: args{s: "A1b2B1a"},
			want: true,
		},
		{
			name: "Mixed alphanumeric non-palindrome",
			args: args{s: "A1b2C3"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
