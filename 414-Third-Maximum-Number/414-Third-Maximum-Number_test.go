package leetcode

import "testing"

func Test_thirdMax(t *testing.T) {
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
			args: args{nums: []int{3, 2, 1}},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{nums: []int{1, 2}},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{nums: []int{2, 2, 3, 1}},
			want: 1,
		},
		{
			name: "Single element",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "Two elements",
			args: args{nums: []int{1, 2}},
			want: 2,
		},
		{
			name: "All elements same",
			args: args{nums: []int{2, 2, 2}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thirdMax(tt.args.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
