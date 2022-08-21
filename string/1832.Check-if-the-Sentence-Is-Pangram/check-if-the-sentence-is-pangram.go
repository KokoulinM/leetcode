package main

import "strings"

func checkIfPangram(sentence string) bool {
	alp := "thequickbrownfoxjumpsoverthelazydog"

	for _, i := range alp {
		isContein := strings.Contains(sentence, string(i))
		if !isContein {
			return false
		}
	}

	return true
}
