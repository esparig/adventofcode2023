package day3

import (
	"adventofcode2023/utils"
	"fmt"
)

func Day3() {

	fmt.Println("Day 3: input test")
	lines := utils.ReadLines("day3/input-test5.txt")
	numbers := getNumbers(lines, false)
	sum := 0
	for _, num := range numbers {
		fmt.Println(num.number, "From: (", num.initialdigit[0], ",", num.initialdigit[1], ") To: (",
			num.finaldigit[0], ",", num.finaldigit[1], ") - ", isPartNumber(lines, num))
		if isPartNumber(lines, num) {
			sum += num.number
		}
	}
	fmt.Println("Sum of part numbers: ", sum)

	fmt.Println("\nDay 3: input")
	lines = utils.ReadLines("day3/input.txt")
	numbers = getNumbers(lines, false)
	sum = 0

	for _, num := range numbers {
		fmt.Print(num.number, " From: (", num.initialdigit[0], ",", num.initialdigit[1], ") To: (",
			num.finaldigit[0], ",", num.finaldigit[1], ")")
		if isPartNumber(lines, num) {
			fmt.Println(" - true")
			sum += num.number
		} else {
			fmt.Println(" - false")
		}
	}

	fmt.Println("Sum of part numbers: ", sum)

}
