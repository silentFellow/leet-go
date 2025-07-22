package leetcode

import "testing"

func Test_maximumUniqueSubarray(t *testing.T) {
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
			args: args{nums: []int{4, 2, 4, 5, 6}},
			want: 17,
		},
		{
			name: "example2",
			args: args{nums: []int{5, 2, 1, 2, 5, 2, 1, 2, 5}},
			want: 8,
		},
		{
			name: "all_same",
			args: args{nums: []int{1, 1, 1, 1}},
			want: 1,
		},
		{
			name: "all_unique",
			args: args{nums: []int{1, 2, 3, 4}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumUniqueSubarray(tt.args.nums); got != tt.want {
				t.Errorf("maximumUniqueSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
