package leetcode

func getHappyString(n int, k int) string {
  arr := []string{"a", "b", "c"}

  for len(arr[len(arr)-1]) != n {
    cur := arr
    for _, val := range cur {
      if val[len(val)-1] != 'a' {
        arr = append(arr, val + "a")
      }

      if val[len(val)-1] != 'b' {
        arr = append(arr, val + "b")
      }

      if val[len(val)-1] != 'c' {
        arr = append(arr, val + "c")
      }
    }

    arr = arr[len(cur):]
  }

  if len(arr) < k {
    return ""
  }

  return arr[k-1]
}
