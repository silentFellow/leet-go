package leetcode

func vowelStrings(words []string, queries [][]int) []int {
	isValid := make([]bool, len(words))
	for i, word := range words {
		vowels := map[byte]struct{}{
			'a': {},
			'e': {},
			'i': {},
			'o': {},
			'u': {},
		}

		_, ok := vowels[word[0]]
		_, ok2 := vowels[word[len(word)-1]]
		if ok && ok2 {
			isValid[i] = true
		} else {
			isValid[i] = false
		}
	}

  prefix := make([]int, len(words)+1)
  for i := range words {
    prefix[i+1] = prefix[i]
    if isValid[i] {
      prefix[i+1]++
    }
  }

  results := make([]int, len(queries))
  for i, v := range queries {
    l, r := v[0], v[1]
    results[i] = prefix[r+1] - prefix[l]
  }


	return results
}
