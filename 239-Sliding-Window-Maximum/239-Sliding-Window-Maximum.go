package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	dequeue := make([]int, 0)
	ans := make([]int, 0)

	for i, val := range nums {
		// remove dequeue index if less than the window
		for len(dequeue) > 0 && dequeue[0] < i-k+1 {
			dequeue = dequeue[1:]
		}

		// remove smaller elements after addition values in dequeue
		for len(dequeue) > 0 && nums[dequeue[len(dequeue)-1]] <= val {
			dequeue = dequeue[:len(dequeue)-1]
		}

		// append current index
		dequeue = append(dequeue, i)

		// append the value to result
		if i >= k-1 {
			ans = append(ans, nums[dequeue[0]])
		}
	}

	return ans
}
