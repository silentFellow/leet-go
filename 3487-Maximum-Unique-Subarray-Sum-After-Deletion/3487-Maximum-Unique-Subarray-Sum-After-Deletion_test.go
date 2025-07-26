package leetcode

import "testing"

func Test_maxSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "all positive unique",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 15,
		},
		{
			name: "all duplicates and zero",
			args: args{nums: []int{1, 1, 0, 1, 1}},
			want: 1,
		},
		{
			name: "mixed with negatives",
			args: args{nums: []int{1, 2, -1, -2, 1, 0, -1}},
			want: 3,
		},
		{
			name: "all negatives",
			args: args{nums: []int{-100}},
			want: -100,
		},
		{
			name: "positive and negative",
			args: args{nums: []int{20, -20}},
			want: 20,
		},
		{
			name: "all zeros",
			args: args{nums: []int{0, 0, 0}},
			want: 0,
		},
		{
			name: "mixed, all unique",
			args: args{nums: []int{5, -2, 3, 0}},
			want: 8,
		},
		{
			name: "mixed, some duplicates",
			args: args{nums: []int{2, 2, -1, 3, 3, 0}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSum(tt.args.nums); got != tt.want {
				t.Errorf("maxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
