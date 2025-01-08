package leetcode

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor()

	// Test case 1: Insert and search "apple"
	trie.Insert("apple")
	if got := trie.Search("apple"); !got {
		t.Errorf("Trie.Search() = %v, want %v", got, true)
	}

	// Test case 2: Search for "app" which is not inserted yet
	if got := trie.Search("app"); got {
		t.Errorf("Trie.Search() = %v, want %v", got, false)
	}

	// Test case 3: Check if "app" is a prefix
	if got := trie.StartsWith("app"); !got {
		t.Errorf("Trie.StartsWith() = %v, want %v", got, true)
	}

	// Test case 4: Insert "app" and search for it
	trie.Insert("app")
	if got := trie.Search("app"); !got {
		t.Errorf("Trie.Search() = %v, want %v", got, true)
	}
}
