package leetcode

// O(n)
func containsDuplicate(nums []int) bool {
	hmap := make(map[int]int)

	for _, val := range nums {
		hmap[val] = hmap[val] + 1
	}

	for _, val := range hmap {
		if val >= 2 {
			return true
		}
	}

	return false
}

// O(nlogn)

// import "sort"

// func containsDuplicate(nums []int) bool {
//   sort.Ints(nums)
//
//   for i := range len(nums)-1 {
//     if nums[i] == nums[i+1] {
//       return true
//     }
//   }
//
//   return false
// }
