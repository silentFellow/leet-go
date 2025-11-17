package leetcode

func kLengthApart(nums []int, k int) bool {
	fo := -1

	for i, v := range nums {
		if v == 1 {
			// (i-fo+1) gives the sequence
			// (i-fo+1-2) remove both one from the seqence
			if fo == -1 || (i-fo+1-2) >= k {
				fo = i
			} else {
				return false
			}
		}
	}

	return true
}
