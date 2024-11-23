package leetcode

import "testing"

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		n        int
		bad      int
		expected int
	}{
		{5, 4, 4},
		{1, 1, 1},
	}

	for _, test := range tests {
		badVersion = test.bad
		result := firstBadVersion(test.n)
		if result != test.expected {
			t.Errorf("firstBadVersion(%d) = %d; want %d", test.n, result, test.expected)
		}
	}
}
