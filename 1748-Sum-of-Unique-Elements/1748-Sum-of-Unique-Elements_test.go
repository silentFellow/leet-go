package leetcode

import "testing"

func Test_sumOfUnique(t *testing.T) {
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
			args: args{nums: []int{1, 2, 3, 2}},
			want: 4,
		},
		{
			name: "example 2",
			args: args{nums: []int{1, 1, 1, 1, 1}},
			want: 0,
		},
		{
			name: "example 3",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 15,
		},
		{
			name: "no elements",
			args: args{nums: []int{}},
			want: 0,
		},
		{
			name: "all unique",
			args: args{nums: []int{10, 20, 30}},
			want: 60,
		},
		{
			name: "all duplicates",
			args: args{nums: []int{5, 5, 5, 5}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfUnique(tt.args.nums); got != tt.want {
				t.Errorf("sumOfUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
