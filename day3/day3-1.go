package day3

import (
	"fmt"
	"unicode"
)

// struct to hold number data (number, initial digit, final digit)
type number struct {
	number       int
	initialdigit []int
	finaldigit   []int
}

func getNumbers(lines []string, verbose bool) []number {
	var initialdigit []int
	var finaldigit []int
	var num int = 0
	// initialize number array
	numbers := make([]number, 0)
	initialdigit = []int{-1, -1}
	finaldigit = []int{-1, -1}
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			current := rune(lines[i][j])
			if unicode.IsDigit(current) {
				if initialdigit[0] == -1 {
					initialdigit = []int{i, j}
				}
				finaldigit = []int{i, j}
				num = num*10 + int(rune(lines[i][j])-'0')
			} else {
				if initialdigit[0] != -1 {
					numbers = append(numbers, number{num, initialdigit, finaldigit})
					num = 0
					initialdigit = []int{-1, -1}
					finaldigit = []int{-1, -1}
				}
			}
		}
		if initialdigit[0] != -1 {
			numbers = append(numbers, number{num, initialdigit, finaldigit})
			num = 0
			initialdigit = []int{-1, -1}
			finaldigit = []int{-1, -1}
		}
	}
	// print numbers if verbose
	if verbose {
		for _, num := range numbers {
			println(num.number, "From: (", num.initialdigit[0], ",", num.initialdigit[1], ") To: (", num.finaldigit[0], ",", num.finaldigit[1], ")")
		}
	}
	return numbers
}

// a rune is symbol if it isn't a digit or a dot
func isSym(r rune) bool {
	if !unicode.IsDigit(r) && r != '.' {
		fmt.Println("Symbol: ", string(r))
	}
	return !unicode.IsDigit(r) && r != '.'
}

// a number is a part number if it's adjacent to a symbol
func isPartNumber(lines []string, n number) bool {
	row := n.initialdigit[0]
	inicol := n.initialdigit[1]
	endcol := n.finaldigit[1]
	// check if the number is adjacent to a symbol
	if inicol > 0 {
		if isSym(rune(lines[row][inicol-1])) { // left
			return true
		}
		if row > 0 {
			if isSym(rune(lines[row-1][inicol-1])) { // up-left
				return true
			}
		}
		if row < len(lines)-1 {
			if isSym(rune(lines[row+1][inicol-1])) { // down-left
				return true
			}
		}
	}
	for i := inicol; i <= endcol; i++ {
		if row > 0 {
			if isSym(rune(lines[row-1][i])) { // up
				return true
			}
		}
		if row < len(lines)-1 {
			if isSym(rune(lines[row+1][i])) { // down
				return true
			}
		}
	}
	if endcol < len(lines[row])-1 {
		if isSym(rune(lines[row][endcol+1])) { // right
			return true
		}
		if row > 0 {
			if isSym(rune(lines[row-1][endcol+1])) { // up-right
				return true
			}
		}
		if row < len(lines)-1 {
			if isSym(rune(lines[row+1][endcol+1])) { // down-right
				return true
			}
		}
	}
	return false
}
