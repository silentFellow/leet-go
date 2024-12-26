package leetcode

// use bottom-up dp to more space optimized solution
func findTargetSumWays(nums []int, target int) int {
	// contains sum with how much ways it can be get
	dp := make(map[int]int)

	// 1 way to get 0 as sum
	dp[0] = 1

	for i := range len(nums) {
		next_dp := make(map[int]int)
		for key, val := range dp {
			next_dp[key+nums[i]] += val
			next_dp[key-nums[i]] += val
		}
		dp = next_dp
	}

	return dp[target]
}

// use top-down dp approach to optimize
// func findTargetSumWays(nums []int, target int) int {
//   dp := make(map[string]int)
//
// 	var backtrack func(idx, sum int) int
// 	backtrack = func(idx, sum int) int {
//     key := fmt.Sprintf("%d, %d", idx, sum)
//
//     if val, ok := dp[key]; ok {
//       return val
//     }
//
// 		if idx == len(nums) {
// 			if sum == target {
// 				return 1
// 			} else {
// 				return 0
// 			}
// 		}
//
// 		dp[key] = backtrack(idx+1, sum+nums[idx]) + backtrack(idx+1, sum-nums[idx])
//     return dp[key]
// 	}
//
// 	return backtrack(0, 0)
// }

// brute force not memory optimized
// func findTargetSumWays(nums []int, target int) int {
// 	var backtrack func(idx, sum int) int
// 	backtrack = func(idx, sum int) int {
// 		if idx == len(nums) {
// 			if sum == target {
// 				return 1
// 			} else {
// 				return 0
// 			}
// 		}
//
// 		return backtrack(idx+1, sum+nums[idx]) + backtrack(idx+1, sum-nums[idx])
// 	}
//
// 	return backtrack(0, 0)
// }
