package leetcode

import "testing"

func Test_maximumCount(t *testing.T) {
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
			args: args{nums: []int{-2, -1, -1, 1, 2, 3}},
			want: 3,
		},
		{
			name: "example2",
			args: args{nums: []int{-3, -2, -1, 0, 0, 1, 2}},
			want: 3,
		},
		{
			name: "example3",
			args: args{nums: []int{5, 20, 66, 1314}},
			want: 4,
		},
		{
			name: "all negatives",
			args: args{nums: []int{-5, -4, -3, -2, -1}},
			want: 5,
		},
		{
			name: "all positives",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 5,
		},
		{
			name: "mixed with zeros",
			args: args{nums: []int{-3, -2, -1, 0, 0, 1, 2, 3}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCount(tt.args.nums); got != tt.want {
				t.Errorf("maximumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
