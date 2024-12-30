package leetcode

func mergeSortedList(nums1, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	} else if len(nums2) == 0 {
		return nums1
	}

	res := make([]int, len(nums1)+len(nums2))
	i, j, temp := 0, 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			res[temp] = nums1[i]
			i++
		} else {
			res[temp] = nums2[j]
			j++
		}

		temp++
	}

	if i == len(nums1) {
		for j := j; j < len(nums2); j, temp = j+1, temp+1 {
			res[temp] = nums2[j]
		}
	} else {
		for i := i; i < len(nums1); i, temp = i+1, temp+1 {
			res[temp] = nums1[i]
		}
	}

	return res
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := mergeSortedList(nums1, nums2)

	if len(arr)%2 == 0 {
		middle := len(arr) / 2
		return float64(arr[middle]+arr[middle-1]) / 2
	} else {
		return float64(arr[len(arr)/2])
	}
}
