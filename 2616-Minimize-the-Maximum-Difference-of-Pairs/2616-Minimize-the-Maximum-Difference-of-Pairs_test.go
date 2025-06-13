package leetcode

import "testing"

func Test_minimizeMax(t *testing.T) {
	type args struct {
		nums []int
		p    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{10, 1, 2, 7, 1, 3}, p: 2},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{nums: []int{4, 2, 1, 2}, p: 1},
			want: 0,
		},
		{
			name: "No pairs needed",
			args: args{nums: []int{5, 3, 1, 9, 7}, p: 0},
			want: 0,
		},
		{
			name: "All identical elements",
			args: args{nums: []int{6, 6, 6, 6, 6, 6}, p: 3},
			want: 0,
		},
		{
			name: "Already sorted input",
			args: args{nums: []int{1, 2, 3, 4, 5, 6}, p: 2},
			want: 1,
		},
		{
			name: "Max difference at edges",
			args: args{nums: []int{1, 1000000000}, p: 1},
			want: 999999999,
		},
		{
			name: "Alternating pattern",
			args: args{nums: []int{1, 10, 1, 10, 1, 10}, p: 3},
			want: 9,
		},
		{
			name: "Minimum input size",
			args: args{nums: []int{1, 2}, p: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimizeMax(tt.args.nums, tt.args.p); got != tt.want {
				t.Errorf("minimizeMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
