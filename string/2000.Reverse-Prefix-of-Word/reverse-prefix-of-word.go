package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	b := []byte(s)

	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}

	return string(b)
}

func reversePrefix(word string, ch byte) string {
	i := strings.Index(word, string(ch))

	fmt.Println(reverseString(word[:i+1]))

	word = reverseString(word[:i+1]) + word[i:]

	return word
}
