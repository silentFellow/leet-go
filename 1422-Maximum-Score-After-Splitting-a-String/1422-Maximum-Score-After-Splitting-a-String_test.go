package leetcode

import "testing"

func Test_maxScore(t *testing.T) {
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
			args: args{s: "011101"},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{s: "00111"},
			want: 5,
		},
		{
			name: "Example 3",
			args: args{s: "1111"},
			want: 3,
		},
		{
			name: "All zeros",
			args: args{s: "0000"},
			want: 3,
		},
		{
			name: "All ones",
			args: args{s: "1111"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.s); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
