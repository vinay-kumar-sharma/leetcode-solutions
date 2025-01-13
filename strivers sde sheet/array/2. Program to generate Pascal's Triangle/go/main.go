package main

import "fmt"

func main() {

	n := 5

	fmt.Println(generate(n))

}
func generate(numRows int) [][]int {

	triangle := make([][]int, numRows)

	triangle[0] = make([]int, 1)
	triangle[0][0] = 1
	for i := 1; i < numRows; i++ {

		triangle[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				triangle[i][j] = 1
			} else {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
			}
		}
	}

	return triangle
}
