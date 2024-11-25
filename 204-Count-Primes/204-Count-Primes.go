package leetcode

func countPrimes(n int) int {
  if n == 0 {return 0}

  flags := make([]bool, n)

  for i := range flags {
    if i == 0 || i == 1 {continue} // since 0 and 1 are non prime
    flags[i] = true
  }

  for i:=1; i<n; i++ {
    if !flags[i] {continue}
    for j:=i*i; j<n; j=j+i {
      flags[j] = false
    }
  }

  count := 0
  for _,val := range flags {
    if val {
      count++
    }
  }

  return count
}
