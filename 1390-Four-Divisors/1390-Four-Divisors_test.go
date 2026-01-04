package leetcode

import "testing"

func Test_sumFourDivisors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"Example 1", []int{21, 4, 7}, 32},     // Only 21 has 4 divisors: 1, 3, 7, 21
		{"Example 2", []int{21, 21}, 64},       // Both 21s, each sum is 32
		{"Example 3", []int{1, 2, 3, 4, 5}, 0}, // No number with exactly 4 divisors
		{"Duplicates", []int{10, 10, 10}, 54},    // 10 appears 3 times, sum is 18*3=54
		{"Large input", []int{99991, 99999}, 0},  // Large numbers, likely no four divisors
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumFourDivisors(tt.nums); got != tt.want {
				t.Errorf("sumFourDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
