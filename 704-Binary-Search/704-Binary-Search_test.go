package leetcode

import "testing"

func Test_search(t *testing.T) {
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
			name: "target exists in the middle",
			args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9},
			want: 4,
		},
		{
			name: "target does not exist",
			args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2},
			want: -1,
		},
		{
			name: "target is the first element",
			args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: -1},
			want: 0,
		},
		{
			name: "target is the last element",
			args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 12},
			want: 5,
		},
		{
			name: "single element array, target exists",
			args: args{nums: []int{5}, target: 5},
			want: 0,
		},
		{
			name: "single element array, target does not exist",
			args: args{nums: []int{5}, target: 1},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
