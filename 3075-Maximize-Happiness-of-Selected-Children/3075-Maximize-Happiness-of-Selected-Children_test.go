package leetcode

import "testing"

func Test_maximumHappinessSum(t *testing.T) {
	type args struct {
		happiness []int
		k         int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// Basic cases
		{name: "example1", args: args{happiness: []int{1, 2, 3}, k: 2}, want: 4},
		{name: "example2", args: args{happiness: []int{1, 1, 1, 1}, k: 2}, want: 1},
		{name: "example3", args: args{happiness: []int{2, 3, 4, 5}, k: 1}, want: 5},

		// Edge cases
		{name: "single_element", args: args{happiness: []int{10}, k: 1}, want: 10},
		{name: "all_zero_after_decrement", args: args{happiness: []int{1, 1, 1}, k: 3}, want: 1},
		{name: "k_equals_n", args: args{happiness: []int{5, 4, 3, 2, 1}, k: 5}, want: 9},

		// Decrement to zero
		{name: "decrement_to_zero", args: args{happiness: []int{3, 2, 1}, k: 3}, want: 4},

		// Random order
		{name: "random_order", args: args{happiness: []int{7, 1, 5, 3, 6, 4}, k: 3}, want: 15},

		// All same values
		{name: "all_same", args: args{happiness: []int{5, 5, 5, 5, 5}, k: 3}, want: 12},

		// Decrement skips negatives
		{name: "decrement_skips_negatives", args: args{happiness: []int{2, 2, 2}, k: 3}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumHappinessSum(tt.args.happiness, tt.args.k); got != tt.want {
				t.Errorf("maximumHappinessSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
