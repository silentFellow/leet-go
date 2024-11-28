package leetcode

import "testing"

func Test_arraySign(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "all positive numbers",
			args: args{nums: []int{1, 2, 3, 4}},
			want: 1,
		},
		{
			name: "all negative numbers",
			args: args{nums: []int{-1, -2, -3, -4}},
			want: 1,
		},
		{
			name: "mixed positive and negative numbers",
			args: args{nums: []int{-1, 2, -3, 4}},
			want: 1,
		},
		{
			name: "contains zero",
			args: args{nums: []int{1, 0, -3, 4}},
			want: 0,
		},
		{
			name: "odd number of negative numbers",
			args: args{nums: []int{-1, 1, -1, 1, -1}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arraySign(tt.args.nums); got != tt.want {
				t.Errorf("arraySign() = %v, want %v", got, tt.want)
			}
		})
	}
}
