package main

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	type tire map[rune]tire

	root := tire{}

	for _, str := range dictionary {
		cur := root
		for _, s := range str {
			if cur[s] == nil {
				cur[s] = tire{}
			}
			cur = cur[s]
		}
		cur['#'] = tire{}
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		cur := root
		for j, w := range word {
			if cur['#'] != nil {
				words[i] = word[:j]
				break
			}
			if cur[w] == nil {
				break
			}
			cur = cur[w]
		}
	}

	return strings.Join(words, " ")
}
