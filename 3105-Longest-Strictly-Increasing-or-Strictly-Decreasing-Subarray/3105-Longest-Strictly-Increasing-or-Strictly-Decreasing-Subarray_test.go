package leetcode

import "testing"

func Test_longestMonotonicSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{nums: []int{1, 4, 3, 3, 2}},
			want: 2,
		},
		{
			name: "example 2",
			args: args{nums: []int{3, 3, 3, 3}},
			want: 1,
		},
		{
			name: "example 3",
			args: args{nums: []int{3, 2, 1}},
			want: 3,
		},
		{
			name: "single element",
			args: args{nums: []int{5}},
			want: 1,
		},
		{
			name: "all increasing",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 5,
		},
		{
			name: "all decreasing",
			args: args{nums: []int{5, 4, 3, 2, 1}},
			want: 5,
		},
		{
			name: "mixed",
			args: args{nums: []int{1, 3, 2, 4, 3, 5}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestMonotonicSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestMonotonicSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
