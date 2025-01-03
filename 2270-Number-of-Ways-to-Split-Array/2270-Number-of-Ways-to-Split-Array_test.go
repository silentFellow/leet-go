package leetcode

import "testing"

func Test_waysToSplitArray(t *testing.T) {
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
			args: args{nums: []int{10, 4, -8, 7}},
			want: 2,
		},
		{
			name: "example2",
			args: args{nums: []int{2, 3, 1, 0}},
			want: 2,
		},
		{
			name: "all_negative",
			args: args{nums: []int{-1, -2, -3, -4, -5}},
			want: 3,
		},
		{
			name: "mixed",
			args: args{nums: []int{1, -1, 1, -1, 1}},
			want: 2,
		},
		{
			name: "single_split",
			args: args{nums: []int{1, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToSplitArray(tt.args.nums); got != tt.want {
				t.Errorf("waysToSplitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
