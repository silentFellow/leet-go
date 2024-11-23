package leetcode


// INFO: This is a mock implementation of isBadVersion, since leetcode hides this function.
var badVersion int

func isBadVersion(version int) bool {
	return version >= badVersion
}


func firstBadVersion(n int) int {
  l, r := 1, n

  for l <= r {
    m := (l+r)/2
    if isBadVersion(m) {
      r = m-1
    } else {
      l = m+1
    }
  }

  return l
}
