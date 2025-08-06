package leetcode

import "testing"

func Test_numOfUnplacedFruits(t *testing.T) {
	type args struct {
		fruits  []int
		baskets []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				fruits:  []int{4, 2, 5},
				baskets: []int{3, 5, 4},
			},
			want: 1,
		},
		// {
		// 	name: "Example 2",
		// 	args: args{
		// 		fruits:  []int{3, 6, 1},
		// 		baskets: []int{6, 4, 7},
		// 	},
		// 	want: 0,
		// },
		// {
		// 	name: "All unplaced",
		// 	args: args{
		// 		fruits:  []int{10, 20, 30},
		// 		baskets: []int{1, 2, 3},
		// 	},
		// 	want: 3,
		// },
		// {
		// 	name: "All placed, same size",
		// 	args: args{
		// 		fruits:  []int{1, 2, 3},
		// 		baskets: []int{1, 2, 3},
		// 	},
		// 	want: 0,
		// },
		// {
		// 	name: "Edge case, n=1, placed",
		// 	args: args{
		// 		fruits:  []int{5},
		// 		baskets: []int{5},
		// 	},
		// 	want: 0,
		// },
		// {
		// 	name: "Edge case, n=1, unplaced",
		// 	args: args{
		// 		fruits:  []int{6},
		// 		baskets: []int{5},
		// 	},
		// 	want: 1,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfUnplacedFruits(tt.args.fruits, tt.args.baskets); got != tt.want {
				t.Errorf("numOfUnplacedFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
