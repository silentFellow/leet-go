package leetcode

import "testing"

func Test_findScore(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "example1",
			args: args{nums: []int{2, 1, 3, 4, 5, 2}},
			want: 7,
		},
		{
			name: "example2",
			args: args{nums: []int{2, 3, 5, 1, 3, 2}},
			want: 5,
		},
		{
			name: "single element",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "already sorted",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 9,
		},
		{
			name: "reverse sorted",
			args: args{nums: []int{5, 4, 3, 2, 1}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findScore(tt.args.nums); got != tt.want {
				t.Errorf("findScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
