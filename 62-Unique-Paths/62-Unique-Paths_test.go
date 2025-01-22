package leetcode

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2x2 grid",
			args: args{m: 2, n: 2},
			want: 2,
		},
		{
			name: "3x3 grid",
			args: args{m: 3, n: 3},
			want: 6,
		},
		{
			name: "3x2 grid",
			args: args{m: 3, n: 2},
			want: 3,
		},
		{
			name: "7x3 grid",
			args: args{m: 7, n: 3},
			want: 28,
		},
		{
			name: "1x1 grid",
			args: args{m: 1, n: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
