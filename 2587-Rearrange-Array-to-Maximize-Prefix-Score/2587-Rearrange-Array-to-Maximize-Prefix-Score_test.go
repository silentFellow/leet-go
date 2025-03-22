package leetcode

import "testing"

func Test_maxScore(t *testing.T) {
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
			args: args{nums: []int{2, -1, 0, 1, -3, 3, -3}},
			want: 6,
		},
		{
			name: "example2",
			args: args{nums: []int{-2, -3, 0}},
			want: 0,
		},
		{
			name: "all_positive",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 5,
		},
		{
			name: "all_negative",
			args: args{nums: []int{-1, -2, -3, -4, -5}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.nums); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
