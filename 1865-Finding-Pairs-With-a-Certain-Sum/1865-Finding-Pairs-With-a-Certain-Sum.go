package leetcode

// NOTE: arr1 does not need to be stored in the FindSumPairs struct,
// since arr1 will not change after initialization.
// We can simply calculate its frequency map during initialization.

type FindSumPairs struct {
	arr1, arr2   []int
	hmap1, hmap2 map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	hmap1, hmap2 := make(map[int]int), make(map[int]int)
	for _, v := range nums1 {
		hmap1[v]++
	}

	for _, v := range nums2 {
		hmap2[v]++
	}
	return FindSumPairs{
		arr1:  nums1,
		arr2:  nums2,
		hmap1: hmap1,
		hmap2: hmap2,
	}
}

func (this *FindSumPairs) Add(index int, val int) {
	v := this.arr2[index]
	change := v + val

	this.arr2[index] = change

	this.hmap2[v]--
	if old, _ := this.hmap2[v]; old == 0 {
		delete(this.hmap2, v)
	}
	this.hmap2[change]++
}

func (this *FindSumPairs) Count(tot int) int {
	cnt := 0
	for k1, v1 := range this.hmap1 {
		if v2, ok := this.hmap2[tot-k1]; ok {
			cnt += (v1 * v2)
		}
	}

	return cnt
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
