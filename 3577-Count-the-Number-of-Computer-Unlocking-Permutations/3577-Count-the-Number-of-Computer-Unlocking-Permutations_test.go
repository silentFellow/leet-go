package leetcode

import "testing"

func Test_countPermutations(t *testing.T) {
	type args struct {
		complexity []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"single element", args{[]int{5}}, 1},
		{"all same elements", args{[]int{3, 3, 3}}, 0},
		{"min appears twice", args{[]int{2, 2, 3}}, 0},
		{"min not unique", args{[]int{1, 2, 1}}, 0},
		{"valid permutation", args{[]int{2, 3, 4}}, 2},
		{"min not first", args{[]int{3, 2, 4}}, 0},
		{"increasing sequence", args{[]int{1, 2, 3, 4}}, 6},
		{"decreasing sequence", args{[]int{4, 3, 2, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPermutations(tt.args.complexity); got != tt.want {
				t.Errorf("countPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
