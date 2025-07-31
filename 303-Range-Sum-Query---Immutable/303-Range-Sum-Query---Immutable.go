package leetcode

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)

	prefixSum := make([]int, n+1)
	for i := range n {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	return NumArray{
		prefixSum: prefixSum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	prefixSum := this.prefixSum

	return prefixSum[right+1] - prefixSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
