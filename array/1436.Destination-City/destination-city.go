package main

import "fmt"

// paths: [][]string{{"D", "B"}, {"C", "A"}, {"B", "C"}},

//{"D", "B"},
//{"C", "A"},
//{"B", "C"}

// Концовкой является маршрут с index == 1 и для котогоро нет соответствия с index == 0

func destCity(paths [][]string) string {
	var output string

	for i := 0; i < len(paths); i++ {
		output = paths[i][1]
		for j := 1; j < len(paths); j++ {
			if output == paths[j][0] {
				fmt.Println("continue")
				continue
			}
		}

		fmt.Println(output)
	}

	return output
}
