package main

import "fmt"

func main() {

	fmt.Println("))())(", "010100", "=>", fmt.Sprintf("%v", canBeValid("))())(", "010100")))
	fmt.Println("()()", "0000", "=>", fmt.Sprintf("%v", canBeValid("()()", "0000")))
	fmt.Println(")", "0", "=>", fmt.Sprintf("%v", canBeValid(")", "0")))
	fmt.Println("(()))", "11011", "=>", fmt.Sprintf("%v", canBeValid("(()))", "11011")))

	fmt.Println("(", "0", "=>", fmt.Sprintf("%v", canBeValid("(", "0")))
	fmt.Println("", "", "=>", fmt.Sprintf("%v", canBeValid("", "")))
	fmt.Println("(((((", "11111", "=>", fmt.Sprintf("%v", canBeValid("(((((", "11111")))
	fmt.Println("()()()", "000000", "=>", fmt.Sprintf("%v", canBeValid("()()()", "000000")))
	fmt.Println("))(((", "00000", "=>", fmt.Sprintf("%v", canBeValid("))(((", "00000")))
}
func canBeValid(s string, locked string) bool {
	if len(s)%2 != 0 {
		return false // Odd-length strings cannot be valid
	}

	minOpen, maxOpen := 0, 0 // Min and max possible open parentheses
	for i := 0; i < len(s); i++ {
		if locked[i] == '1' {
			if s[i] == '(' {
				minOpen++
				maxOpen++
			} else {
				minOpen--
				maxOpen--
			}
		} else {
			minOpen--
			maxOpen++
		}
		if maxOpen < 0 {
			return false // Too many ')'
		}
		minOpen = max(minOpen, 0) // Min open cannot go below 0
	}

	return minOpen == 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
