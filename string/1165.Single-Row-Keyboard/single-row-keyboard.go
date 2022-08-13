package main

import (
	"strings"
)

func abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func calculateTime(keyboard string, word string) int {
	var result int
	var prefIndex int

	for _, v := range word {
		index := strings.Index(keyboard, string(v))

		result += abs(index, prefIndex)
		prefIndex = index
	}

	return result
}
