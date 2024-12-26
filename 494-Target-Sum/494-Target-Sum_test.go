package leetcode

import "testing"

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				nums:   []int{1, 1, 1, 1, 1},
				target: 3,
			},
			want: 5,
		},
		{
			name: "example2",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 1,
		},
		{
			name: "example3",
			args: args{
				nums:   []int{1, 2, 1},
				target: 2,
			},
			want: 2,
		},
		{
			name: "example4",
			args: args{
				nums:   []int{2, 3, 5, 7},
				target: 8,
			},
			want: 0,
		},
		{
			name: "example5",
			args: args{
				nums:   []int{0, 0, 0, 0, 0},
				target: 0,
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
