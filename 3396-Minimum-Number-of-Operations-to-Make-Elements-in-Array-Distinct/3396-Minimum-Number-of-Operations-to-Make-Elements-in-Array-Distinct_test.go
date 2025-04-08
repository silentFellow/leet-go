package leetcode

import "testing"

func Test_minimumOperations(t *testing.T) {
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
			args: args{nums: []int{1, 2, 3, 4, 2, 3, 3, 5, 7}},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{nums: []int{4, 5, 6, 4, 4}},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{nums: []int{6, 7, 8, 9}},
			want: 0,
		},
		{
			name: "All elements distinct",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 0,
		},
		{
			name: "All elements the same",
			args: args{nums: []int{1, 1, 1, 1, 1}},
			want: 2,
		},
		{
			name: "Empty array",
			args: args{nums: []int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.args.nums); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
