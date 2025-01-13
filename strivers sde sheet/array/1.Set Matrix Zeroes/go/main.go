package main

import "fmt"

func main() {

	array01 := [][]int{{0, 1}}

	fmt.Println("array=>", fmt.Sprintf("%v", setZeroes(array01)))

	array := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}

	fmt.Println("array=>", fmt.Sprintf("%v", setZeroes(array)))

	array1 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}

	fmt.Println("array1=>", fmt.Sprintf("%v", setZeroes(array1)))
}

func setZeroes(matrix [][]int) [][]int {
	column0 := 1
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0

				if j != 0 {
					matrix[0][j] = 0
				} else {
					column0 = 0
				}
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if column0 == 0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

	return matrix
}
