package leetcode

import "math"

func closestPrimes(left int, right int) []int {
  allPrimes := make([]bool, right+1)
  for i:=2; i <= right; i++ {
    allPrimes[i] = true
  }
  for i:=2; i*i<=right; i++ {
    if allPrimes[i] {
      for j:=i*i; j<=right; j+=i {
        allPrimes[j] = false
      }
    }
  }

  primesInBetween := []int{}
  for i:=left; i<=right; i++ {
    if allPrimes[i] {
      primesInBetween = append(primesInBetween, i)
    }
  }

  minDiff := math.MaxInt
  result := []int{-1, -1}

  for i:=1; i<len(primesInBetween); i++ {
    diff := primesInBetween[i] - primesInBetween[i-1]
    if diff < minDiff {
      minDiff = diff
      result[0], result[1] = primesInBetween[i-1], primesInBetween[i]
    }
  }

  return result
}
