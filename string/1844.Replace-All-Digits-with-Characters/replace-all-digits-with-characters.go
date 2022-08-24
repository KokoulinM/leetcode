package main

import "fmt"

func replaceDigits(s string) string {
	result := []byte(s)

	for i, v := range result {
		if v < 58 {
			result[i] = result[i-1] + (result[i] - 48)
		}
	}

	fmt.Println(string(result))

	return string(result)
}
