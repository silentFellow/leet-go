package leetcode

import "testing"

func Test_findContentChildren(t *testing.T) {
	type args struct {
		g []int
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{g: []int{1, 2, 3}, s: []int{1, 1}},
			want: 1,
		},
		{
			name: "example2",
			args: args{g: []int{1, 2}, s: []int{1, 2, 3}},
			want: 2,
		},
		{
			name: "no cookies",
			args: args{g: []int{1, 2, 3}, s: []int{}},
			want: 0,
		},
		{
			name: "all children content",
			args: args{g: []int{1, 2, 3}, s: []int{3, 2, 1}},
			want: 3,
		},
		{
			name: "more cookies than children",
			args: args{g: []int{2, 3}, s: []int{1, 2, 3, 4}},
			want: 2,
		},
		{
			name: "large values",
			args: args{g: []int{100000}, s: []int{100000}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContentChildren(tt.args.g, tt.args.s); got != tt.want {
				t.Errorf("findContentChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}
