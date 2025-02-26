package leetcode

import "testing"

func Test_maxAbsoluteSum(t *testing.T) {
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
			args: args{nums: []int{1, -3, 2, 3, -4}},
			want: 5,
		},
		{
			name: "example2",
			args: args{nums: []int{2, -5, 1, -4, 3, -2}},
			want: 8,
		},
		{
			name: "all positive",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 15,
		},
		{
			name: "all negative",
			args: args{nums: []int{-1, -2, -3, -4, -5}},
			want: 15,
		},
		{
			name: "mixed",
			args: args{nums: []int{-1, 2, -3, 4, -5}},
			want: 5,
		},
		{
			name: "single element positive",
			args: args{nums: []int{10}},
			want: 10,
		},
		{
			name: "single element negative",
			args: args{nums: []int{-10}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAbsoluteSum(tt.args.nums); got != tt.want {
				t.Errorf("maxAbsoluteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
