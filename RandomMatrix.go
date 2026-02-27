package main

import (
	"math/rand"
)

func RandomMatrix(rows, cols, minVal, maxVal int) [][]int {
	matrix := make([][]int, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)

		for j := 0; j < cols; j++ {
			matrix[i][j] = rand.Intn(maxVal-minVal+1) + minVal
		}
	}
	return matrix
}
