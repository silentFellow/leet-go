package leetcode

import "testing"

func Test_maxOperations(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{s: "1001101"},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{s: "00111"},
			want: 0,
		},
		{
			name: "All ones",
			args: args{s: "11111"},
			want: 0,
		},
		{
			name: "All zeros",
			args: args{s: "00000"},
			want: 0,
		},
		{
			name: "Single 1 at start",
			args: args{s: "100000"},
			want: 1,
		},
		{
			name: "Single 1 at end",
			args: args{s: "000001"},
			want: 0,
		},
		{
			name: "Multiple 1s and 0s",
			args: args{s: "110011"},
			want: 2,
		},
		{
			name: "Long sequence of 1s then 0s",
			args: args{s: "1111100000"},
			want: 5,
		},
		{
			name: "Long sequence of 0s then 1s",
			args: args{s: "0000011111"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxOperations(tt.args.s); got != tt.want {
				t.Errorf("maxOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
