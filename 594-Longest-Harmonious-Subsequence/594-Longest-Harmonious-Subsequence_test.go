package leetcode

import "testing"

func Test_findLHS(t *testing.T) {
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
			args: args{nums: []int{1, 3, 2, 2, 5, 2, 3, 7}},
			want: 5,
		},
		{
			name: "example2",
			args: args{nums: []int{1, 2, 3, 4}},
			want: 2,
		},
		{
			name: "example3",
			args: args{nums: []int{1, 1, 1, 1}},
			want: 0,
		},
		{
			name: "empty",
			args: args{nums: []int{}},
			want: 0,
		},
		{
			name: "single element",
			args: args{nums: []int{5}},
			want: 0,
		},
		{
			name: "two elements diff 1",
			args: args{nums: []int{2, 3}},
			want: 2,
		},
		{
			name: "two elements diff >1",
			args: args{nums: []int{2, 5}},
			want: 0,
		},
		{
			name: "multiple pairs",
			args: args{nums: []int{1, 2, 2, 1}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLHS(tt.args.nums); got != tt.want {
				t.Errorf("findLHS() = %v, want %v", got, tt.want)
			}
		})
	}
}
