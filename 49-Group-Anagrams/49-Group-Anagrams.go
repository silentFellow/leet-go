package leetcode

func groupAnagrams(strs []string) [][]string {
  hmap := make(map[[26]int][]string)

  for _, val := range strs {
    arr := getAnagramArray(val)
    hmap[arr] = append(hmap[arr], val)
  }

  ans := make([][]string, 0, len(hmap))
  for _, val := range hmap {
    ans = append(ans, val)
  }

  return ans
}

func getAnagramArray(str string) [26]int {
  arr := [26]int{}

  for _, val := range str {
    arr[val - 'a']++
  }

  return arr
}
