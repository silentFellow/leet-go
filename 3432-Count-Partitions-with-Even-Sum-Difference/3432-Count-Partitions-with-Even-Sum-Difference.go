package leetcode

func countPartitions(nums []int) int {
	fullSum := 0
	for _, num := range nums {
		fullSum += num
	}

	ans := 0

	leftSum := 0
	for i:=0; i<len(nums)-1; i++ {
		num := nums[i]
		leftSum += num
		fullSum -= num

		if (fullSum-leftSum)%2 == 0 {
			ans++
		}
	}

	return ans
}
