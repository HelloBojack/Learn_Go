package main

import (
	"fmt"
)

type Trie struct {
	children map[rune]*Trie
	isEnd    bool
}

func Constructor() *Trie {
	return &Trie{children: make(map[rune]*Trie, 26)}
}

func (t *Trie) Insert(word string) {

	for _, w := range word {
		if t.children[w] == nil {
			t.children[w] = Constructor()
		}
		t = t.children[w]
	}

	t.isEnd = true
}

func (t *Trie) Search(word string) bool {
	for _, w := range word {
		if t.children[w] == nil {
			return false
		}
		t = t.children[w]
	}

	return t.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	for _, w := range prefix {
		if t.children[w] == nil {
			return false
		}
		t = t.children[w]
	}

	return true
}

func main() {
	word := "apple"
	word2 := "app"
	obj := Constructor()
	obj.Insert(word)
	fmt.Println(obj)
	param_2 := obj.Search(word)
	param_3 := obj.Search(word2)
	fmt.Println(param_2, param_3)
}

/**
* Your Trie object will be instantiated and called as such:
* obj := Constructor();
* obj.Insert(word);
* param_2 := obj.Search(word);
* param_3 := obj.StartsWith(prefix);
 */
