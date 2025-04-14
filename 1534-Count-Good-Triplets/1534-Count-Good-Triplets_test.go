package leetcode

import "testing"

func Test_countGoodTriplets(t *testing.T) {
	type args struct {
		arr []int
		a   int
		b   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				arr: []int{3, 0, 1, 1, 9, 7},
				a:   7,
				b:   2,
				c:   3,
			},
			want: 4,
		},
		{
			name: "example2",
			args: args{
				arr: []int{1, 1, 2, 2, 3},
				a:   0,
				b:   0,
				c:   1,
			},
			want: 0,
		},
		{
			name: "all_elements_same",
			args: args{
				arr: []int{1, 1, 1, 1},
				a:   0,
				b:   0,
				c:   0,
			},
			want: 4,
		},
		{
			name: "no_good_triplets",
			args: args{
				arr: []int{10, 20, 30, 40},
				a:   5,
				b:   5,
				c:   5,
			},
			want: 0,
		},
		{
			name: "large_input",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				a:   10,
				b:   10,
				c:   10,
			},
			want: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodTriplets(tt.args.arr, tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("countGoodTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
