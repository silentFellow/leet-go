package leetcode

import "testing"

func Test_divideArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{nums: []int{3, 2, 3, 2, 2, 2}},
			want: true,
		},
		{
			name: "example2",
			args: args{nums: []int{1, 2, 3, 4}},
			want: false,
		},
		{
			name: "all_pairs",
			args: args{nums: []int{1, 1, 2, 2, 3, 3, 4, 4}},
			want: true,
		},
		{
			name: "no_pairs",
			args: args{nums: []int{1, 2, 3, 4, 5, 6}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideArray(tt.args.nums); got != tt.want {
				t.Errorf("divideArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
