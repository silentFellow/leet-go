package leetcode

import "testing"

func Test_hasIncreasingSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{nums: []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, k: 3},
			want: true,
		},
		{
			name: "Example 2",
			args: args{nums: []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}, k: 5},
			want: false,
		},
		{
			name: "No increasing subarrays",
			args: args{nums: []int{5, 4, 3, 2, 1, 0}, k: 2},
			want: false,
		},
		{
			name: "Adjacent at start",
			args: args{nums: []int{1, 2, 3, 4, 5}, k: 2},
			want: true,
		},
		{
			name: "All elements same",
			args: args{nums: []int{1, 1, 1, 1, 1, 1}, k: 2},
			want: false,
		},
		{
			name: "Edge case k=2",
			args: args{nums: []int{1, 2, 3, 4}, k: 2},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasIncreasingSubarrays(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("hasIncreasingSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
