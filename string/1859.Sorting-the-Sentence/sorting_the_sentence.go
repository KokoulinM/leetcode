package main

import (
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	tArr := strings.Split(s, " ")

	result := make([]string, len(tArr))

	for i := 0; i < len(tArr); i++ {
		pos := tArr[i][len(tArr[i])-1]

		p, _ := strconv.Atoi(string(pos))

		result[p-1] = tArr[i][:len(tArr[i])-1]
	}

	return strings.Join(result, " ")
}
