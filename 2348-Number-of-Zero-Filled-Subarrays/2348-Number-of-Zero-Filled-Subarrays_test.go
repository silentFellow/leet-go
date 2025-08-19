package leetcode

import "testing"

func Test_zeroFilledSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{nums: []int{1, 3, 0, 0, 2, 0, 0, 4}},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{nums: []int{0, 0, 0, 2, 0, 0}},
			want: 9,
		},
		{
			name: "Example 3",
			args: args{nums: []int{2, 10, 2019}},
			want: 0,
		},
		{
			name: "All zeros",
			args: args{nums: []int{0, 0, 0, 0}},
			want: 10,
		},
		{
			name: "Single zero",
			args: args{nums: []int{0}},
			want: 1,
		},
		{
			name: "No zeros",
			args: args{nums: []int{1, 2, 3}},
			want: 0,
		},
		{
			name: "Zeros at start and end",
			args: args{nums: []int{0, 0, 1, 0, 0}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zeroFilledSubarray(tt.args.nums); got != tt.want {
				t.Errorf("zeroFilledSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
