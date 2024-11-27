package leetcode

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "example2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
		{
			name: "no_occurrences",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				val:  6,
			},
			want: 5,
		},
		{
			name: "all_occurrences",
			args: args{
				nums: []int{2, 2, 2, 2, 2},
				val:  2,
			},
			want: 0,
		},
		{
			name: "mixed_occurrences",
			args: args{
				nums: []int{4, 1, 2, 3, 4, 4, 5},
				val:  4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
