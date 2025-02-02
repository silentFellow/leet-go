package leetcode

import "testing"

func Test_check(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example_1",
			args: args{nums: []int{3, 2, 1}},
			want: false,
		},
		{
			name: "sorted and rotated",
			args: args{nums: []int{3, 4, 5, 1, 2}},
			want: true,
		},
		{
			name: "sorted but not rotated",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: true,
		},
		{
			name: "not sorted",
			args: args{nums: []int{2, 1, 3, 4, 5}},
			want: false,
		},
		{
			name: "single element",
			args: args{nums: []int{1}},
			want: true,
		},
		{
			name: "two elements sorted and rotated",
			args: args{nums: []int{2, 1}},
			want: true,
		},
		{
			name: "two elements not sorted",
			args: args{nums: []int{1, 3}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.nums); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
