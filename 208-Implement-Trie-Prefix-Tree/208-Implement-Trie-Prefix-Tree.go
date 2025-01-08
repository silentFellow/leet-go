package leetcode

type TrieNode struct {
	Children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{},
	}
}

func (this *Trie) Insert(word string) {
	root := this.root
	for _, char := range word {
		idx := char - 'a'
		if root.Children[idx] == nil {
			root.Children[idx] = &TrieNode{}
		}

		root = root.Children[idx]
	}

	root.isEnd = true
}

func (this *Trie) Search(word string) bool {
	root := this.root
	for _, char := range word {
		idx := char - 'a'
		if root.Children[idx] == nil {
			return false
		}

		root = root.Children[idx]
	}

	return root.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	root := this.root
	for _, char := range prefix {
		idx := char - 'a'
		if root.Children[idx] == nil {
			return false
		}

    root = root.Children[idx]
	}

	return !root.isEnd
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
