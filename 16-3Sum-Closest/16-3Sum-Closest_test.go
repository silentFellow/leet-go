package leetcode

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		{
			name: "example2",
			args: args{
				nums:   []int{0, 0, 0},
				target: 1,
			},
			want: 0,
		},
		{
			name: "example3",
			args: args{
				nums:   []int{1, 1, 1, 1},
				target: 3,
			},
			want: 3,
		},
		{
			name: "example4",
			args: args{
				nums:   []int{-1, 0, 1, 1},
				target: 2,
			},
			want: 2,
		},
		{
			name: "example5",
			args: args{
				nums:   []int{1, 2, 3, 4},
				target: 6,
			},
			want: 6,
		},
		{
			name: "example6",
			args: args{
				nums:   []int{2, 3, 8, 9, 10},
				target: 16,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
