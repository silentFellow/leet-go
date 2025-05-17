package leetcode

import "testing"

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "example1",
			args: args{nums: []int{2, 0, 2, 1, 1, 0}},
		},
		{
			name: "example2",
			args: args{nums: []int{2, 0, 1}},
		},
		{
			name: "all zeros",
			args: args{nums: []int{0, 0, 0}},
		},
		{
			name: "all ones",
			args: args{nums: []int{1, 1, 1}},
		},
		{
			name: "all twos",
			args: args{nums: []int{2, 2, 2}},
		},
		{
			name: "mixed",
			args: args{nums: []int{1, 2, 0, 1, 2, 0, 1}},
		},
		{
			name: "single element 0",
			args: args{nums: []int{0}},
		},
		{
			name: "single element 1",
			args: args{nums: []int{1}},
		},
		{
			name: "single element 2",
			args: args{nums: []int{2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
		})
	}
}
