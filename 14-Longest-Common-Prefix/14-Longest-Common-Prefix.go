package leetcode

import "strings"

type trieNode struct {
	children [26]*trieNode
	isEnd    bool
}

type trie struct {
	root *trieNode
}

func initTrie() *trie {
	return &trie{
		root: &trieNode{},
	}
}

func (t *trie) insert(w string) {
	temp := t.root

	for _, char := range w {
		point := char - 'a'
		if temp.children[point] == nil {
			temp.children[point] = &trieNode{}
		}
		temp = temp.children[point]
	}

	temp.isEnd = true
}

func longestCommonPrefix(strs []string) string {
	charTrie := initTrie()

	for _, v := range strs {
		charTrie.insert(v)
	}

	var ans strings.Builder
	temp := charTrie.root
  for temp != nil && !temp.isEnd {
		childCount := 0
		nextIdx := 0
		for i, child := range temp.children {
			if child != nil {
				childCount++
				if childCount >= 2 {
					return ans.String()
				}

				nextIdx = i
			}
		}

		ans.WriteByte(byte(nextIdx) + 'a')
		temp = temp.children[nextIdx]
	}

	return ans.String()
}
