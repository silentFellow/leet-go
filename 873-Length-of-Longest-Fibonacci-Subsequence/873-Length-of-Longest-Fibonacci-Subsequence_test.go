package leetcode

import "testing"

func Test_lenLongestFibSubseq(t *testing.T) {
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
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8}},
			want: 5,
		},
		{
			name: "example2",
			args: args{arr: []int{1, 3, 7, 11, 12, 14, 18}},
			want: 3,
		},
		{
			name: "no_fib_sequence",
			args: args{arr: []int{1, 4, 6, 10, 15}},
			want: 3,
		},
		{
			name: "single_fib_sequence",
			args: args{arr: []int{1, 2, 3}},
			want: 3,
		},
		{
			name: "multiple_fib_sequences",
			args: args{arr: []int{1, 2, 3, 5, 8, 13, 21}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lenLongestFibSubseq(tt.args.arr); got != tt.want {
				t.Errorf("lenLongestFibSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
