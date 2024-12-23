package leetcode

import "testing"

func Test_removeDuplicates(t *testing.T) {
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
			args: args{nums: []int{1, 1, 2}},
			want: 2,
		},
		{
			name: "Test case 2",
			args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			want: 5,
		},
		{
			name: "Test case 3",
			args: args{nums: []int{1, 1, 1, 1, 1}},
			want: 1,
		},
		{
			name: "Test case 4",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 5,
		},
		{
			name: "Test case 5",
			args: args{nums: []int{1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
