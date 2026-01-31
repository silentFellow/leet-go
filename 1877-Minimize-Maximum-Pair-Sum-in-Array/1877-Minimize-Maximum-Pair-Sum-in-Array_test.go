package leetcode

import "testing"

func Test_minPairSum(t *testing.T) {
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
			args: args{nums: []int{3, 5, 2, 3}},
			want: 7,
		},
		{
			name: "Example 2",
			args: args{nums: []int{3, 5, 4, 2, 4, 6}},
			want: 8,
		},
		{
			name: "All same elements",
			args: args{nums: []int{4, 4, 4, 4}},
			want: 8,
		},
		{
			name: "Ascending order",
			args: args{nums: []int{1, 2, 3, 4}},
			want: 5,
		},
		{
			name: "Descending order",
			args: args{nums: []int{8, 7, 6, 5, 4, 3, 2, 1}},
			want: 9,
		},
		{
			name: "Large numbers",
			args: args{nums: []int{100000, 1, 99999, 2}},
			want: 100001,
		},
		{
			name: "Minimum input size",
			args: args{nums: []int{1, 2}},
			want: 3,
		},
		{
			name: "Duplicates",
			args: args{nums: []int{5, 5, 5, 5, 5, 5}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPairSum(tt.args.nums); got != tt.want {
				t.Errorf("minPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
