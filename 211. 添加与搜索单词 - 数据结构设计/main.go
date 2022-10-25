package main

import "fmt"

type WordDictionary struct {
	words  [26]*WordDictionary
	isWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (t *WordDictionary) AddWord(word string) {
	node := t
	for _, w := range word {
		ch := w - 'a'
		if node.words[ch] == nil {

			node.words[ch] = &WordDictionary{}
		}
		node = node.words[ch]
	}

	node.isWord = true

}

func dfs(word string, node *WordDictionary, index int) bool {
	if len(word) == index {
		return node.isWord
	}

	s := word[index]
	chs := s - 'a'
	if s != '.' {
		ch := node.words[chs]
		if ch != nil && dfs(word, ch, index+1) {
			return true
		}
	} else {

		for i := range node.words {
			ch := node.words[i]
			if ch != nil && dfs(word, ch, index+1) {
				return true
			}
		}

	}

	return false
}

func (t *WordDictionary) Search(word string) bool {

	return dfs(word, t, 0)
}

func main() {
	word := "dad"
	obj := Constructor()
	obj.AddWord(word)
	// param_2 := obj.Search(".ad")
	fmt.Println(obj)
}
