package leetcode

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{1, 2, 3, 3, 2, 2}},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{nums: []int{1, 2, 3, 4}},
			want: 1,
		},
		{
			name: "All same max value",
			args: args{nums: []int{3, 3, 3, 3}},
			want: 4,
		},
		{
			name: "Multiple groups of max value",
			args: args{nums: []int{3, 3, 3, 2, 3, 3}},
			want: 3,
		},
		{
			name: "Max at end",
			args: args{nums: []int{1, 2, 2, 4, 4, 4}},
			want: 3,
		},
		{
			name: "Max at beginning",
			args: args{nums: []int{5, 5, 1, 2, 3}},
			want: 2,
		},
		{
			name: "Non-contiguous max values",
			args: args{nums: []int{7, 1, 7, 2, 7}},
			want: 1,
		},
		{
			name: "One element",
			args: args{nums: []int{99}},
			want: 1,
		},
		{
			name: "Large identical elements",
			args: args{nums: make([]int, 100000)},
			want: 100000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
