package leetcode

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{nums: []int{3, 4, 5, 2}},
			want: 12,
		},
		{
			name: "example 2",
			args: args{nums: []int{1, 5, 4, 5}},
			want: 16,
		},
		{
			name: "example 3",
			args: args{nums: []int{3, 7}},
			want: 12,
		},
		{
			name: "minimum length",
			args: args{nums: []int{1, 2}},
			want: 0,
		},
		{
			name: "same elements",
			args: args{nums: []int{5, 5, 5, 5}},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
