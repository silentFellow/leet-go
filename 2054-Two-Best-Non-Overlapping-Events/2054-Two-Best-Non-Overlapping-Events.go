package leetcode

import (
	"sort"
)

func maxTwoEvents(events [][]int) int {
	maxSum := 0 // answer

  // sort based on start time
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

  // precompute max values after that index
	maxValues := make([]int, len(events))
	maxVal := events[len(events)-1][2]
	for i := len(events) - 1; i >= 0; i-- {
		maxVal = max(maxVal, events[i][2])
		maxValues[i] = maxVal
	}

  // iterate over each start time
	for i, event := range events {
    // set the next start time threshold
		thresh := event[1] + 1

    // binary search to find next start time after current ends
		l, r := i, len(events)-1
		for l < r {
			m := (l + r) / 2

			if events[m][0] < thresh {
				l = m + 1
			} else {
				r = m
			}
		}

    // find max value after the found element using precomputed maxvalue array
		if events[l][0] >= thresh { // if next found
			maxSum = max(maxSum, maxValues[l]+event[2])
		} else {
      // important edge case: in question it is atmost two
      // so single time interval value can be also maximum
			maxSum = max(maxSum, event[2])
		}
	}

	return maxSum
}
