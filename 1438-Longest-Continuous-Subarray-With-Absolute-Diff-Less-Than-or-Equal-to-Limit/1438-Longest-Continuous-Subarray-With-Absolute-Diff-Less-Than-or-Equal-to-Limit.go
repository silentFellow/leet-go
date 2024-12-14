package leetcode

// INTITUTION: limit should not exceed windows max and windows min
// since, all other elements should be within limit
// if abs(max - min) < limit

func longestSubarray(nums []int, limit int) int {
  l := 0
  ans := 0

  minDeque := []int{}
  maxDeque := []int{}

  for r, val := range nums {
    // remove all small elements when new val is added to maxDeque
    for len(maxDeque) > 0 && nums[maxDeque[len(maxDeque)-1]] <= val {
      maxDeque = maxDeque[:len(maxDeque)-1]
    }

    // remove all large elements when new val is added to maxDeque
    for len(minDeque) > 0 && nums[minDeque[len(minDeque)-1]] >= val {
      minDeque = minDeque[:len(minDeque)-1]
    }

    // append the index to both min and max dequeue
    minDeque = append(minDeque, r)
    maxDeque = append(maxDeque, r)

    // check absolute difference if exceeds the limit
    for nums[maxDeque[0]] - nums[minDeque[0]] > limit {
      l++

      // remove all index from both dequeues if they are less than l
      for maxDeque[0] < l {
        maxDeque = maxDeque[1:]
      }
      for minDeque[0] < l {
        minDeque = minDeque[1:]
      }
    }

    ans = max(ans, r-l+1)
  }

  return ans
}
