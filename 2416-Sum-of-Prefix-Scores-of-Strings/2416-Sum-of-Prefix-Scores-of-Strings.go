package leetcode

type trieNode struct {
	children [26]*trieNode
	count    int
}

type trie struct {
	root *trieNode
}

func getTrie() trie {
	return trie{root: &trieNode{}}
}

func(t trie) insert(word string) {
  root := t.root
  for _, char := range word {
    idx := char - 'a'
    if root.children[idx] == nil {
      root.children[idx] = &trieNode{count: 0}
    }

    root.children[idx].count++
    root = root.children[idx]
  }
}

func(t trie) getScore(word string) int {
  count := 0
  root := t.root

  for _, char := range word {
    idx := char - 'a'
    if root.children[idx] == nil {
      return count
    }

    count += root.children[idx].count
    root = root.children[idx]
  }

  return count
}

func sumPrefixScores(words []string) []int {
  t := getTrie()

  for _, word := range words {
    t.insert(word)
  }

  ans := []int{}
  for _, word := range words {
    score := t.getScore(word)
    ans = append(ans, score)
  }

  return ans
}
