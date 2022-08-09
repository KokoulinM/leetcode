package main

func matrixHorizontal(mat [][]int) []int {
	var result []int

	for i := 0; i < len(mat); i++ {
		if i % 2 == 0 {
			for j := 0; j < len(mat); j++ {
				result = append(result, mat[i][j])
			}
		} else {
			for j := len(mat[i]) - 1; j >= 0; j-- {
				result = append(result, mat[i][j])
			}
		}
	}

	return result
}