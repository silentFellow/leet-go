package leetcode

import "testing"

func TestFindSumPairsScenario(t *testing.T) {
	// Scenario based on the example in the problem description
	obj := Constructor([]int{1, 1, 2, 2, 2, 3}, []int{1, 4, 5, 2, 5, 4})

	if got := obj.Count(7); got != 8 {
		t.Errorf("Count(7) = %v, want %v", got, 8)
	}

	obj.Add(3, 2) // nums2 = [1,4,5,4,5,4]
	if got := obj.Count(8); got != 2 {
		t.Errorf("Count(8) = %v, want %v", got, 2)
	}

	if got := obj.Count(4); got != 1 {
		t.Errorf("Count(4) = %v, want %v", got, 1)
	}

	obj.Add(0, 1) // nums2 = [2,4,5,4,5,4]
	obj.Add(1, 1) // nums2 = [2,5,5,4,5,4]
	if got := obj.Count(7); got != 11 {
		t.Errorf("Count(7) = %v, want %v", got, 11)
	}
}

func TestFindSumPairsEdgeCases(t *testing.T) {
	// Edge case: nums1 and nums2 of length 1
	obj := Constructor([]int{1}, []int{2})
	if got := obj.Count(3); got != 1 {
		t.Errorf("Count(3) = %v, want %v", got, 1)
	}
	obj.Add(0, 1)
	if got := obj.Count(3); got != 0 {
		t.Errorf("Count(3) after add = %v, want %v", got, 0)
	}
	if got := obj.Count(4); got != 1 {
		t.Errorf("Count(4) after add = %v, want %v", got, 1)
	}
}

func TestFindSumPairsNoPairs(t *testing.T) {
	obj := Constructor([]int{10, 20}, []int{30, 40})
	if got := obj.Count(100); got != 0 {
		t.Errorf("Count(100) = %v, want %v", got, 0)
	}
}

func TestFindSumPairsMultipleAdds(t *testing.T) {
	obj := Constructor([]int{2, 3}, []int{1, 2})
	obj.Add(0, 3) // nums2 = [4,2]
	obj.Add(1, 2) // nums2 = [4,4]
	if got := obj.Count(6); got != 2 {
		t.Errorf("Count(6) = %v, want %v", got, 2)
	}
}
