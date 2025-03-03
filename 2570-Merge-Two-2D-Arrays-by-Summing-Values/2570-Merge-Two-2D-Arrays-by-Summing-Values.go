package leetcode

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	maxv := 0
	for _, v := range nums1 {
		maxv = max(maxv, v[0])
	}
	for _, v := range nums2 {
		maxv = max(maxv, v[0])
	}

	sum := make([]int, maxv)
  for _, v := range nums1 {
    sum[v[0]-1] += v[1]
  }
  for _, v := range nums2 {
    sum[v[0]-1] += v[1]
  }

  ans := make([][]int, 0)
  for i, v := range sum {
    if v != 0 {
      ans = append(ans, []int{i+1, v})
    }
  }

  return ans
}
