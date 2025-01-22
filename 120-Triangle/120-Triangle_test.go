package leetcode

import "testing"

func Test_minimumTotal(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{triangle: [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			}},
			want: 11,
		},
		{
			name: "example2",
			args: args{triangle: [][]int{
				{-10},
			}},
			want: -10,
		},
		{
			name: "example3",
			args: args{triangle: [][]int{
				{1},
				{2, 3},
				{3, 4, 5},
				{6, 7, 8, 9},
			}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotal(tt.args.triangle); got != tt.want {
				t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
