package leetcode

func maxFrequencyElements(nums []int) int {
	hmap := make(map[int]int)
	for _, v := range nums {
		hmap[v]++
	}

	maxFreq := 0
	for _, freq := range hmap {
		maxFreq = max(maxFreq, freq)
	}

	maxCount := 0
	for _, freq := range hmap {
		if freq == maxFreq {
			maxCount++
		}
	}

	return maxCount * maxFreq
}
