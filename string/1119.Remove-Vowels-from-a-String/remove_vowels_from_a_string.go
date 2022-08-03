package main

import (
	"strings"
)

func removeVowels(s string) string {
	except := map[string]string{"a": "a", "e": "e", "i": "i", "o": "o", "u": "u"}
	var result []string

	for _, str := range s {
		if string(str) != except[string(str)] {
			result = append(result, string(str))
		}
	}

	return strings.Join(result, "")
}
