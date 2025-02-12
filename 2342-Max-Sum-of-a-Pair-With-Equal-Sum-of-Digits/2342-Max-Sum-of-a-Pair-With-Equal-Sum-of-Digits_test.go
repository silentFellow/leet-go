package leetcode

import "testing"

func Test_maximumSum(t *testing.T) {
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
			args: args{nums: []int{18, 43, 36, 13, 7}},
			want: 54,
		},
		{
			name: "example2",
			args: args{nums: []int{10, 12, 19, 14}},
			want: -1,
		},
		{
			name: "no_pairs",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: -1,
		},
		{
			name: "multiple_pairs",
			args: args{nums: []int{51, 42, 33, 24, 15}},
			want: 93, // 51 + 42
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSum(tt.args.nums); got != tt.want {
				t.Errorf("maximumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
