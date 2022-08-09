package main

func diagonalSum(mat [][]int) int {
    var sum int

	for i := 0; i < len(mat); i++ {		
		sum += mat[i][i]
		j := len(mat) - i - 1

		if i != j {
			sum += mat[i][j]
		}
	}

	return sum
}