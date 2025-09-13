package leetcode

import "testing"

func Test_countPairs(t *testing.T) {
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
			args: args{nums: []int{-1, 1, 2, 3, 1}, target: 2},
			want: 3,
		},
		{
			name: "example2",
			args: args{nums: []int{-6, 2, 5, -2, -7, -1, 3}, target: -2},
			want: 10,
		},
		{
			name: "all pairs valid",
			args: args{nums: []int{-10, -9, -8}, target: 0},
			want: 3,
		},
		{
			name: "no pairs valid",
			args: args{nums: []int{10, 20, 30}, target: 0},
			want: 0,
		},
		{
			name: "single element",
			args: args{nums: []int{1}, target: 2},
			want: 0,
		},
		{
			name: "duplicates",
			args: args{nums: []int{1, 1, 1, 1}, target: 3},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
