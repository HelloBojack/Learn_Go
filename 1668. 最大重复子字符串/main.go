package main

import "strings"

func maxRepeating(sequence string, word string) int {
	res := 0
	temp := word

	for len(temp) <= len(sequence) {
		if !strings.Contains(sequence, temp) {
			break
		} else {
			res++
			temp = temp + word
		}

	}

	return res
}
