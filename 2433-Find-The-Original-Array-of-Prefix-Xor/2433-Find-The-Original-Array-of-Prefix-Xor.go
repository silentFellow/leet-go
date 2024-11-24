package leetcode

func findArray(pref []int) []int {
  xor := 0

  for i, val := range pref {
    pref[i] = xor ^ val
    xor ^= pref[i]
  }

  return pref
}
