package leetcode

import "testing"

func Test_isOneBitCharacter(t *testing.T) {
	type args struct {
		bits []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{bits: []int{1, 0, 0}},
			want: true,
		},
		{
			name: "Example 2",
			args: args{bits: []int{1, 1, 1, 0}},
			want: false,
		},
		{
			name: "Single zero",
			args: args{bits: []int{0}},
			want: true,
		},
		{
			name: "Ends with two-bit character",
			args: args{bits: []int{1, 0}},
			want: false,
		},
		{
			name: "Multiple one-bit characters",
			args: args{bits: []int{0, 0, 0}},
			want: true,
		},
		{
			name: "Two two-bit characters",
			args: args{bits: []int{1, 1, 0}},
			want: true,
		},
		{
			name: "Long sequence ending with one-bit",
			args: args{bits: []int{1, 1, 0, 0, 1, 0, 0}},
			want: true,
		},
		{
			name: "Long sequence ending with two-bit",
			args: args{bits: []int{1, 1, 0, 0, 1, 1}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOneBitCharacter(tt.args.bits); got != tt.want {
				t.Errorf("isOneBitCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
