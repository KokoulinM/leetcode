package main

import (
	"fmt"
)

func cellsInRange(s string) []string {
	var result []string
	startCell := s[0]
	startRows := s[1]

	endCell := s[3]
	endRows := s[4]

	for i := startCell; i <= endCell; i++ {
		for j := startRows; j <= endRows; j++ {
			result = append(result, fmt.Sprintf("%s%s", string(i), string(j)))
		}
	}

	return result
}
