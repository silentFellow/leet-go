package leetcode

import "testing"

func Test_tupleSameProduct(t *testing.T) {
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
			args: args{nums: []int{2, 3, 4, 6}},
			want: 8,
		},
		{
			name: "example2",
			args: args{nums: []int{1, 2, 4, 5, 10}},
			want: 16,
		},
		{
			name: "example3",
			args: args{nums: []int{2, 3, 4, 6, 8, 12}},
			want: 40,
		},
		{
			name: "no tuples",
			args: args{nums: []int{1, 2, 3}},
			want: 0,
		},
		{
			name: "multiple tuples",
			args: args{nums: []int{1, 2, 3, 4, 6, 12}},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tupleSameProduct(tt.args.nums); got != tt.want {
				t.Errorf("tupleSameProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
