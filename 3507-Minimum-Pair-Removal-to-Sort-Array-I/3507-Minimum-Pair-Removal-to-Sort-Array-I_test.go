package leetcode

import "testing"

func Test_minimumPairRemoval(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{nums: []int{5, 2, 3, 1}},
			want: 2,
		},
		{
			name: "example2",
			args: args{nums: []int{1, 2, 2}},
			want: 0,
		},
		{
			name: "already sorted",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 0,
		},
		{
			name: "all decreasing",
			args: args{nums: []int{5, 4, 3, 2, 1}},
			want: 4,
		},
		{
			name: "negative numbers",
			args: args{nums: []int{-1, -2, -3, -4}},
			want: 3,
		},
		{
			name: "single element",
			args: args{nums: []int{7}},
			want: 0,
		},
		{
			name: "two elements sorted",
			args: args{nums: []int{1, 2}},
			want: 0,
		},
		{
			name: "two elements unsorted",
			args: args{nums: []int{2, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumPairRemoval(tt.args.nums); got != tt.want {
				t.Errorf("minimumPairRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}
