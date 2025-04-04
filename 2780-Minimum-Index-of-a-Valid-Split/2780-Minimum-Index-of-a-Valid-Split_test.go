package leetcode

import "testing"

func Test_minimumIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{1, 2, 2, 2}},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{nums: []int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}},
			want: 4,
		},
		{
			name: "Example 3",
			args: args{nums: []int{3, 3, 3, 3, 7, 2, 2}},
			want: -1,
		},
		{
			name: "Single element array",
			args: args{nums: []int{1}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumIndex(tt.args.nums); got != tt.want {
				t.Errorf("minimumIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
