package leetcode

import "testing"

func Test_minSubarray(t *testing.T) {
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
			args: args{nums: []int{3, 1, 4, 2}, p: 6},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{nums: []int{6, 3, 5, 2}, p: 9},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{nums: []int{1, 2, 3}, p: 3},
			want: 0,
		},
		{
			name: "Impossible case",
			args: args{nums: []int{2, 3, 1}, p: 7},
			want: -1,
		},
		{
			name: "Large p, no removal needed",
			args: args{nums: []int{1000000000, 1, 1, 1, 1}, p: 1000000004},
			want: 0,
		},
		{
			name: "All elements same, p divides sum",
			args: args{nums: []int{2, 2, 2, 2}, p: 8},
			want: 0,
		},
		{
			name: "All elements same, p does not divide sum",
			args: args{nums: []int{2, 2, 2, 2}, p: 3},
			want: 1,
		},
		{
			name: "Subarray at end",
			args: args{nums: []int{1, 2, 3, 4, 5}, p: 6},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubarray(tt.args.nums, tt.args.p); got != tt.want {
				t.Errorf("minSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
