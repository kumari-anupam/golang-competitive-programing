package main

import "fmt"

func main() {
	matrix := [][]int{{1,1,1},{1,0,1},{1,1,0}}
	setZeroes(matrix)

	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	col0 := 1
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i ++ {
		for j := 0; j < m ; j ++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				if j != 0 {
					matrix[0][j] = 0
				} else {
					col0 = 0
				}
			}
		}
	}

	// travese matrix and put zero based on first row and column

	for i := 1; i < n; i ++ {
		for j := 1; j < m; j ++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// set first row to zero
	if matrix[0][0] == 0 {
		for j := 0; j < m; j ++ {
			matrix[0][j] = 0
		}
	}

	// set first column to zero
	if col0 == 0 {
		for i := 0; i < n; i ++ {
			matrix[i][0] = 0
		}
	}
}

