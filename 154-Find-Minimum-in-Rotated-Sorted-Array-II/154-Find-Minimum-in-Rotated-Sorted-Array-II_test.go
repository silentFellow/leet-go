package leetcode

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{nums: []int{1, 3, 5}},
			want: 1,
		},
		{
			name: "Test case 2",
			args: args{nums: []int{2, 2, 2, 0, 1}},
			want: 0,
		},
		{
			name: "Test case 3",
			args: args{nums: []int{4, 5, 6, 7, 0, 1, 4}},
			want: 0,
		},
		{
			name: "Test case 4",
			args: args{nums: []int{10, 1, 10, 10, 10}},
			want: 1,
		},
		{
			name: "Test case 5",
			args: args{nums: []int{3, 3, 1, 3}},
			want: 1,
		},
		{
			name: "Test case 6",
			args: args{nums: []int{3, 1, 3}},
			want: 1,
		},
		{
			name: "Test case 7",
			args: args{nums: []int{1, 2, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
