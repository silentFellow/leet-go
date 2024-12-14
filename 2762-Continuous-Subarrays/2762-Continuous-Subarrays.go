package leetcode

func continuousSubarrays(nums []int) int64 {
	ans := 0
	l := 0

	// dequeues to keep track of min and max elements indices
	maxDequeue := make([]int, 0)
	minDequeue := make([]int, 0)

	for r, val := range nums {
		// pop all small elements if current val is added
		for len(maxDequeue) > 0 && nums[maxDequeue[len(maxDequeue)-1]] <= val {
			maxDequeue = maxDequeue[:len(maxDequeue)-1]
		}

		// pop all large elements if current val is added
		for len(minDequeue) > 0 && nums[minDequeue[len(minDequeue)-1]] >= val {
			minDequeue = minDequeue[:len(minDequeue)-1]
		}

		// append both index
		minDequeue = append(minDequeue, r)
		maxDequeue = append(maxDequeue, r)

		// check and pop all elements when max-min difference is greater than 2
		for nums[maxDequeue[0]]-nums[minDequeue[0]] > 2 {
			l++

			// pop all elements in queue if less than l
			if l > maxDequeue[0] {
				maxDequeue = maxDequeue[1:]
			}
			if l > minDequeue[0] {
				minDequeue = minDequeue[1:]
			}
		}

		ans += r - l + 1
	}

	return int64(ans)
}
