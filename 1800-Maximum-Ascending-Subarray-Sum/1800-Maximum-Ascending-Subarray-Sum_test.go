package leetcode

import "testing"

func Test_maxAscendingSum(t *testing.T) {
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
			args: args{nums: []int{10, 20, 30, 5, 10, 50}},
			want: 65,
		},
		{
			name: "example2",
			args: args{nums: []int{10, 20, 30, 40, 50}},
			want: 150,
		},
		{
			name: "example3",
			args: args{nums: []int{12, 17, 15, 13, 10, 11, 12}},
			want: 33,
		},
		{
			name: "single element",
			args: args{nums: []int{5}},
			want: 5,
		},
		{
			name: "all descending",
			args: args{nums: []int{5, 4, 3, 2, 1}},
			want: 5,
		},
		{
			name: "mixed",
			args: args{nums: []int{1, 2, 3, 2, 3, 4, 1, 2}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAscendingSum(tt.args.nums); got != tt.want {
				t.Errorf("maxAscendingSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
