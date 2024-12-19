package leetcode

import "testing"

func Test_maxChunksToSorted(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{arr: []int{4, 3, 2, 1, 0}},
			want: 1,
		},
		{
			name: "example2",
			args: args{arr: []int{1, 0, 2, 3, 4}},
			want: 4,
		},
		{
			name: "example3",
			args: args{arr: []int{1, 2, 0, 3}},
			want: 2,
		},
		{
			name: "sorted array",
			args: args{arr: []int{0, 1, 2, 3, 4}},
			want: 5,
		},
		{
			name: "reverse sorted array",
			args: args{arr: []int{4, 3, 2, 1, 0}},
			want: 1,
		},
		{
			name: "single element",
			args: args{arr: []int{0}},
			want: 1,
		},
		{
			name: "two elements",
			args: args{arr: []int{1, 0}},
			want: 1,
		},
		{
			name: "three elements",
			args: args{arr: []int{0, 2, 1}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunksToSorted(tt.args.arr); got != tt.want {
				t.Errorf("maxChunksToSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
