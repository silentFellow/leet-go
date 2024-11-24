package leetcode

func minSum(nums1 []int, nums2 []int) int64 {
	var s1, z1, s2, z2 int

	for _, val := range nums1 {
		if val == 0 {
			z1++
		}
		s1 += val
	}

	for _, val := range nums2 {
		if val == 0 {
			z2++
		}
		s2 += val
	}

	if (z1 == 0 && s1 < (s2+z2)) || (z2 == 0 && s2 < (s1+z1)) {
		return -1
	}

	if (s1 + z1) > (s2 + z2) {
		return int64(s1 + z1)
	} else {
		return int64(s2 + z2)
	}
}
