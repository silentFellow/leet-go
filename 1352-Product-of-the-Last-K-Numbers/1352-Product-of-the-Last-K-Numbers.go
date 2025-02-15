package leetcode

// TODO: can be optimized through prefix products

type ProductOfNumbers struct {
	nums []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{}
}

func (this *ProductOfNumbers) Add(num int) {
	this.nums = append(this.nums, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	v := 1
	n := len(this.nums) - 1

	for i := range k {
		v *= this.nums[n-i]
	}

	return v
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
