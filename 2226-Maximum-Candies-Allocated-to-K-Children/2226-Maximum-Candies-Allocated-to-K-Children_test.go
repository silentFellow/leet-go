package leetcode

import "testing"

func Test_maximumCandies(t *testing.T) {
	type args struct {
		candies []int
		k       int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				candies: []int{5, 8, 6},
				k:       3,
			},
			want: 5,
		},
		{
			name: "example2",
			args: args{
				candies: []int{2, 5},
				k:       11,
			},
			want: 0,
		},
		{
			name: "example1",
			args: args{
				candies: []int{5, 8, 6},
				k:       3,
			},
			want: 5,
		},
		{
			name: "example2",
			args: args{
				candies: []int{2, 5},
				k:       11,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCandies(tt.args.candies, tt.args.k); got != tt.want {
				t.Errorf("maximumCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMax(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMax(tt.args.arr); got != tt.want {
				t.Errorf("findMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
