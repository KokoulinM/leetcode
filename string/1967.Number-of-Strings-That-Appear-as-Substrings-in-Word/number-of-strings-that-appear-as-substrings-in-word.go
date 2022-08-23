package main

import "strings"

func numOfStrings(patterns []string, word string) int {
	var output int

	for _, v := range patterns {
		if strings.Contains(word, string(v)) {
			output++
		}
	}

	return output
}
