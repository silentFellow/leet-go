package leetcode

import "testing"

func TestMinTimeToVisitAllPoints(t *testing.T) {
	tests := []struct {
		name     string
		points   [][]int
		expected int
	}{
		{
			name:     "Example from problem",
			points:   [][]int{{1, 1}, {3, 4}, {-1, 0}},
			expected: 7,
		},
		{
			name:     "Horizontal movement",
			points:   [][]int{{0, 0}, {5, 0}},
			expected: 5,
		},
		{
			name:     "Vertical movement",
			points:   [][]int{{0, 0}, {0, 7}},
			expected: 7,
		},
		{
			name:     "Diagonal movement",
			points:   [][]int{{2, 2}, {5, 5}},
			expected: 3,
		},
		{
			name:     "All on straight diagonal",
			points:   [][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}},
			expected: 3,
		},
		{
			name:     "Zigzag",
			points:   [][]int{{0, 0}, {2, 3}, {4, 1}, {1, 1}},
			expected: 8,
		},
		{
			name:     "All points the same",
			points:   [][]int{{1, 1}, {1, 1}, {1, 1}},
			expected: 0,
		},
		{
			name:     "Negative coordinates",
			points:   [][]int{{-2, -3}, {-4, -1}, {-1, -5}},
			expected: 6,
		},
		{
			name:     "Large jump",
			points:   [][]int{{0, 0}, {1000, 1000}},
			expected: 1000,
		},
		{
			name:     "Single point",
			points:   [][]int{{5, 5}},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minTimeToVisitAllPoints(tt.points)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
