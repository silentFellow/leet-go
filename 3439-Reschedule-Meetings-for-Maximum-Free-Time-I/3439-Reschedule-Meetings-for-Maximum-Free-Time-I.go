package leetcode

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	freeTime := make([]int, 0)

	freeTime = append(freeTime, startTime[0]-0)
	for i := 1; i < len(startTime); i++ {
		freeTime = append(freeTime, startTime[i]-endTime[i-1])
	}
	freeTime = append(freeTime, eventTime-endTime[len(endTime)-1])

	windowSum := 0
	for i := range k + 1 {
		windowSum += freeTime[i]
	}
	ans := windowSum

	for i := k + 1; i < len(freeTime); i++ {
		windowSum = windowSum - freeTime[i-k-1] + freeTime[i]
		ans = max(ans, windowSum)
	}

	return ans
}
