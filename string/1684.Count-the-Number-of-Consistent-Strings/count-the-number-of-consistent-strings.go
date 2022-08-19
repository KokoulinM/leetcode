package main

import (
	"fmt"
	"sort"
	"strings"
)

// allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
// allowed = "abc", words = ["a","b","c","ab","ac","bc","abc"]

func countConsistentStrings(allowed string, words []string) int {
	var result int

	for _, w := range words {
		s := strings.Split(w, "")
		sort.Strings(s)

		s2 := strings.Split(allowed, "")
		sort.Strings(s2)

		fmt.Println(strings.Join(s2, ""), strings.Join(s, ""))

		if strings.Index(strings.Join(s2, ""), strings.Join(s, "")) == -1 {
			continue
		}
		result++
	}

	return result
}
