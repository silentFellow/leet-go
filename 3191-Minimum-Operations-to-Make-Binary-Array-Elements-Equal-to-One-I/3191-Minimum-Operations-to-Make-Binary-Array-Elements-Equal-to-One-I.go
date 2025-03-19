package leetcode

func minOperations(nums []int) int {
	if sum(nums) == len(nums) {
		return 0
	}

	ans := 0
	for i := range len(nums) - 2 {
		if nums[i] == 0 {
			for j := i; j < i+3; j++ {
				nums[j] = change(nums[j])
			}

			ans++
		}
	}

	if sum(nums) == len(nums) {
		return ans
	} else {
		return -1
	}
}

func sum(arr []int) int {
	ans := 0
	for _, v := range arr {
		ans += v
	}

	return ans
}

func change(v int) int {
	if v == 1 {
		return 0
	} else {
		return 1
	}
}
